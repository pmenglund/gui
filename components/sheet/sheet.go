package sheet

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/pmenglund/gui/components/button"
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
	Title       string
	Description string
	Trigger     g.Node
	HTMX        public.Props
}

func Sheet(p Props, children ...g.Node) g.Node {
	id := a11y.ID("sheet", p.ID, p.Title, p.DataTestID)
	titleID := id + "-title"
	descID := id + "-description"
	trigger := p.Trigger
	if trigger == nil {
		trigger = button.Button(button.Props{Variant: button.VariantOutline}, g.Text("Open sheet"))
	}

	return h.Div(
		append(
			render.Attrs(
				id,
				"inline-flex",
				p.DataTestID,
				p.HTMX,
				p.Attributes,
				h.Data("ui-controller", "sheet"),
				h.Data("ui-state", "closed"),
			),
			h.Span(h.Data("ui-trigger", ""), trigger),
			h.Div(
				h.Data("ui-content", ""),
				h.Hidden("hidden"),
				h.Class("hidden fixed inset-0 z-50"),
				h.Button(h.Type("button"), h.Class("absolute inset-0 bg-black/40"), g.Attr("data-ui-close", "")),
				h.Aside(
					h.Class(tw.Join("absolute inset-y-0 right-0 z-10 grid w-full max-w-md gap-4 border-l bg-[rgb(var(--ui-surface))] p-6 shadow-2xl", p.Class)),
					h.Role("dialog"),
					h.Aria("modal", "true"),
					h.Aria("labelledby", titleID),
					g.If(p.Description != "", h.Aria("describedby", descID)),
					h.Div(
						h.Class("grid gap-2"),
						h.H2(h.ID(titleID), h.Class("text-xl font-semibold"), g.Text(p.Title)),
						g.If(p.Description != "", h.P(h.ID(descID), h.Class("text-sm text-[rgb(var(--ui-muted-foreground))]"), g.Text(p.Description))),
					),
					h.Div(h.Class("grid gap-4"), g.Group(children)),
					h.Div(h.Class("flex justify-end"), h.Button(h.Type("button"), h.Class("rounded-full border px-4 py-2 text-sm"), g.Attr("data-ui-close", ""), g.Text("Close"))),
				),
			),
		)...,
	)
}
