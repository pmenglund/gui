package checkbox

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	public "github.com/pmenglund/gui/htmx"
	"github.com/pmenglund/gui/internal/a11y"
	"github.com/pmenglund/gui/internal/render"
	"github.com/pmenglund/gui/internal/tw"
)

type Props struct {
	ID          string
	Class       string
	Attributes  []g.Node
	DataTestID  string
	Name        string
	Value       string
	Label       string
	Checked     bool
	Disabled    bool
	Required    bool
	Invalid     bool
	DescribedBy string
	HTMX        public.Props
}

// Checkbox renders a checkbox control with optional inline content.
func Checkbox(p Props, children ...g.Node) g.Node {
	id := a11y.ID("checkbox", p.ID, p.Name, p.Label, p.DataTestID)
	labelContent := children
	if len(labelContent) == 0 && p.Label != "" {
		labelContent = []g.Node{g.Text(p.Label)}
	}

	inputAttrs := render.Attrs(
		id,
		tw.Join(
			"h-4 w-4 rounded border bg-[rgb(var(--ui-surface))] text-[rgb(var(--ui-primary))] focus:outline-none focus-visible:ring-2 focus-visible:ring-[rgb(var(--ui-ring))]",
			tw.When(p.Invalid, "border-[rgb(var(--ui-danger))]"),
		),
		p.DataTestID,
		p.HTMX,
		nil,
		h.Type("checkbox"),
		g.If(p.Name != "", h.Name(p.Name)),
		g.If(p.Value != "", h.Value(p.Value)),
		g.If(p.Checked, h.Checked()),
		g.If(p.Disabled, h.Disabled()),
		g.If(p.Required, h.Required()),
		g.If(p.Invalid, h.Aria("invalid", "true")),
		g.If(p.DescribedBy != "", h.Aria("describedby", p.DescribedBy)),
	)

	return h.Label(append(
		render.Attrs("", tw.Join("inline-flex items-center gap-3 text-sm", p.Class), "", public.Props{}, p.Attributes),
		h.Input(inputAttrs...),
		h.Span(g.Group(labelContent)),
	)...)
}
