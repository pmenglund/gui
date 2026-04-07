package selectui

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	public "github.com/pmenglund/gui/htmx"
	"github.com/pmenglund/gui/internal/render"
	"github.com/pmenglund/gui/internal/tw"
)

type Option struct {
	Value    string
	Label    string
	Disabled bool
}

type Props struct {
	ID          string
	Class       string
	Attributes  []g.Node
	DataTestID  string
	Name        string
	Value       string
	Placeholder string
	Options     []Option
	Disabled    bool
	Required    bool
	Invalid     bool
	DescribedBy string
	HTMX        public.Props
}

// Select renders a native select element with the provided options.
func Select(p Props) g.Node {
	children := make([]g.Node, 0, len(p.Options)+1)
	if p.Placeholder != "" {
		children = append(children, h.Option(h.Value(""), g.Text(p.Placeholder)))
	}
	for _, option := range p.Options {
		children = append(children, h.Option(
			h.Value(option.Value),
			g.If(option.Disabled, h.Disabled()),
			g.If(option.Value == p.Value, h.Selected()),
			g.Text(option.Label),
		))
	}

	return h.Select(append(
		render.Attrs(
			p.ID,
			tw.Join(
				"h-11 w-full rounded-[var(--ui-radius)] border bg-[rgb(var(--ui-surface))] px-3 text-sm text-[rgb(var(--ui-foreground))] shadow-sm focus:outline-none focus-visible:ring-2 focus-visible:ring-[rgb(var(--ui-ring))] focus-visible:ring-offset-2 focus-visible:ring-offset-[rgb(var(--ui-background))]",
				tw.When(p.Invalid, "border-[rgb(var(--ui-danger))]"),
				tw.When(p.Disabled, "cursor-not-allowed opacity-60"),
				p.Class,
			),
			p.DataTestID,
			p.HTMX,
			p.Attributes,
			g.If(p.Name != "", h.Name(p.Name)),
			g.If(p.Disabled, h.Disabled()),
			g.If(p.Required, h.Required()),
			g.If(p.Invalid, h.Aria("invalid", "true")),
			g.If(p.DescribedBy != "", h.Aria("describedby", p.DescribedBy)),
		),
		children...,
	)...)
}
