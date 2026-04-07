package radiogroup

import (
	"fmt"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	public "github.com/pmenglund/gui/htmx"
	"github.com/pmenglund/gui/internal/a11y"
	"github.com/pmenglund/gui/internal/render"
	"github.com/pmenglund/gui/internal/tw"
)

type Option struct {
	Value       string
	Label       string
	Description string
	Disabled    bool
}

type Props struct {
	ID          string
	Class       string
	Attributes  []g.Node
	DataTestID  string
	Name        string
	Legend      string
	Value       string
	Options     []Option
	Disabled    bool
	Required    bool
	Invalid     bool
	DescribedBy string
	HTMX        public.Props
}

// RadioGroup renders a fieldset of mutually exclusive options.
func RadioGroup(p Props) g.Node {
	items := make([]g.Node, 0, len(p.Options)+1)
	groupID := p.ID
	if groupID == "" {
		groupID = a11y.InstanceID("radiogroup", "", p.Name, p.Legend, p.DataTestID)
	}
	if p.Legend != "" {
		items = append(items, h.Legend(h.Class("mb-2 text-sm font-medium"), g.Text(p.Legend)))
	}

	for i, option := range p.Options {
		optionID := fmt.Sprintf("%s-option-%d", groupID, i+1)
		input := h.Input(render.Attrs(
			optionID,
			tw.Join("mt-1 h-4 w-4 border text-[rgb(var(--ui-primary))] focus:outline-none focus-visible:ring-2 focus-visible:ring-[rgb(var(--ui-ring))]"),
			"",
			p.HTMX,
			nil,
			h.Type("radio"),
			h.Name(p.Name),
			h.Value(option.Value),
			g.If(option.Value == p.Value, h.Checked()),
			g.If(option.Disabled || p.Disabled, h.Disabled()),
			g.If(p.Required, h.Required()),
			g.If(p.Invalid, h.Aria("invalid", "true")),
			g.If(p.DescribedBy != "", h.Aria("describedby", p.DescribedBy)),
		)...)

		copy := option
		items = append(items, h.Label(
			h.Class("flex gap-3 rounded-[var(--ui-radius)] border border-[rgb(var(--ui-border))] bg-[rgb(var(--ui-surface))] px-3 py-3"),
			input,
			h.Span(
				h.Class("grid gap-1"),
				h.Span(h.Class("text-sm font-medium"), g.Text(copy.Label)),
				g.If(copy.Description != "", h.Span(h.Class("text-sm text-[rgb(var(--ui-muted-foreground))]"), g.Text(copy.Description))),
			),
		))
	}

	return h.FieldSet(append(
		render.Attrs(
			p.ID,
			tw.Join("grid gap-3", tw.When(p.Invalid, "text-[rgb(var(--ui-danger))]"), p.Class),
			p.DataTestID,
			public.Props{},
			p.Attributes,
			g.If(p.DescribedBy != "", h.Aria("describedby", p.DescribedBy)),
		),
		items...,
	)...)
}
