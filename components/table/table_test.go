package table_test

import (
	"strings"
	"testing"

	g "maragu.dev/gomponents"

	"github.com/pmenglund/goth/components/table"
	"github.com/pmenglund/goth/internal/testutil"
)

func TestTableRowDataTestIDAndAttributesRenderOnTR(t *testing.T) {
	t.Parallel()

	got := testutil.Render(t, table.Table(table.Props{
		Columns: []table.Column{{Header: "Service"}},
		Rows: []table.Row{{
			DataTestID: "state-issue-review-COLIN-94",
			Attributes: []g.Node{
				g.Attr("aria-selected", "true"),
			},
			Cells: []g.Node{g.Text("API")},
		}},
	}))
	if !strings.Contains(got, `<tr class="border-t" data-testid="state-issue-review-COLIN-94" aria-selected="true">`) {
		t.Fatalf("table row missing tr hooks: %s", got)
	}
}

func TestTableCellItemsRenderCellHooks(t *testing.T) {
	t.Parallel()

	got := testutil.Render(t, table.Table(table.Props{
		CellClass: "host-cell",
		Columns:   []table.Column{{Header: "Service"}},
		Rows: []table.Row{{
			CellItems: []table.Cell{{
				Content:    g.Text("API"),
				Class:      "service-cell",
				DataTestID: "service-cell",
				Attributes: []g.Node{
					g.Attr("data-state", "healthy"),
				},
			}},
		}},
	}))
	if !strings.Contains(got, `<td class="px-4 py-3 text-sm host-cell service-cell" data-testid="service-cell" data-state="healthy">API</td>`) {
		t.Fatalf("table CellItems missing td hooks: %s", got)
	}
}

func TestTableLegacyCellsRenderUnchanged(t *testing.T) {
	t.Parallel()

	got := testutil.Render(t, table.Table(table.Props{
		Columns: []table.Column{{Header: "Service"}},
		Rows: []table.Row{{
			Cells: []g.Node{g.Text("API")},
		}},
	}))
	want := `<div class="overflow-hidden rounded-[var(--ui-radius)] border bg-[rgb(var(--ui-surface))]"><table class="min-w-full border-separate border-spacing-0"><thead><tr class="bg-[rgb(var(--ui-surface-strong))]"><th class="px-4 py-3 text-left text-xs font-semibold uppercase tracking-[0.18em] text-[rgb(var(--ui-muted-foreground))]">Service</th></tr></thead><tbody><tr class="border-t"><td class="px-4 py-3 text-sm">API</td></tr></tbody></table></div>`
	if got != want {
		t.Fatalf("legacy table cells changed\nwant:\n%s\n\ngot:\n%s", want, got)
	}
}
