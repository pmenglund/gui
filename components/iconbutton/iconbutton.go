package iconbutton

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/pmenglund/goth/components/button"
	public "github.com/pmenglund/goth/htmx"
)

type Props struct {
	ID         string
	Class      string
	Attributes []g.Node
	Variant    button.Variant
	Size       button.Size
	Type       string
	Disabled   bool
	DataTestID string
	Label      string
	Icon       g.Node
	HTMX       public.Props
}

// IconButton renders a button optimized for icon-only content.
func IconButton(p Props, children ...g.Node) g.Node {
	icon := p.Icon
	if len(children) > 0 {
		icon = g.Group(children)
	}
	return button.Button(button.Props{
		ID:         p.ID,
		Class:      p.Class,
		Attributes: append([]g.Node{h.Aria("label", p.Label)}, p.Attributes...),
		Variant:    p.Variant,
		Size:       p.Size,
		Type:       p.Type,
		Disabled:   p.Disabled,
		DataTestID: p.DataTestID,
		HTMX:       p.HTMX,
	}, icon)
}
