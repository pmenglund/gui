package toast

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/pmenglund/gui/components/button"
	public "github.com/pmenglund/gui/htmx"
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
	Action      g.Node
	HTMX        public.Props
}

// Toast renders a transient notification surface with the provided content.
func Toast(p Props, children ...g.Node) g.Node {
	trigger := p.Trigger
	if trigger == nil {
		trigger = button.Button(button.Props{Variant: button.VariantOutline}, g.Text("Show toast"))
	}

	content := []g.Node{
		h.Div(
			h.Class("grid gap-1"),
			h.H3(h.Class("font-semibold"), g.Text(p.Title)),
			g.If(p.Description != "", h.P(h.Class("text-sm text-[rgb(var(--ui-muted-foreground))]"), g.Text(p.Description))),
			g.Group(children),
		),
	}
	if p.Action != nil {
		content = append(content, p.Action)
	}
	content = append(content, h.Button(h.Type("button"), h.Class("rounded-full border px-3 py-2 text-sm"), g.Attr("data-ui-close", ""), g.Text("Dismiss")))

	return h.Div(
		append(
			render.Attrs(
				p.ID,
				tw.Join("inline-flex", p.Class),
				p.DataTestID,
				p.HTMX,
				p.Attributes,
				h.Data("ui-controller", "toast"),
				h.Data("ui-state", "closed"),
			),
			h.Span(h.Data("ui-trigger", ""), trigger),
			h.Div(
				h.Data("ui-content", ""),
				h.Hidden("hidden"),
				h.Class("hidden fixed bottom-6 right-6 z-50 flex max-w-sm items-start gap-3 rounded-[var(--ui-radius)] border bg-[rgb(var(--ui-surface))] p-4 shadow-2xl"),
				h.Role("status"),
				g.Group(content),
			),
		)...,
	)
}
