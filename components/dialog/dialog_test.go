package dialog_test

import (
	"strings"
	"testing"

	g "maragu.dev/gomponents"

	"github.com/pmenglund/goth/components/dialog"
	"github.com/pmenglund/goth/internal/testutil"
)

func TestDialogOpenRendersVisibleInitialContent(t *testing.T) {
	t.Parallel()

	got := testutil.Render(t, dialog.Dialog(dialog.Props{
		ID:    "review-dialog",
		Title: "Review rollout",
		Open:  true,
	}, g.Text("Ready to review.")))

	if strings.Contains(got, `data-ui-content="" class="hidden fixed`) {
		t.Fatalf("open dialog rendered hidden utility class on content: %s", got)
	}
	if strings.Contains(got, `data-ui-content="" class="fixed inset-0 z-50 grid place-items-center bg-black/50 p-4" hidden`) {
		t.Fatalf("open dialog rendered hidden attribute on content: %s", got)
	}
	if !strings.Contains(got, `data-ui-content="" class="fixed inset-0 z-50 grid place-items-center bg-black/50 p-4"`) {
		t.Fatalf("open dialog missing visible content class: %s", got)
	}
}
