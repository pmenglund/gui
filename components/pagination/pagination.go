package pagination

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	public "github.com/pmenglund/gui/htmx"
	ihtmx "github.com/pmenglund/gui/internal/htmx"
	"github.com/pmenglund/gui/internal/render"
	"github.com/pmenglund/gui/internal/tw"
)

type Item struct {
	Label    string
	Href     string
	Current  bool
	Disabled bool
	HTMX     public.Props
}

type Props struct {
	ID         string
	Class      string
	Attributes []g.Node
	DataTestID string
	Items      []Item
}

func Pagination(p Props) g.Node {
	nodes := make([]g.Node, 0, len(p.Items))
	for _, item := range p.Items {
		linkAttrs := []g.Node{
			h.Class(tw.Join(
				"inline-flex min-w-10 items-center justify-center rounded-full border px-3 py-2 text-sm font-medium",
				tw.When(item.Current, "border-[rgb(var(--ui-primary))] bg-[rgb(var(--ui-primary))] text-[rgb(var(--ui-primary-foreground))]"),
				tw.When(!item.Current, "border-[rgb(var(--ui-border))] bg-[rgb(var(--ui-surface))]"),
				tw.When(item.Disabled, "pointer-events-none opacity-50"),
			)),
		}
		if item.Href != "" {
			linkAttrs = append(linkAttrs, h.Href(item.Href))
		}
		linkAttrs = append(linkAttrs, ihtmx.Attrs(item.HTMX)...)
		nodes = append(nodes, h.Li(h.A(append(linkAttrs, g.Text(item.Label))...)))
	}

	return h.Nav(
		append(
			render.Attrs(
				p.ID,
				tw.Join("flex items-center justify-center", p.Class),
				p.DataTestID,
				public.Props{},
				p.Attributes,
				h.Aria("label", "Pagination"),
			),
			h.Ul(h.Class("flex flex-wrap items-center gap-2"), g.Group(nodes)),
		)...,
	)
}
