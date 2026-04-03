package render

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	public "github.com/pmenglund/gui/htmx"
	ihtmx "github.com/pmenglund/gui/internal/htmx"
)

// Attrs merges the common component attributes in stable order.
func Attrs(id, className, dataTestID string, hx public.Props, attributes []g.Node, extra ...g.Node) []g.Node {
	out := make([]g.Node, 0, len(attributes)+len(extra)+4)
	if id != "" {
		out = append(out, h.ID(id))
	}
	if className != "" {
		out = append(out, h.Class(className))
	}
	if dataTestID != "" {
		out = append(out, h.Data("testid", dataTestID))
	}
	out = append(out, ihtmx.Attrs(hx)...)
	out = append(out, extra...)
	out = append(out, attributes...)
	return out
}

// Nodes returns the non-nil nodes in order.
func Nodes(nodes ...g.Node) []g.Node {
	out := make([]g.Node, 0, len(nodes))
	for _, node := range nodes {
		if node != nil {
			out = append(out, node)
		}
	}
	return out
}
