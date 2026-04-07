package render

import (
	"strings"
	"unicode"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	public "github.com/pmenglund/gui/htmx"
	ihtmx "github.com/pmenglund/gui/internal/htmx"
)

type typedNode interface {
	Type() g.NodeType
}

// Attrs merges the common component attributes in stable order while reserving
// component-managed attributes and delegated runtime hooks from raw passthrough.
func Attrs(id, className, dataTestID string, hx public.Props, attributes []g.Node, extra ...g.Node) []g.Node {
	internal := make([]g.Node, 0, len(extra)+4)
	if id != "" {
		internal = append(internal, h.ID(id))
	}
	if className != "" {
		internal = append(internal, h.Class(className))
	}
	if dataTestID != "" {
		internal = append(internal, h.Data("testid", dataTestID))
	}
	internal = append(internal, ihtmx.Attrs(hx)...)
	internal = append(internal, flattenNodes(extra)...)

	passthroughAttrs, passthroughOther := splitPassthroughNodes(attributes)
	protected := protectedAttributeNames(internal)
	filtered := filterPassthroughAttributes(passthroughAttrs, protected)

	out := make([]g.Node, 0, len(internal)+len(filtered)+len(passthroughOther))
	out = append(out, internal...)
	out = append(out, filtered...)
	out = append(out, passthroughOther...)
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

func flattenNodes(nodes []g.Node) []g.Node {
	out := make([]g.Node, 0, len(nodes))
	for _, node := range nodes {
		appendFlattened(&out, node)
	}
	return out
}

func appendFlattened(out *[]g.Node, node g.Node) {
	if node == nil {
		return
	}
	if group, ok := node.(g.Group); ok {
		for _, child := range group {
			appendFlattened(out, child)
		}
		return
	}
	*out = append(*out, node)
}

func splitPassthroughNodes(nodes []g.Node) (attrNodes []g.Node, otherNodes []g.Node) {
	flat := flattenNodes(nodes)
	attrNodes = make([]g.Node, 0, len(flat))
	otherNodes = make([]g.Node, 0, len(flat))
	for _, node := range flat {
		if isAttributeNode(node) {
			attrNodes = append(attrNodes, node)
			continue
		}
		otherNodes = append(otherNodes, node)
	}
	return attrNodes, otherNodes
}

func protectedAttributeNames(nodes []g.Node) map[string]struct{} {
	protected := make(map[string]struct{}, len(nodes))
	for _, node := range flattenNodes(nodes) {
		if !isAttributeNode(node) {
			continue
		}
		name, ok := extractAttributeName(node)
		if !ok {
			continue
		}
		protected[name] = struct{}{}
	}
	return protected
}

func filterPassthroughAttributes(nodes []g.Node, protected map[string]struct{}) []g.Node {
	out := make([]g.Node, 0, len(nodes))
	for _, node := range nodes {
		name, ok := extractAttributeName(node)
		if !ok {
			out = append(out, node)
			continue
		}
		if strings.HasPrefix(name, "data-ui-") {
			continue
		}
		if _, exists := protected[name]; exists {
			continue
		}
		out = append(out, node)
	}
	return out
}

func isAttributeNode(node g.Node) bool {
	typed, ok := node.(typedNode)
	return ok && typed.Type() == g.AttributeType
}

func extractAttributeName(node g.Node) (string, bool) {
	var b strings.Builder
	if err := node.Render(&b); err != nil {
		return "", false
	}

	rendered := strings.TrimSpace(b.String())
	if rendered == "" {
		return "", false
	}

	end := strings.IndexFunc(rendered, func(r rune) bool {
		return unicode.IsSpace(r) || r == '='
	})
	if end == 0 {
		return "", false
	}
	if end == -1 {
		return strings.ToLower(rendered), true
	}
	return strings.ToLower(rendered[:end]), true
}
