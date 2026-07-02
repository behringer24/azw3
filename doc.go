// Package azw3 provides a write-only, high-level API for building AZW3
// (Kindle KF8) ebooks. It wraps github.com/behringer24/mobi, which implements
// the underlying KF8 record and container format.
//
// Its shape mirrors github.com/writingtoole/epub: a Book you configure
// with plain Add/Set verb methods, no fluent chaining, no functional
// options, errors returned rather than panics.
//
// # Scope
//
// azw3 is write-only: it has no facilities for reading or modifying
// existing AZW3/MOBI files.
//
// # Content vs. navigation
//
// As in EPUB, content and navigation are separate concerns. AddChapter
// appends XHTML to the book's flat reading-order flow (the KF8 format,
// like EPUB's spine, has no native concept of nested content). AddNavpoint
// builds a Navpoint tree that can be nested arbitrarily deep, independent
// of that flat flow; each Navpoint either targets the Id of a chapter
// added via AddChapter, or has no target and exists purely to group its
// children under a label (e.g. a "Part I" heading).
//
// Whenever at least one Navpoint has been added, Write/Serialize render
// the tree as a nested table of contents page and insert it as the
// book's first chapter, with links resolving to hidden anchors AddChapter
// already placed at the start of each chapter's content. This is how
// azw3 offers a nested table of contents despite KF8's flat native
// navigation: the Kindle reader's own "Go To" chapter list is generated
// by leotaku/mobi directly from the flat chapter flow (one entry per
// chapter, in reading order) and does not reflect Navpoint nesting.
package azw3
