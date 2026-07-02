# azw3

A write-only Go library for building AZW3 (Kindle KF8) ebooks — add metadata,
chapters, images, and a nested table of contents, then write out a `.azw3`
file.

`azw3` does not read or modify existing AZW3/MOBI files. It only creates new
ones from scratch.

## Status

Early stage. The core API is implemented and tested (metadata, chapters,
images, cover, stylesheets, nested navigation, output), but it hasn't yet
been exercised against a wide range of real-world content or Kindle
hardware/apps. Font embedding is not supported yet. See
[Design notes and limitations](#design-notes-and-limitations) below before
relying on it for anything important.

## Installation

```sh
go get github.com/behringer24/azw3
```

## Quick start

```go
package main

import (
	"log"
	"os"

	"github.com/behringer24/azw3"
)

func main() {
	b := azw3.New()
	b.SetTitle("De Vita Caesarum")
	b.AddAuthor("Suetonius")
	if err := b.AddLanguage("la"); err != nil {
		log.Fatal(err)
	}

	cover, err := os.ReadFile("cover.jpg")
	if err != nil {
		log.Fatal(err)
	}
	coverID, err := b.AddImage("cover", cover)
	if err != nil {
		log.Fatal(err)
	}
	b.SetCoverImage(coverID)

	ch1, err := b.AddChapter("ch1", "Julius Caesar", "<h1>Julius Caesar</h1><p>...</p>")
	if err != nil {
		log.Fatal(err)
	}
	ch2, err := b.AddChapter("ch2", "Augustus", "<h1>Augustus</h1><p>...</p>", 1)
	if err != nil {
		log.Fatal(err)
	}

	// Optional: a nested table of contents, independent of the flat
	// chapter order above.
	toc := b.AddNavpoint("The Twelve Caesars", "", 0)
	toc.AddNavpoint("Julius Caesar", string(ch1), 0)
	toc.AddNavpoint("Augustus", string(ch2), 1)

	if err := b.Write("caesars.azw3"); err != nil {
		log.Fatal(err)
	}
}
```

## API overview

The API mirrors [writingtoole/epub](https://github.com/writingtoole/epub):
plain `Add`/`Set` verb methods, no fluent chaining, no functional options,
errors returned rather than panics.

- **Metadata** — `SetTitle`, `AddAuthor`, `AddContributor` (validated
  against the MARC relator code vocabulary), `AddPublisher`, `AddLanguage`
  (BCP 47), `AddDescription`, `SetPublishedDate`, `SetUniqueID`,
  `SetFixedLayout`, `SetRightToLeft`.
- **Content** — `AddChapter` appends XHTML to the book's flat reading-order
  flow. `AddImage`/`AddImageFile` register images and return a ready-to-use
  `<img src="...">` value. `AddStylesheet`/`AddStylesheetFile` register CSS
  that is automatically applied to every chapter.
- **Navigation** — `AddNavpoint` builds a `Navpoint` tree that can be
  nested arbitrarily deep, independently of the flat content flow. Each
  node either targets the `Id` of a chapter or has no target and exists
  purely to group its children under a label (e.g. a "Part I" heading).
  Whenever at least one `Navpoint` exists, `Write`/`Serialize` render the
  tree as a table of contents page and insert it as the book's first
  chapter.
- **Cover** — `SetCoverImage` promotes a previously added image to be the
  book's cover.
- **Output** — `Write(path)` and `Serialize()`.

See the [package documentation](https://pkg.go.dev/github.com/behringer24/azw3)
(or run `go doc` in this repository) for full details on every method.

## Design notes and limitations

AZW3/KF8 is a reverse-engineered, sparsely documented format. This library
is built on top of [leotaku/mobi](https://github.com/leotaku/mobi), which
implements the underlying KF8 record and container format, and deviates
from a literal EPUB-style API in a few places the format forces:

- **Content is flat.** Like EPUB's spine, KF8 has no native concept of
  nested chapters. The Kindle reader's own "Go To" chapter list is
  generated directly from the flat chapter list added via `AddChapter`
  and does not reflect `Navpoint` nesting — only azw3's own generated
  table of contents page does.
- **No filesystem-like resource paths.** Unlike EPUB's zip archive, KF8
  has no addressable files. The `path` argument to `AddChapter`/`AddImage`/
  `AddStylesheet` is a caller-chosen, book-internal name used only to
  detect duplicates, not a real path.
- **`Navpoint` targets whole chapters only.** The underlying writer
  generates its native navigation index strictly one entry per top-level
  chapter; there's no lower-level addressing for arbitrary points within
  a chapter's content.
- **No font embedding yet.** Support for this in the underlying writer is
  unclear/undocumented; left out until it can be verified.
- **No JavaScript.** Kindle's rendering engine doesn't execute it, so
  there's no `AddJavaScript`.

## Acknowledgments

- [leotaku/mobi](https://github.com/leotaku/mobi) implements the KF8
  writer this library builds on.
- [writingtoole/epub](https://github.com/writingtoole/epub) is the API
  this library's shape is modeled after.
- [Calibre](https://github.com/kovidgoyal/calibre)'s MOBI/KF8 writer
  source was a helpful reference for understanding the format itself.

## License

[MIT](LICENSE)
