// Command basic generates a small sample AZW3 ebook, demonstrating
// metadata, a cover image, a stylesheet, an inline image, and a nested
// table of contents. Run it from the repository root:
//
//	go run ./examples/basic
//
// It writes example.azw3 to the current directory.
package main

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"

	"github.com/behringer24/azw3"
)

func main() {
	b := azw3.New()
	b.SetTitle("The Lighthouse Keeper's Log")
	b.AddAuthor("A. Sample Author")
	b.AddDescription("A short, fictional sample book generated to exercise the azw3 library.")
	b.AddPublisher("azw3 examples")
	if err := b.AddLanguage("en"); err != nil {
		log.Fatalf("AddLanguage: %v", err)
	}
	if err := b.AddContributor("A. N. Editor", "edt"); err != nil {
		log.Fatalf("AddContributor: %v", err)
	}

	if _, err := b.AddStylesheet("style", stylesheet); err != nil {
		log.Fatalf("AddStylesheet: %v", err)
	}

	coverID, err := b.AddImage("cover", coverImage())
	if err != nil {
		log.Fatalf("AddImage cover: %v", err)
	}
	b.SetCoverImage(coverID)

	lampID, err := b.AddImage("lamp-illustration", lampImage())
	if err != nil {
		log.Fatalf("AddImage lamp: %v", err)
	}

	partOne := b.AddNavpoint("Part One: Arrival", "", 0)
	partTwo := b.AddNavpoint("Part Two: The Storm", "", 1)

	ch1, err := b.AddChapter("ch1", "Chapter One: The Appointment", chapterOne, 0)
	if err != nil {
		log.Fatalf("AddChapter ch1: %v", err)
	}
	partOne.AddNavpoint("Chapter One: The Appointment", string(ch1), 0)

	ch2, err := b.AddChapter("ch2", "Chapter Two: The First Night", chapterTwo(lampID), 1)
	if err != nil {
		log.Fatalf("AddChapter ch2: %v", err)
	}
	partOne.AddNavpoint("Chapter Two: The First Night", string(ch2), 1)

	ch3, err := b.AddChapter("ch3", "Chapter Three: The Wind Rises", chapterThree, 2)
	if err != nil {
		log.Fatalf("AddChapter ch3: %v", err)
	}
	partTwo.AddNavpoint("Chapter Three: The Wind Rises", string(ch3), 0)

	if err := b.Write("example.azw3"); err != nil {
		log.Fatalf("Write: %v", err)
	}
	log.Println("wrote example.azw3")
}

const stylesheet = `
body { font-family: serif; }
h1 { text-align: center; font-variant: small-caps; }
p { text-indent: 1.5em; margin: 0; }
p.first { text-indent: 0; }
`

const chapterOne = `
<h1>Chapter One: The Appointment</h1>
<p class="first">The letter arrived on a Tuesday, smelling faintly of salt
and tar, and informing Miram Cole that she had been appointed keeper of the
Fenwick Point light for the coming winter.</p>
<p>She read it twice by the kitchen window, then a third time aloud to the
cat, who was unimpressed.</p>
<p>By Thursday she had packed one trunk, sold her armchair, and bought a
lantern she did not yet know how to trim.</p>
`

func chapterTwo(lampID azw3.Id) string {
	return `
<h1>Chapter Two: The First Night</h1>
<p class="first">The lighthouse smelled of oil and cold stone. Miram climbed
the spiral stair with her hand trailing along the iron rail, counting steps
out of habit rather than need.</p>
<p>At the top, the great lamp waited, dark and patient.</p>
<p><img src="` + string(lampID) + `" alt="A simple line drawing of an oil lamp"/></p>
<p>She lit it just as the sun dropped into the sea, and stood a long while
watching the beam sweep the water, before she remembered she was hungry.</p>
`
}

const chapterThree = `
<h1>Chapter Three: The Wind Rises</h1>
<p class="first">By the second week, Miram had learned the tower's every
creak, and could tell from the pitch of the wind alone whether the night
would be calm or long.</p>
<p>Tonight, it was long.</p>
<p>She banked the lamp, checked the oil twice, and settled into the
keeper's chair to wait out the storm, log book open on her knee.</p>
`

// coverImage renders a simple placeholder cover: a vertical gradient from
// deep blue to near-black, evoking a night sky over the sea.
func coverImage() []byte {
	const w, h = 600, 800
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		t := float64(y) / float64(h)
		c := color.RGBA{
			R: uint8(10 + 20*(1-t)),
			G: uint8(20 + 40*(1-t)),
			B: uint8(40 + 120*(1-t)),
			A: 255,
		}
		for x := 0; x < w; x++ {
			img.Set(x, y, c)
		}
	}
	// A simple "lighthouse beam" wedge near the top third.
	beamY := h / 3
	for y := beamY - 4; y <= beamY+4; y++ {
		for x := w / 2; x < w; x++ {
			img.Set(x, y, color.RGBA{R: 255, G: 240, B: 200, A: 255})
		}
	}
	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, img, &jpeg.Options{Quality: 85}); err != nil {
		log.Fatalf("encode cover: %v", err)
	}
	return buf.Bytes()
}

// lampImage renders a very simple line-art oil lamp: a triangle body on a
// rectangular base.
func lampImage() []byte {
	const w, h = 200, 200
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	white := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	black := color.RGBA{A: 255}
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			img.Set(x, y, white)
		}
	}
	// Base
	for y := 150; y < 170; y++ {
		for x := 60; x < 140; x++ {
			img.Set(x, y, black)
		}
	}
	// Body (triangle-ish, widening toward the base)
	for y := 60; y < 150; y++ {
		half := (y - 60) / 2
		for x := 100 - half; x <= 100+half; x++ {
			img.Set(x, y, black)
		}
	}
	// Flame
	for y := 30; y < 60; y++ {
		half := (60 - y) / 6
		for x := 100 - half; x <= 100+half; x++ {
			img.Set(x, y, black)
		}
	}
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, img); err != nil {
		log.Fatalf("encode lamp: %v", err)
	}
	return buf.Bytes()
}
