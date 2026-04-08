package breadcrumbs

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
	Items      []Item
}

// Breadcrumbs renders a breadcrumb trail from the provided items.
func Breadcrumbs(p Props) g.Node {
	crumbs := make([]g.Node, 0, len(p.Items))
	for i, item := range p.Items {
		content := g.Node(g.Text(item.Label))
		if item.Href != "" && !item.Current {
			linkChildren := append(append([]g.Node{h.Href(item.Href), h.Class("hover:underline")}, ihtmx.Attrs(item.HTMX)...), g.Text(item.Label))
			content = h.A(linkChildren...)
		}
		crumbs = append(crumbs, h.Li(
			h.Class("inline-flex items-center gap-2"),
			content,
			g.If(i < len(p.Items)-1, h.Span(h.Class("text-[rgb(var(--ui-muted-foreground))]"), g.Text("/"))),
		))
	}

	return h.Nav(
		append(
			render.Attrs(
				p.ID,
				tw.Join("text-sm text-[rgb(var(--ui-muted-foreground))]", p.Class),
				p.DataTestID,
				public.Props{},
				p.Attributes,
				h.Aria("label", "Breadcrumb"),
			),
			h.Ol(h.Class("flex flex-wrap items-center gap-2"), g.Group(crumbs)),
		)...,
	)
}
