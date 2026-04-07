package tabs_test

import (
	"strings"
	"testing"

	"github.com/pmenglund/gui/components/tabs"
	"github.com/pmenglund/gui/internal/testutil"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestTabsRenderServerState(t *testing.T) {
	t.Parallel()

	got := testutil.Render(t, tabs.Tabs(tabs.Props{
		ID: "server-state",
		Items: []tabs.Item{
			{Label: "Overview", Panel: h.P(g.Text("Overview panel"))},
			{Label: "Activity", Panel: h.P(g.Text("Activity panel"))},
		},
	}))

	checks := []string{
		`data-ui-value="overview"`,
		`id="server-state-trigger-overview"`,
		`aria-controls="server-state-panel-overview" aria-selected="true" tabindex="0" data-ui-trigger="" data-ui-target="overview"`,
		`id="server-state-trigger-activity"`,
		`aria-controls="server-state-panel-activity" aria-selected="false" tabindex="-1" data-ui-trigger="" data-ui-target="activity"`,
		`id="server-state-panel-overview" class="rounded-[var(--ui-radius)] border bg-[rgb(var(--ui-surface))] p-4" role="tabpanel" aria-labelledby="server-state-trigger-overview" data-ui-content="" data-ui-target="overview"`,
		`id="server-state-panel-activity" class="rounded-[var(--ui-radius)] border bg-[rgb(var(--ui-surface))] p-4" role="tabpanel" aria-labelledby="server-state-trigger-activity" hidden="hidden" data-ui-content="" data-ui-target="activity"`,
	}

	for _, check := range checks {
		if !strings.Contains(got, check) {
			t.Fatalf("tabs markup missing %q: %s", check, got)
		}
	}
}
