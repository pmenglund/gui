package skeleton

import (
	"strconv"

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
	Width      string
	Height     string
	Circle     bool
	Count      int
}

func Skeleton(p Props) g.Node {
	count := p.Count
	if count == 0 {
		count = 1
	}

	items := make([]g.Node, 0, count)
	for i := 0; i < count; i++ {
		items = append(items, h.Div(
			render.Attrs(
				"",
				tw.Join(
					"animate-pulse bg-[rgb(var(--ui-muted))]",
					defaultString(p.Width, "w-full"),
					defaultString(p.Height, "h-4"),
					tw.When(p.Circle, "rounded-full"),
					tw.When(!p.Circle, "rounded-[calc(var(--ui-radius)*0.75)]"),
				),
				"",
				public.Props{},
				nil,
				g.Attr("aria-hidden", "true"),
			)...,
		))
	}

	return h.Div(append(
		render.Attrs(
			p.ID,
			tw.Join("grid gap-2", p.Class),
			p.DataTestID,
			public.Props{},
			p.Attributes,
			h.Data("count", strconv.Itoa(count)),
		),
		items...,
	)...)
}

func defaultString(value, fallback string) string {
	if value == "" {
		return fallback
	}
	return value
}
