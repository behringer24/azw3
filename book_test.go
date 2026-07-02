package azw3_test

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"testing"

	"github.com/behringer24/azw3"
)

func testPNG(t *testing.T) []byte {
	t.Helper()
	img := image.NewRGBA(image.Rect(0, 0, 4, 4))
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			img.Set(x, y, color.RGBA{R: 255, A: 255})
		}
	}
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, img); err != nil {
		t.Fatalf("encode test png: %v", err)
	}
	return buf.Bytes()
}

func TestWriteMinimalBook(t *testing.T) {
	b := azw3.New()
	b.SetTitle("Test Book")
	b.AddAuthor("Jane Doe")
	if err := b.AddLanguage("en"); err != nil {
		t.Fatalf("AddLanguage: %v", err)
	}

	if _, err := b.AddStylesheet("style", "body { font-family: serif; }"); err != nil {
		t.Fatalf("AddStylesheet: %v", err)
	}

	imgID, err := b.AddImage("cover.png", testPNG(t))
	if err != nil {
		t.Fatalf("AddImage: %v", err)
	}
	b.SetCoverImage(imgID)

	imgID2, err := b.AddImage("inline.png", testPNG(t))
	if err != nil {
		t.Fatalf("AddImage inline: %v", err)
	}

	_, err = b.AddChapter("ch1", "Chapter One", fmt.Sprintf(`<h1>Chapter One</h1><p>Hello, world.</p><img src="%s"/>`, imgID2))
	if err != nil {
		t.Fatalf("AddChapter ch1: %v", err)
	}
	_, err = b.AddChapter("ch2", "Chapter Two", "<h1>Chapter Two</h1><p>More text.</p>", 1)
	if err != nil {
		t.Fatalf("AddChapter ch2: %v", err)
	}

	data, err := b.Serialize()
	if err != nil {
		t.Fatalf("Serialize: %v", err)
	}
	if len(data) == 0 {
		t.Fatal("Serialize returned empty data")
	}

	dir := t.TempDir()
	path := filepath.Join(dir, "test.azw3")
	if err := b.Write(path); err != nil {
		t.Fatalf("Write: %v", err)
	}
	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("stat written file: %v", err)
	}
	if info.Size() == 0 {
		t.Fatal("written file is empty")
	}
	t.Logf("wrote %d bytes to %s", info.Size(), path)
}

func TestWriteWithNoChaptersFails(t *testing.T) {
	b := azw3.New()
	b.SetTitle("Empty Book")
	if _, err := b.Serialize(); err != azw3.ErrNoContent {
		t.Fatalf("expected ErrNoContent, got %v", err)
	}
}

func TestDuplicatePath(t *testing.T) {
	b := azw3.New()
	if _, err := b.AddChapter("dup", "T", "content"); err != nil {
		t.Fatalf("AddChapter: %v", err)
	}
	if _, err := b.AddChapter("dup", "T2", "content2"); err != azw3.ErrDuplicatePath {
		t.Fatalf("expected ErrDuplicatePath, got %v", err)
	}
}

func TestAddContributorValidatesRole(t *testing.T) {
	b := azw3.New()
	if err := b.AddContributor("Jane Doe", "edt"); err != nil {
		t.Fatalf("AddContributor with valid role: %v", err)
	}
	if err := b.AddContributor("Jane Doe", "not-a-real-role"); err != azw3.ErrInvalidContributorRole {
		t.Fatalf("expected ErrInvalidContributorRole, got %v", err)
	}
}

func TestAddLanguageRejectsInvalidTag(t *testing.T) {
	b := azw3.New()
	if err := b.AddLanguage("not-a-lang-tag-!!"); err != azw3.ErrInvalidLanguage {
		t.Fatalf("expected ErrInvalidLanguage, got %v", err)
	}
}

func TestNestedNavpointTOCIsRendered(t *testing.T) {
	b := azw3.New()
	b.SetTitle("Book With TOC")

	ch1, err := b.AddChapter("ch1", "Chapter One", "<p>one</p>")
	if err != nil {
		t.Fatalf("AddChapter ch1: %v", err)
	}
	ch2, err := b.AddChapter("ch2", "Chapter Two", "<p>two</p>", 1)
	if err != nil {
		t.Fatalf("AddChapter ch2: %v", err)
	}

	part := b.AddNavpoint("Part I", "", 0)
	part.AddNavpoint("Chapter One", string(ch1), 0)
	part.AddNavpoint("Chapter Two", string(ch2), 1)

	data, err := b.Serialize()
	if err != nil {
		t.Fatalf("Serialize: %v", err)
	}
	if len(data) == 0 {
		t.Fatal("Serialize returned empty data")
	}
}

func TestNavpointUnknownTargetFailsSerialize(t *testing.T) {
	b := azw3.New()
	b.SetTitle("Broken TOC")
	if _, err := b.AddChapter("ch1", "Chapter One", "<p>one</p>"); err != nil {
		t.Fatalf("AddChapter: %v", err)
	}
	b.AddNavpoint("Ghost Chapter", "does-not-exist", 0)

	_, err := b.Serialize()
	if !errors.Is(err, azw3.ErrChapterNotFound) {
		t.Fatalf("expected ErrChapterNotFound, got %v", err)
	}
}
