package input_test

import (
	"strings"
	"testing"

	"github.com/pmenglund/gui/components/input"
	public "github.com/pmenglund/gui/htmx"
	"github.com/pmenglund/gui/internal/testutil"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestInputPreservesTypedSemanticAndHTMXAttrs(t *testing.T) {
	t.Parallel()

	got := testutil.Render(t, input.Input(input.Props{
		Type:    "email",
		Invalid: true,
		HTMX:    public.Props{Target: "#target"},
		Attributes: []g.Node{
			h.Type("text"),
			h.Aria("invalid", "false"),
			g.Attr("hx-target", "#other"),
			g.Attr("title", "extra"),
		},
	}))

	for _, want := range []string{
		`type="email"`,
		`aria-invalid="true"`,
		`hx-target="#target"`,
		`title="extra"`,
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("rendered input missing %s: %s", want, got)
		}
	}

	for _, blocked := range []string{
		`type="text"`,
		`aria-invalid="false"`,
		`hx-target="#other"`,
	} {
		if strings.Contains(got, blocked) {
			t.Fatalf("rendered input unexpectedly contained %s: %s", blocked, got)
		}
	}
}
