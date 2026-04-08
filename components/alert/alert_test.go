package alert_test

import (
	"strings"
	"testing"

	"github.com/pmenglund/goth/components/alert"
	"github.com/pmenglund/goth/internal/testutil"
)

func TestAlertSlotClassesRenderOnGeneratedTitleAndDescription(t *testing.T) {
	t.Parallel()

	got := testutil.Render(t, alert.Alert(alert.Props{
		Title:            "Heads up",
		Description:      "Check the queue.",
		TitleClass:       "host-title",
		DescriptionClass: "host-description",
	}))
	if !strings.Contains(got, `<h4 class="font-semibold host-title">Heads up</h4>`) {
		t.Fatalf("alert title missing slot class: %s", got)
	}
	if !strings.Contains(got, `<p class="text-sm text-[rgb(var(--ui-muted-foreground))] host-description">Check the queue.</p>`) {
		t.Fatalf("alert description missing slot class: %s", got)
	}
}
