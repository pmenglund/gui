package render_test

import (
	"strings"
	"testing"

	public "github.com/pmenglund/gui/htmx"
	"github.com/pmenglund/gui/internal/render"
	"github.com/pmenglund/gui/internal/testutil"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestAttrsFiltersConflictingProtectedAttributes(t *testing.T) {
	t.Parallel()

	node := h.Div(render.Attrs(
		"",
		"",
		"",
		public.Props{Target: "#target"},
		[]g.Node{
			h.Type("text"),
			h.Aria("labelledby", "raw-title"),
			g.Attr("hx-target", "#other"),
			g.Attr("title", "extra"),
		},
		h.Type("email"),
		h.Aria("labelledby", "safe-title"),
	)...)

	got := testutil.Render(t, node)

	for _, want := range []string{
		`type="email"`,
		`aria-labelledby="safe-title"`,
		`hx-target="#target"`,
		`title="extra"`,
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("rendered markup missing %s: %s", want, got)
		}
	}

	for _, blocked := range []string{
		`type="text"`,
		`aria-labelledby="raw-title"`,
		`hx-target="#other"`,
	} {
		if strings.Contains(got, blocked) {
			t.Fatalf("rendered markup unexpectedly contained %s: %s", blocked, got)
		}
	}
}

func TestAttrsReservesDataUIAttributes(t *testing.T) {
	t.Parallel()

	node := h.Div(render.Attrs(
		"",
		"",
		"",
		public.Props{},
		[]g.Node{
			g.Attr("data-ui-controller", "bad"),
			g.Attr("data-ui-custom", "bad"),
			g.Attr("data-qa", "safe"),
		},
	)...)

	got := testutil.Render(t, node)

	if !strings.Contains(got, `data-qa="safe"`) {
		t.Fatalf("rendered markup missing safe passthrough attribute: %s", got)
	}
	for _, blocked := range []string{
		`data-ui-controller="bad"`,
		`data-ui-custom="bad"`,
	} {
		if strings.Contains(got, blocked) {
			t.Fatalf("rendered markup unexpectedly contained reserved attribute %s: %s", blocked, got)
		}
	}
}
