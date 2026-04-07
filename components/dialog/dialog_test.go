package dialog_test

import (
	"strings"
	"testing"

	"github.com/pmenglund/gui/components/dialog"
	"github.com/pmenglund/gui/internal/testutil"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestDialogPreservesProtectedRuntimeAndARIAAttrs(t *testing.T) {
	t.Parallel()

	got := testutil.Render(t, dialog.Dialog(dialog.Props{
		Title: "Review rollout",
		Attributes: []g.Node{
			h.ID("bad-id"),
			h.Class("bad-class"),
			g.Attr("data-ui-controller", "bad"),
			g.Attr("data-ui-state", "bad"),
			g.Attr("title", "extra"),
		},
	}, g.Text("Body copy")))

	for _, want := range []string{
		`id="dialog-review-rollout"`,
		`class="inline-flex"`,
		`data-ui-controller="dialog"`,
		`data-ui-state="closed"`,
		`aria-labelledby="dialog-review-rollout-title"`,
		`title="extra"`,
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("rendered dialog missing %s: %s", want, got)
		}
	}

	for _, blocked := range []string{
		`id="bad-id"`,
		`class="bad-class"`,
		`data-ui-controller="bad"`,
		`data-ui-state="bad"`,
	} {
		if strings.Contains(got, blocked) {
			t.Fatalf("rendered dialog unexpectedly contained %s: %s", blocked, got)
		}
	}
}
