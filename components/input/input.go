package input

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	public "github.com/pmenglund/goth/htmx"
	"github.com/pmenglund/goth/internal/render"
	"github.com/pmenglund/goth/internal/tw"
)

type Props struct {
	ID           string
	Class        string
	Attributes   []g.Node
	DataTestID   string
	Name         string
	Type         string
	Value        string
	Placeholder  string
	Disabled     bool
	Required     bool
	Invalid      bool
	AutoComplete string
	DescribedBy  string
	HTMX         public.Props
}

// Input renders a single-line text input with the provided attributes.
func Input(p Props) g.Node {
	inputType := p.Type
	if inputType == "" {
		inputType = "text"
	}

	className := tw.Join(
		"h-11 w-full rounded-[var(--ui-radius)] border bg-[rgb(var(--ui-surface))] px-3 text-sm text-[rgb(var(--ui-foreground))] shadow-sm transition focus:outline-none focus-visible:ring-2 focus-visible:ring-[rgb(var(--ui-ring))] focus-visible:ring-offset-2 focus-visible:ring-offset-[rgb(var(--ui-background))]",
		tw.When(p.Disabled, "cursor-not-allowed opacity-60"),
		tw.When(p.Invalid, "border-[rgb(var(--ui-danger))] focus-visible:ring-[rgb(var(--ui-danger))]"),
		p.Class,
	)

	attrs := render.Attrs(
		p.ID,
		className,
		p.DataTestID,
		p.HTMX,
		p.Attributes,
		h.Type(inputType),
		g.If(p.Name != "", h.Name(p.Name)),
		g.If(p.Value != "", h.Value(p.Value)),
		g.If(p.Placeholder != "", h.Placeholder(p.Placeholder)),
		g.If(p.AutoComplete != "", h.AutoComplete(p.AutoComplete)),
		g.If(p.DescribedBy != "", h.Aria("describedby", p.DescribedBy)),
		g.If(p.Invalid, h.Aria("invalid", "true")),
		g.If(p.Disabled, h.Disabled()),
		g.If(p.Required, h.Required()),
	)
	return h.Input(attrs...)
}
