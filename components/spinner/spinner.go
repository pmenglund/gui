package spinner

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	public "github.com/pmenglund/goth/htmx"
	"github.com/pmenglund/goth/internal/render"
	"github.com/pmenglund/goth/internal/tw"
)

type Props struct {
	ID         string
	Class      string
	Attributes []g.Node
	DataTestID string
	Label      string
	Size       string
}

// Spinner renders an animated loading indicator.
func Spinner(p Props) g.Node {
	size := p.Size
	if size == "" {
		size = "h-5 w-5"
	}

	label := p.Label
	if label == "" {
		label = "Loading"
	}

	return h.Span(
		append(
			render.Attrs(
				p.ID,
				tw.Join("inline-flex items-center gap-2 text-sm text-[rgb(var(--ui-muted-foreground))]", p.Class),
				p.DataTestID,
				public.Props{},
				p.Attributes,
				h.Role("status"),
			),
			h.Span(h.Class(tw.Join(size, "inline-block animate-spin rounded-full border-2 border-[rgb(var(--ui-muted))] border-t-[rgb(var(--ui-primary))]"))),
			h.Span(g.Text(label)),
		)...,
	)
}
