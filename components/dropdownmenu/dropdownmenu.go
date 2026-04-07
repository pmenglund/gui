package dropdownmenu

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/pmenglund/gui/components/button"
	public "github.com/pmenglund/gui/htmx"
	ihtmx "github.com/pmenglund/gui/internal/htmx"
	"github.com/pmenglund/gui/internal/render"
	"github.com/pmenglund/gui/internal/tw"
)

type Item struct {
	Label     string
	Href      string
	Dangerous bool
	HTMX      public.Props
}

type Props struct {
	ID         string
	Class      string
	Attributes []g.Node
	DataTestID string
	Trigger    g.Node
	Items      []Item
	HTMX       public.Props
}

// DropdownMenu renders a trigger and popover menu for the provided items.
func DropdownMenu(p Props) g.Node {
	trigger := p.Trigger
	if trigger == nil {
		trigger = button.Button(button.Props{Variant: button.VariantOutline}, g.Text("Open menu"))
	}

	menuItems := make([]g.Node, 0, len(p.Items))
	for _, item := range p.Items {
		className := tw.Join(
			"block rounded-[calc(var(--ui-radius)*0.75)] px-3 py-2 text-sm hover:bg-[rgb(var(--ui-surface-strong))]",
			tw.When(item.Dangerous, "text-[rgb(var(--ui-danger))]"),
		)
		if item.Href != "" {
			linkChildren := append(append([]g.Node{h.Href(item.Href), h.Class(className)}, ihtmx.Attrs(item.HTMX)...), g.Text(item.Label))
			menuItems = append(menuItems, h.A(linkChildren...))
			continue
		}
		buttonChildren := append(append([]g.Node{h.Type("button"), h.Class(className)}, ihtmx.Attrs(item.HTMX)...), g.Text(item.Label))
		menuItems = append(menuItems, h.Button(buttonChildren...))
	}

	return h.Div(
		append(
			render.Attrs(
				p.ID,
				tw.Join("relative inline-flex", p.Class),
				p.DataTestID,
				p.HTMX,
				p.Attributes,
				h.Data("ui-controller", "dropdownmenu"),
				h.Data("ui-state", "closed"),
			),
			h.Span(h.Data("ui-trigger", ""), trigger),
			h.Div(
				h.Data("ui-content", ""),
				h.Hidden("hidden"),
				h.Class("hidden absolute right-0 top-full z-40 mt-2 grid min-w-56 gap-1 rounded-[var(--ui-radius)] border bg-[rgb(var(--ui-surface))] p-2 shadow-xl"),
				g.Group(menuItems),
			),
		)...,
	)
}
