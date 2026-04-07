package label

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	public "github.com/pmenglund/gui/htmx"
	"github.com/pmenglund/gui/internal/render"
	"github.com/pmenglund/gui/internal/tw"
)

type Props struct {
	ID         string
	Class      string
	Attributes []g.Node
	DataTestID string
	For        string
	Muted      bool
	Required   bool
}

// Label renders a label element with optional required-state markup.
func Label(p Props, children ...g.Node) g.Node {
	contents := children
	if p.Required {
		contents = append(contents, h.Span(h.Class("ml-1 text-[rgb(var(--ui-danger))]"), g.Text("*")))
	}

	return h.Label(append(
		render.Attrs(
			p.ID,
			tw.Join("inline-flex items-center text-sm font-medium", tw.When(p.Muted, "text-[rgb(var(--ui-muted-foreground))]"), p.Class),
			p.DataTestID,
			public.Props{},
			p.Attributes,
			g.If(p.For != "", h.For(p.For)),
		),
		contents...,
	)...)
}
