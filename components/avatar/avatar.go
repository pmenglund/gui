package avatar

import (
	"strings"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	public "github.com/pmenglund/gui/htmx"
	"github.com/pmenglund/gui/internal/render"
	"github.com/pmenglund/gui/internal/tw"
)

type Props struct {
	ID         string
	Class      string
	Attributes []g.Node
	DataTestID string
	Name       string
	Src        string
	Alt        string
	Size       string
}

func Avatar(p Props) g.Node {
	size := p.Size
	if size == "" {
		size = "h-12 w-12"
	}

	if p.Src != "" {
		return h.Img(render.Attrs(
			p.ID,
			tw.Join(size, "rounded-full border object-cover", p.Class),
			p.DataTestID,
			public.Props{},
			p.Attributes,
			h.Src(p.Src),
			h.Alt(p.Alt),
		)...)
	}

	initials := "?"
	if trimmed := strings.TrimSpace(p.Name); trimmed != "" {
		parts := strings.Fields(trimmed)
		initials = strings.ToUpper(string(parts[0][0]))
		if len(parts) > 1 {
			initials += strings.ToUpper(string(parts[len(parts)-1][0]))
		}
	}

	return h.Div(
		append(
			render.Attrs(
				p.ID,
				tw.Join(size, "inline-flex items-center justify-center rounded-full border bg-[rgb(var(--ui-surface-strong))] font-semibold text-[rgb(var(--ui-muted-foreground))]", p.Class),
				p.DataTestID,
				public.Props{},
				p.Attributes,
				h.Aria("label", p.Name),
			),
			g.Text(initials),
		)...,
	)
}
