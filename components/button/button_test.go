package button_test

import (
	"strings"
	"testing"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/pmenglund/goth/components/button"
	"github.com/pmenglund/goth/components/classmode"
	public "github.com/pmenglund/goth/htmx"
	"github.com/pmenglund/goth/internal/testutil"
)

func TestButtonRootClassMergesByDefault(t *testing.T) {
	t.Parallel()

	got := testutil.Render(t, button.Button(button.Props{Class: "host-button"}, g.Text("Save")))
	if !strings.Contains(got, `class="inline-flex`) {
		t.Fatalf("button missing default root classes: %s", got)
	}
	if !strings.Contains(got, `host-button`) {
		t.Fatalf("button missing caller root class: %s", got)
	}
}

func TestButtonRootClassReplaceUsesOnlyCallerClass(t *testing.T) {
	t.Parallel()

	got := testutil.Render(t, button.Button(button.Props{
		ClassMode: classmode.ClassReplace,
		Class:     "host-button",
	}, g.Text("Save")))
	if strings.Contains(got, `inline-flex`) || strings.Contains(got, `rounded-[var(--ui-radius)]`) {
		t.Fatalf("button class replacement kept default classes: %s", got)
	}
	if !strings.Contains(got, `class="host-button"`) {
		t.Fatalf("button missing replacement class: %s", got)
	}
}

func TestButtonRootAttributesKeepStableOrder(t *testing.T) {
	t.Parallel()

	got := testutil.Render(t, button.Button(button.Props{
		ID:         "save",
		Class:      "host-button",
		DataTestID: "save-button",
		Type:       "submit",
		HTMX: public.Props{
			Get:    "/save",
			Target: "#panel",
		},
		Attributes: []g.Node{h.Aria("label", "Save settings")},
	}, g.Text("Save")))

	want := `<button id="save" class="inline-flex items-center justify-center gap-2 rounded-[var(--ui-radius)] border text-sm font-medium transition focus:outline-none focus-visible:ring-2 focus-visible:ring-[rgb(var(--ui-ring))] focus-visible:ring-offset-2 focus-visible:ring-offset-[rgb(var(--ui-background))] border-[rgb(var(--ui-primary))] bg-[rgb(var(--ui-primary))] text-[rgb(var(--ui-primary-foreground))] hover:opacity-90 h-11 px-4 host-button" data-testid="save-button" hx-get="/save" hx-target="#panel" type="submit" aria-label="Save settings">Save</button>`
	if got != want {
		t.Fatalf("button attributes rendered in unexpected order\nwant:\n%s\n\ngot:\n%s", want, got)
	}
}
