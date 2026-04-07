package gui_test

import (
	"strings"
	"testing"

	"github.com/pmenglund/gui/examples/showcase/app"
	"github.com/pmenglund/gui/internal/testutil"
	g "maragu.dev/gomponents"
)

func TestShowcasePageGoldens(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name   string
		golden string
		node   g.Node
	}{
		{name: "primitives", golden: "showcase_primitives.html", node: app.Primitives()},
		{name: "forms", golden: "showcase_forms.html", node: app.Forms()},
		{name: "interactive", golden: "showcase_interactive.html", node: app.Interactive()},
		{name: "htmx", golden: "showcase_htmx.html", node: app.HTMXPage()},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := testutil.Render(t, tc.node)
			testutil.CompareGolden(t, tc.golden, got)
		})
	}
}

func TestHTMXFragmentsContainStableHooks(t *testing.T) {
	t.Parallel()

	activity := testutil.Render(t, app.ActivityFragment(2))
	if !strings.Contains(activity, `hx-get="/partials/activity?page=1"`) {
		t.Fatalf("activity fragment missing page 1 htmx link: %s", activity)
	}
	if !strings.Contains(activity, `hx-target="#activity-target"`) {
		t.Fatalf("activity fragment missing stable htmx target: %s", activity)
	}

	overlay := testutil.Render(t, app.OverlayFragment())
	if !strings.Contains(overlay, `data-ui-controller="dialog"`) {
		t.Fatalf("overlay fragment missing delegated controller hook: %s", overlay)
	}
	if !strings.Contains(overlay, `data-ui-trigger=""`) {
		t.Fatalf("overlay fragment missing delegated trigger hook: %s", overlay)
	}
	if !strings.Contains(overlay, `role="dialog" tabindex="-1"`) {
		t.Fatalf("overlay fragment missing focusable dialog surface: %s", overlay)
	}

	validation := testutil.Render(t, app.ValidationFragment("ops@example.com"))
	if !strings.Contains(validation, `Looks good for a company email.`) {
		t.Fatalf("validation fragment did not render success copy: %s", validation)
	}
}
