package fielddescription

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

// FieldDescription renders helper text associated with a form field.
func FieldDescription(p Props, children ...g.Node) g.Node {
	return h.P(append(
		render.Attrs(
			p.ID,
			tw.Join("text-sm text-[rgb(var(--ui-muted-foreground))]", p.Class),
			p.DataTestID,
			public.Props{},
			p.Attributes,
		),
		children...,
	)...)
}
