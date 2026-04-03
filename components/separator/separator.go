package separator

import (
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
	Vertical   bool
	HTMX       public.Props
}

func Separator(p Props) g.Node {
	className := "my-2 h-px w-full bg-[rgb(var(--ui-border))]"
	if p.Vertical {
		className = "mx-2 h-full min-h-6 w-px bg-[rgb(var(--ui-border))]"
	}
	return h.Hr(render.Attrs(
		p.ID,
		tw.Join(className, p.Class),
		p.DataTestID,
		p.HTMX,
		p.Attributes,
		h.Role("separator"),
	)...)
}
