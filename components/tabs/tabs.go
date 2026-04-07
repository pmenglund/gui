package tabs

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	public "github.com/pmenglund/gui/htmx"
	"github.com/pmenglund/gui/internal/a11y"
	"github.com/pmenglund/gui/internal/render"
	"github.com/pmenglund/gui/internal/tw"
)

type Item struct {
	Key   string
	Label string
	Panel g.Node
}

type Props struct {
	ID         string
	Class      string
	Attributes []g.Node
	DataTestID string
	Items      []Item
	Value      string
	HTMX       public.Props
}

type normalizedItem struct {
	key   string
	label string
	panel g.Node
}

// Tabs renders a tab list and associated panels from the provided items.
func Tabs(p Props) g.Node {
	items := make([]normalizedItem, 0, len(p.Items))
	for _, item := range p.Items {
		key := item.Key
		if key == "" {
			key = a11y.Slug(item.Label)
		}
		items = append(items, normalizedItem{
			key:   key,
			label: item.Label,
			panel: item.Panel,
		})
	}

	value := p.Value
	if len(items) > 0 && !hasValue(items, value) {
		value = items[0].key
	}

	triggers := make([]g.Node, 0, len(items))
	panels := make([]g.Node, 0, len(items))
	rootID := a11y.ID("tabs", p.ID, p.DataTestID)

	for _, item := range items {
		panelID := rootID + "-panel-" + item.key
		triggerID := rootID + "-trigger-" + item.key
		current := item.key == value

		triggers = append(triggers, h.Button(
			h.Type("button"),
			h.ID(triggerID),
			h.Class(tw.Join(
				"rounded-full px-3 py-2 text-sm font-medium transition hover:bg-[rgb(var(--ui-surface-strong))]",
				tw.When(current, "bg-[rgb(var(--ui-primary))] text-[rgb(var(--ui-primary-foreground))]"),
			)),
			h.Role("tab"),
			h.Aria("controls", panelID),
			h.Aria("selected", boolString(current)),
			h.TabIndex(tabIndex(current)),
			h.Data("ui-trigger", ""),
			h.Data("ui-target", item.key),
			g.Text(item.label),
		))

		panels = append(panels, h.Div(
			h.ID(panelID),
			h.Class("rounded-[var(--ui-radius)] border bg-[rgb(var(--ui-surface))] p-4"),
			h.Role("tabpanel"),
			h.Aria("labelledby", triggerID),
			g.If(!current, h.Hidden("hidden")),
			h.Data("ui-content", ""),
			h.Data("ui-target", item.key),
			item.panel,
		))
	}

	return h.Div(
		append(
			render.Attrs(
				rootID,
				tw.Join("grid gap-4", p.Class),
				p.DataTestID,
				p.HTMX,
				p.Attributes,
				h.Data("ui-controller", "tabs"),
				h.Data("ui-value", value),
			),
			h.Div(h.Class("inline-flex flex-wrap gap-2 rounded-full bg-[rgb(var(--ui-surface-strong))] p-1"), h.Role("tablist"), g.Group(triggers)),
			h.Div(h.Class("grid gap-4"), g.Group(panels)),
		)...,
	)
}

func hasValue(items []normalizedItem, value string) bool {
	for _, item := range items {
		if item.key == value {
			return true
		}
	}
	return false
}

func boolString(ok bool) string {
	if ok {
		return "true"
	}
	return "false"
}

func tabIndex(current bool) string {
	if current {
		return "0"
	}
	return "-1"
}
