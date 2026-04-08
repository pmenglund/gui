package fielderror

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
}

// FieldError renders an inline validation message for a form field.
func FieldError(p Props, children ...g.Node) g.Node {
	return h.P(append(
		render.Attrs(
			p.ID,
			tw.Join("text-sm font-medium text-[rgb(var(--ui-danger))]", p.Class),
			p.DataTestID,
			public.Props{},
			p.Attributes,
		),
		children...,
	)...)
}
