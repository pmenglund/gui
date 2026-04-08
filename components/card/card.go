package card

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/pmenglund/goth/components/classmode"
	public "github.com/pmenglund/goth/htmx"
	"github.com/pmenglund/goth/internal/render"
	"github.com/pmenglund/goth/internal/tw"
)

// Props configures Card rendering.
type Props struct {
	ID               string
	Class            string
	ClassMode        classmode.ClassMode
	Attributes       []g.Node
	DataTestID       string
	Title            string
	Description      string
	Footer           g.Node
	HeaderClass      string
	TitleClass       string
	DescriptionClass string
	FooterClass      string
	HTMX             public.Props
}

// Card renders a bordered surface for grouping related content.
func Card(p Props, children ...g.Node) g.Node {
	body := []g.Node{}
	if p.Title != "" || p.Description != "" {
		body = append(body, h.Header(
			h.Class(tw.Join("grid gap-1", p.HeaderClass)),
			g.If(p.Title != "", h.H3(h.Class(tw.Join("text-lg font-semibold", p.TitleClass)), g.Text(p.Title))),
			g.If(p.Description != "", h.P(h.Class(tw.Join("text-sm text-[rgb(var(--ui-muted-foreground))]", p.DescriptionClass)), g.Text(p.Description))),
		))
	}
	body = append(body, children...)
	if p.Footer != nil {
		body = append(body, h.Footer(h.Class(tw.Join("pt-2", p.FooterClass)), p.Footer))
	}

	return h.Section(append(
		render.Attrs(
			p.ID,
			tw.Classes(p.ClassMode, "grid gap-4 rounded-[var(--ui-radius)] border bg-[rgb(var(--ui-surface))] p-6 shadow-sm", p.Class),
			p.DataTestID,
			p.HTMX,
			p.Attributes,
		),
		body...,
	)...)
}
