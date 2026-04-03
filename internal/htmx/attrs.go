package htmx

import (
	g "maragu.dev/gomponents"

	public "github.com/pmenglund/gui/htmx"
)

// Attrs converts stable HTMX props into gomponents attributes.
func Attrs(p public.Props) []g.Node {
	if p.Disabled {
		return nil
	}
	return filter(
		attr("hx-get", p.Get),
		attr("hx-post", p.Post),
		attr("hx-put", p.Put),
		attr("hx-patch", p.Patch),
		attr("hx-delete", p.Delete),
		attr("hx-trigger", p.Trigger),
		attr("hx-target", p.Target),
		attr("hx-swap", p.Swap),
		attr("hx-select", p.Select),
		attr("hx-include", p.Include),
		attr("hx-indicator", p.Indicator),
		attr("hx-push-url", p.PushURL),
		attr("hx-confirm", p.Confirm),
		attr("hx-encoding", p.Encoding),
		attr("hx-vals", p.Values),
	)
}

func attr(name, value string) g.Node {
	if value == "" {
		return nil
	}
	return g.Attr(name, value)
}

func filter(nodes ...g.Node) []g.Node {
	out := make([]g.Node, 0, len(nodes))
	for _, node := range nodes {
		if node != nil {
			out = append(out, node)
		}
	}
	return out
}
