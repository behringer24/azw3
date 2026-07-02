package azw3

import "testing"

func TestRenderNavpointsNestedWithGrouping(t *testing.T) {
	root := []*Navpoint{
		{label: "Preface", target: "pref", order: 0},
	}
	part := &Navpoint{label: "Part I", target: "", order: 1}
	part.children = []*Navpoint{
		{label: "Chapter One", target: "ch1", order: 0},
		{label: "Chapter Two", target: "ch2", order: 1},
	}
	root = append(root, part)

	anchors := map[Id]string{
		"pref": "azw3-ch1",
		"ch1":  "azw3-ch2",
		"ch2":  "azw3-ch3",
	}

	got, err := renderNavpoints(root, anchors)
	if err != nil {
		t.Fatalf("renderNavpoints: %v", err)
	}

	want := `<ol>` +
		`<li><a href="#azw3-ch1">Preface</a></li>` +
		`<li>Part I<ol>` +
		`<li><a href="#azw3-ch2">Chapter One</a></li>` +
		`<li><a href="#azw3-ch3">Chapter Two</a></li>` +
		`</ol></li>` +
		`</ol>`
	if got != want {
		t.Fatalf("got:\n%s\nwant:\n%s", got, want)
	}
}

func TestRenderNavpointsUnknownTarget(t *testing.T) {
	nodes := []*Navpoint{{label: "Ghost", target: "nope"}}
	if _, err := renderNavpoints(nodes, map[Id]string{}); err == nil {
		t.Fatal("expected error for unknown target, got nil")
	}
}

func TestRenderNavpointsEscapesLabel(t *testing.T) {
	nodes := []*Navpoint{{label: `A & B <em>"C"</em>`, target: ""}}
	got, err := renderNavpoints(nodes, map[Id]string{})
	if err != nil {
		t.Fatalf("renderNavpoints: %v", err)
	}
	want := `<ol><li>A &amp; B &lt;em&gt;&#34;C&#34;&lt;/em&gt;</li></ol>`
	if got != want {
		t.Fatalf("got:\n%s\nwant:\n%s", got, want)
	}
}

func TestSortedNavpointsBreaksTiesByInsertionOrder(t *testing.T) {
	a := &Navpoint{label: "a", order: 1}
	b := &Navpoint{label: "b", order: 0}
	c := &Navpoint{label: "c", order: 0}
	sorted := sortedNavpoints([]*Navpoint{a, b, c})
	if len(sorted) != 3 || sorted[0] != b || sorted[1] != c || sorted[2] != a {
		t.Fatalf("unexpected order: %+v", sorted)
	}
}
