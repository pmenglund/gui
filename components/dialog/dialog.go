package dialog

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
	Footer      g.Node
	Open        bool
	HTMX        public.Props
}

// Dialog renders a modal dialog shell and its trigger content.
func Dialog(p Props, children ...g.Node) g.Node {
	id := a11y.InstanceID("dialog", p.ID, p.Title, p.DataTestID)
	titleID := id + "-title"
	descID := id + "-description"
	trigger := p.Trigger
	if trigger == nil {
		trigger = button.Button(button.Props{Variant: button.VariantOutline}, g.Text("Open dialog"))
	}

	content := h.Div(
		h.Data("ui-content", ""),
		h.Class("hidden fixed inset-0 z-50 grid place-items-center bg-black/50 p-4"),
		g.If(!p.Open, h.Hidden("hidden")),
		h.Button(h.Type("button"), h.Class("absolute inset-0 cursor-default"), g.Attr("data-ui-close", "")),
		h.Div(
			h.Class(tw.Join("relative z-10 grid w-full max-w-lg gap-4 rounded-[var(--ui-radius)] border bg-[rgb(var(--ui-surface))] p-6 shadow-2xl", p.Class)),
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
			h.Div(h.Class("flex items-center justify-end gap-3"), p.Footer, h.Button(h.Type("button"), h.Class("rounded-full border px-4 py-2 text-sm"), g.Attr("data-ui-close", ""), g.Text("Close"))),
		),
	)

	return h.Div(
		append(
			render.Attrs(
				id,
				"inline-flex",
				p.DataTestID,
				p.HTMX,
				p.Attributes,
				h.Data("ui-controller", "dialog"),
				h.Data("ui-state", ternary(p.Open, "open", "closed")),
			),
			h.Span(h.Data("ui-trigger", ""), trigger),
			content,
		)...,
	)
}

func ternary(ok bool, a, b string) string {
	if ok {
		return a
	}
	return b
}
