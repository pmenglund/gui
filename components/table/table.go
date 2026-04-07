package table

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	public "github.com/pmenglund/gui/htmx"
	"github.com/pmenglund/gui/internal/render"
	"github.com/pmenglund/gui/internal/tw"
)

type Column struct {
	Header string
	Class  string
}

type Row struct {
	Cells []g.Node
}

type Props struct {
	ID         string
	Class      string
	Attributes []g.Node
	DataTestID string
	Columns    []Column
	Rows       []Row
	Caption    string
	HTMX       public.Props
}

// Table renders tabular data from the provided columns and rows.
func Table(p Props) g.Node {
	headers := make([]g.Node, 0, len(p.Columns))
	for _, column := range p.Columns {
		headers = append(headers, h.Th(h.Class(tw.Join("px-4 py-3 text-left text-xs font-semibold uppercase tracking-[0.18em] text-[rgb(var(--ui-muted-foreground))]", column.Class)), g.Text(column.Header)))
	}

	rows := make([]g.Node, 0, len(p.Rows))
	for _, row := range p.Rows {
		cells := make([]g.Node, 0, len(row.Cells))
		for _, cell := range row.Cells {
			cells = append(cells, h.Td(h.Class("px-4 py-3 text-sm"), cell))
		}
		rows = append(rows, h.Tr(h.Class("border-t"), g.Group(cells)))
	}

	table := h.Table(
		h.Class("min-w-full border-separate border-spacing-0"),
		g.If(p.Caption != "", h.Caption(h.Class("sr-only"), g.Text(p.Caption))),
		h.THead(h.Tr(h.Class("bg-[rgb(var(--ui-surface-strong))]"), g.Group(headers))),
		h.TBody(g.Group(rows)),
	)

	return h.Div(
		append(
			render.Attrs(
				p.ID,
				tw.Join("overflow-hidden rounded-[var(--ui-radius)] border bg-[rgb(var(--ui-surface))]", p.Class),
				p.DataTestID,
				p.HTMX,
				p.Attributes,
			),
			table,
		)...,
	)
}
