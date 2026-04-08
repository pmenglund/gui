package table

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/pmenglund/goth/components/classmode"
	public "github.com/pmenglund/goth/htmx"
	"github.com/pmenglund/goth/internal/render"
	"github.com/pmenglund/goth/internal/tw"
)

// Column configures a table header cell.
type Column struct {
	Header     string
	Class      string
	Attributes []g.Node
	DataTestID string
}

// Cell configures a table body cell.
type Cell struct {
	Content    g.Node
	Class      string
	Attributes []g.Node
	DataTestID string
}

// Props configures Table rendering.
type Props struct {
	ID              string
	Class           string
	ClassMode       classmode.ClassMode
	Attributes      []g.Node
	DataTestID      string
	Columns         []Column
	Rows            []Row
	Caption         string
	TableClass      string
	HeadClass       string
	HeaderRowClass  string
	HeaderCellClass string
	BodyClass       string
	RowClass        string
	CellClass       string
	HTMX            public.Props
}

// Row configures a table body row.
type Row struct {
	Cells      []g.Node
	CellItems  []Cell
	Class      string
	Attributes []g.Node
	DataTestID string
}

// Table renders tabular data from the provided columns and rows.
func Table(p Props) g.Node {
	headers := make([]g.Node, 0, len(p.Columns))
	for _, column := range p.Columns {
		headers = append(headers, h.Th(append(
			render.Attrs(
				"",
				tw.Join("px-4 py-3 text-left text-xs font-semibold uppercase tracking-[0.18em] text-[rgb(var(--ui-muted-foreground))]", p.HeaderCellClass, column.Class),
				column.DataTestID,
				public.Props{},
				column.Attributes,
			),
			g.Text(column.Header),
		)...))
	}

	rows := make([]g.Node, 0, len(p.Rows))
	for _, row := range p.Rows {
		var cells []g.Node
		if len(row.CellItems) > 0 {
			cells = make([]g.Node, 0, len(row.CellItems))
			for _, cell := range row.CellItems {
				cells = append(cells, h.Td(append(
					render.Attrs(
						"",
						tw.Join("px-4 py-3 text-sm", p.CellClass, cell.Class),
						cell.DataTestID,
						public.Props{},
						cell.Attributes,
					),
					cell.Content,
				)...))
			}
		} else {
			cells = make([]g.Node, 0, len(row.Cells))
			for _, cell := range row.Cells {
				cells = append(cells, h.Td(h.Class(tw.Join("px-4 py-3 text-sm", p.CellClass)), cell))
			}
		}
		rows = append(rows, h.Tr(append(
			render.Attrs(
				"",
				tw.Join("border-t", p.RowClass, row.Class),
				row.DataTestID,
				public.Props{},
				row.Attributes,
			),
			g.Group(cells),
		)...))
	}

	table := h.Table(append(
		render.Attrs(
			"",
			tw.Join("min-w-full border-separate border-spacing-0", p.TableClass),
			"",
			public.Props{},
			nil,
		),
		g.If(p.Caption != "", h.Caption(h.Class("sr-only"), g.Text(p.Caption))),
		h.THead(append(
			render.Attrs("", p.HeadClass, "", public.Props{}, nil),
			h.Tr(h.Class(tw.Join("bg-[rgb(var(--ui-surface-strong))]", p.HeaderRowClass)), g.Group(headers)),
		)...),
		h.TBody(append(
			render.Attrs("", p.BodyClass, "", public.Props{}, nil),
			g.Group(rows),
		)...),
	)...)

	return h.Div(
		append(
			render.Attrs(
				p.ID,
				tw.Classes(p.ClassMode, "overflow-hidden rounded-[var(--ui-radius)] border bg-[rgb(var(--ui-surface))]", p.Class),
				p.DataTestID,
				p.HTMX,
				p.Attributes,
			),
			table,
		)...,
	)
}
