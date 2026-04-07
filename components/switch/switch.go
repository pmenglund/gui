package switchui

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	public "github.com/pmenglund/gui/htmx"
	"github.com/pmenglund/gui/internal/a11y"
	"github.com/pmenglund/gui/internal/render"
	"github.com/pmenglund/gui/internal/tw"
)

type Props struct {
	ID         string
	Class      string
	Attributes []g.Node
	DataTestID string
	Name       string
	Label      string
	Checked    bool
	Disabled   bool
	Required   bool
	HTMX       public.Props
}

// Switch renders a binary toggle control with optional inline content.
func Switch(p Props, children ...g.Node) g.Node {
	id := a11y.InstanceID("switch", p.ID, p.Name, p.Label, p.DataTestID)
	labelContent := children
	if len(labelContent) == 0 && p.Label != "" {
		labelContent = []g.Node{g.Text(p.Label)}
	}

	input := h.Input(render.Attrs(
		id,
		"peer sr-only",
		p.DataTestID,
		p.HTMX,
		nil,
		h.Type("checkbox"),
		g.If(p.Name != "", h.Name(p.Name)),
		g.If(p.Checked, h.Checked()),
		g.If(p.Disabled, h.Disabled()),
		g.If(p.Required, h.Required()),
	)...)

	return h.Label(append(
		render.Attrs("", tw.Join("inline-flex items-center gap-3 text-sm", p.Class), "", public.Props{}, p.Attributes),
		input,
		h.Span(h.Class("relative inline-flex h-6 w-11 items-center rounded-full bg-[rgb(var(--ui-muted))] transition peer-checked:bg-[rgb(var(--ui-primary))]"),
			h.Span(h.Class("absolute left-1 h-4 w-4 rounded-full bg-white transition peer-checked:translate-x-5")),
		),
		h.Span(g.Group(labelContent)),
	)...)
}
