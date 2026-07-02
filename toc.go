package azw3

import (
	"fmt"
	"html"
	"sort"
	"strings"
)

// sortedNavpoints returns nodes ordered by their order field, breaking
// ties by insertion order (the order the corresponding AddNavpoint calls
// were made in).
func sortedNavpoints(nodes []*Navpoint) []*Navpoint {
	sorted := make([]*Navpoint, len(nodes))
	copy(sorted, nodes)
	sort.SliceStable(sorted, func(i, j int) bool {
		return sorted[i].order < sorted[j].order
	})
	return sorted
}

// renderNavpoints renders nodes as a nested <ol> table of contents,
// resolving each non-empty target against anchors (chapter Id -> anchor
// id). Returns ErrChapterNotFound if a target doesn't match any chapter.
func renderNavpoints(nodes []*Navpoint, anchors map[Id]string) (string, error) {
	if len(nodes) == 0 {
		return "", nil
	}

	var b strings.Builder
	b.WriteString("<ol>")
	for _, n := range nodes {
		item, err := renderNavpoint(n, anchors)
		if err != nil {
			return "", err
		}
		b.WriteString(item)
	}
	b.WriteString("</ol>")
	return b.String(), nil
}

func renderNavpoint(n *Navpoint, anchors map[Id]string) (string, error) {
	label := html.EscapeString(n.label)

	entry := label
	if n.target != "" {
		anchor, ok := anchors[Id(n.target)]
		if !ok {
			return "", fmt.Errorf("%w: %q", ErrChapterNotFound, n.target)
		}
		entry = fmt.Sprintf(`<a href="#%s">%s</a>`, anchor, label)
	}

	children, err := renderNavpoints(sortedNavpoints(n.children), anchors)
	if err != nil {
		return "", err
	}

	return "<li>" + entry + children + "</li>", nil
}
