package card

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

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
	Footer      g.Node
	HTMX        public.Props
}

func Card(p Props, children ...g.Node) g.Node {
	body := []g.Node{}
	if p.Title != "" || p.Description != "" {
		body = append(body, h.Header(
			h.Class("grid gap-1"),
			g.If(p.Title != "", h.H3(h.Class("text-lg font-semibold"), g.Text(p.Title))),
			g.If(p.Description != "", h.P(h.Class("text-sm text-[rgb(var(--ui-muted-foreground))]"), g.Text(p.Description))),
		))
	}
	body = append(body, children...)
	if p.Footer != nil {
		body = append(body, h.Footer(h.Class("pt-2"), p.Footer))
	}

	return h.Section(append(
		render.Attrs(
			p.ID,
			tw.Join("grid gap-4 rounded-[var(--ui-radius)] border bg-[rgb(var(--ui-surface))] p-6 shadow-sm", p.Class),
			p.DataTestID,
			p.HTMX,
			p.Attributes,
		),
		body...,
	)...)
}
