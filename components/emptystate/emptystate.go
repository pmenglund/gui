package emptystate

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	public "github.com/pmenglund/goth/htmx"
	"github.com/pmenglund/goth/internal/render"
	"github.com/pmenglund/goth/internal/tw"
)

type Props struct {
	ID           string
	Class        string
	Attributes   []g.Node
	DataTestID   string
	Eyebrow      string
	Title        string
	Description  string
	Illustration g.Node
	Action       g.Node
	HTMX         public.Props
}

// EmptyState renders a placeholder view for empty content.
func EmptyState(p Props) g.Node {
	return h.Section(
		append(
			render.Attrs(
				p.ID,
				tw.Join("grid place-items-center gap-4 rounded-[var(--ui-radius)] border border-dashed bg-[rgb(var(--ui-surface))] p-10 text-center", p.Class),
				p.DataTestID,
				p.HTMX,
				p.Attributes,
			),
			g.If(p.Illustration != nil, p.Illustration),
			g.If(p.Eyebrow != "", h.P(h.Class("ui-kicker"), g.Text(p.Eyebrow))),
			h.Div(
				h.Class("grid gap-2"),
				h.H2(h.Class("text-2xl font-semibold"), g.Text(p.Title)),
				g.If(p.Description != "", h.P(h.Class("text-sm text-[rgb(var(--ui-muted-foreground))]"), g.Text(p.Description))),
			),
			g.If(p.Action != nil, p.Action),
		)...,
	)
}
