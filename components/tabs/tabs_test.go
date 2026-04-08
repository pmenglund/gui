package tabs_test

import (
	"strings"
	"testing"

	g "maragu.dev/gomponents"

	"github.com/pmenglund/goth/components/tabs"
	"github.com/pmenglund/goth/internal/testutil"
)

func TestTabsRenderInitialSelectionState(t *testing.T) {
	t.Parallel()

	got := testutil.Render(t, tabs.Tabs(tabs.Props{
		ID:    "release-tabs",
		Value: "activity",
		Items: []tabs.Item{
			{Key: "overview", Label: "Overview", Panel: g.Text("Overview panel")},
			{Key: "activity", Label: "Activity", Panel: g.Text("Activity panel")},
		},
	}))

	if !strings.Contains(got, `id="release-tabs-trigger-overview"`) || !strings.Contains(got, `aria-selected="false" tabindex="-1"`) {
		t.Fatalf("inactive tab trigger missing initial accessibility state: %s", got)
	}
	if !strings.Contains(got, `id="release-tabs-trigger-activity"`) || !strings.Contains(got, `aria-selected="true" tabindex="0"`) {
		t.Fatalf("active tab trigger missing initial accessibility state: %s", got)
	}
	if !strings.Contains(got, `id="release-tabs-panel-overview" class="rounded-[var(--ui-radius)] border bg-[rgb(var(--ui-surface))] p-4 hidden" role="tabpanel" aria-labelledby="release-tabs-trigger-overview" hidden`) {
		t.Fatalf("inactive tab panel missing initial hidden state: %s", got)
	}
}
