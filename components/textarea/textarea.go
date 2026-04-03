package textarea

import (
	"strconv"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	public "github.com/pmenglund/gui/htmx"
	"github.com/pmenglund/gui/internal/render"
	"github.com/pmenglund/gui/internal/tw"
)

type Props struct {
	ID          string
	Class       string
	Attributes  []g.Node
	DataTestID  string
	Name        string
	Value       string
	Placeholder string
	Rows        int
	Disabled    bool
	Required    bool
	Invalid     bool
	DescribedBy string
	HTMX        public.Props
}

func Textarea(p Props) g.Node {
	rows := p.Rows
	if rows == 0 {
		rows = 4
	}

	attrs := render.Attrs(
		p.ID,
		tw.Join(
			"w-full rounded-[var(--ui-radius)] border bg-[rgb(var(--ui-surface))] px-3 py-2 text-sm text-[rgb(var(--ui-foreground))] shadow-sm transition focus:outline-none focus-visible:ring-2 focus-visible:ring-[rgb(var(--ui-ring))] focus-visible:ring-offset-2 focus-visible:ring-offset-[rgb(var(--ui-background))]",
			tw.When(p.Disabled, "cursor-not-allowed opacity-60"),
			tw.When(p.Invalid, "border-[rgb(var(--ui-danger))] focus-visible:ring-[rgb(var(--ui-danger))]"),
			p.Class,
		),
		p.DataTestID,
		p.HTMX,
		p.Attributes,
		g.If(p.Name != "", h.Name(p.Name)),
		h.Rows(strconv.Itoa(rows)),
		g.If(p.Placeholder != "", h.Placeholder(p.Placeholder)),
		g.If(p.DescribedBy != "", h.Aria("describedby", p.DescribedBy)),
		g.If(p.Invalid, h.Aria("invalid", "true")),
		g.If(p.Disabled, h.Disabled()),
		g.If(p.Required, h.Required()),
	)
	return h.Textarea(append(attrs, g.Text(p.Value))...)
}
