package azw3

// Navpoint is a node in the book's table of contents tree. See the
// package doc ("Content vs. navigation") for why this exists separately
// from the flat content flow built by Book.AddChapter.
//
// A Navpoint's target must be the Id of a chapter added via
// Book.AddChapter, or the empty string for a pure grouping node with no
// link of its own (e.g. a "Part I" heading whose children are the
// chapters that belong to it). This mirrors the fact that the
// underlying KF8 writer's native navigation is generated per whole
// chapter; there is no lower-level addressing for arbitrary points
// within a chapter's content.
type Navpoint struct {
	label  string
	target string
	order  int

	children []*Navpoint
}

// AddNavpoint adds a child navigation point under n, pointing at target
// (the Id of a chapter added via Book.AddChapter, or "" for a grouping
// node with no link of its own). order controls this node's position
// among its siblings.
func (n *Navpoint) AddNavpoint(label, target string, order int) *Navpoint {
	child := &Navpoint{label: label, target: target, order: order}
	n.children = append(n.children, child)
	return child
}

// Label returns the navigation point's display label.
func (n *Navpoint) Label() string {
	return n.label
}

// Target returns the chapter Id this navigation point resolves to, or ""
// if it is a grouping node with no link of its own.
func (n *Navpoint) Target() string {
	return n.target
}
