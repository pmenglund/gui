package navbar

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	public "github.com/pmenglund/goth/htmx"
	ihtmx "github.com/pmenglund/goth/internal/htmx"
	"github.com/pmenglund/goth/internal/render"
	"github.com/pmenglund/goth/internal/tw"
)

type Item struct {
	Label   string
	Href    string
	Current bool
	HTMX    public.Props
}

type Props struct {
	ID         string
	Class      string
	Attributes []g.Node
	DataTestID string
	Brand      g.Node
	Items      []Item
	Actions    []g.Node
}

// Navbar renders a horizontal navigation bar from the provided items.
func Navbar(p Props) g.Node {
	links := make([]g.Node, 0, len(p.Items))
	for _, item := range p.Items {
		attrs := []g.Node{
			h.Class(tw.Join(
				"rounded-full px-3 py-2 text-sm font-medium transition hover:bg-[rgb(var(--ui-surface-strong))]",
				tw.When(item.Current, "bg-[rgb(var(--ui-surface-strong))]"),
			)),
			h.Href(item.Href),
		}
		if item.Current {
			attrs = append(attrs, h.Aria("current", "page"))
		}
		attrs = append(attrs, ihtmx.Attrs(item.HTMX)...)
		linkChildren := append(attrs, g.Text(item.Label))
		links = append(links, h.A(linkChildren...))
	}

	return h.Nav(
		append(
			render.Attrs(
				p.ID,
				tw.Join("flex flex-col gap-4 rounded-[var(--ui-radius)] border bg-[rgb(var(--ui-surface))] p-4 shadow-sm md:flex-row md:items-center md:justify-between", p.Class),
				p.DataTestID,
				public.Props{},
				p.Attributes,
			),
			h.Div(h.Class("flex items-center gap-3"), p.Brand),
			h.Div(h.Class("flex flex-wrap items-center gap-2"), g.Group(links)),
			h.Div(h.Class("flex flex-wrap items-center gap-2"), g.Group(p.Actions)),
		)...,
	)
}
