package gui_test

import (
	"regexp"
	"strings"
	"testing"

	"github.com/pmenglund/gui/components/checkbox"
	"github.com/pmenglund/gui/components/dialog"
	"github.com/pmenglund/gui/components/formfield"
	"github.com/pmenglund/gui/components/input"
	"github.com/pmenglund/gui/components/radiogroup"
	"github.com/pmenglund/gui/components/sheet"
	switchui "github.com/pmenglund/gui/components/switch"
	"github.com/pmenglund/gui/components/tabs"
	"github.com/pmenglund/gui/internal/testutil"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

var (
	idPattern              = regexp.MustCompile(`\sid="([^"]+)"`)
	forPattern             = regexp.MustCompile(`\sfor="([^"]+)"`)
	labelledByPattern      = regexp.MustCompile(`\saria-labelledby="([^"]+)"`)
	describedByPattern     = regexp.MustCompile(`\saria-describedby="([^"]+)"`)
	controlsPattern        = regexp.MustCompile(`\saria-controls="([^"]+)"`)
	relationshipIDPatterns = []*regexp.Regexp{
		forPattern,
		labelledByPattern,
		describedByPattern,
		controlsPattern,
	}
)

func TestRepeatedRelationshipComponentsGenerateDistinctIDs(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name  string
		build func() g.Node
	}{
		{
			name: "dialog",
			build: func() g.Node {
				return dialog.Dialog(dialog.Props{
					Title:       "Review rollout",
					Description: "Confirm the canary has stayed healthy.",
				}, h.P(g.Text("Ship it.")))
			},
		},
		{
			name: "sheet",
			build: func() g.Node {
				return sheet.Sheet(sheet.Props{
					Title:       "Operator notes",
					Description: "Secondary details live here.",
				}, h.P(g.Text("Escalation path.")))
			},
		},
		{
			name: "tabs",
			build: func() g.Node {
				return tabs.Tabs(tabs.Props{
					Items: []tabs.Item{
						{Key: "overview", Label: "Overview", Panel: h.P(g.Text("Overview"))},
						{Key: "activity", Label: "Activity", Panel: h.P(g.Text("Activity"))},
					},
				})
			},
		},
		{
			name: "formfield",
			build: func() g.Node {
				return formfield.FormField(formfield.Props{
					Label:       "Work email",
					Description: "Used for notifications.",
					Error:       "Use a company email address.",
					Builder: func(ids formfield.IDs) g.Node {
						return input.Input(input.Props{
							ID:          ids.ControlID,
							Type:        "email",
							DescribedBy: strings.TrimSpace(strings.Join([]string{ids.DescriptionID, ids.ErrorID}, " ")),
						})
					},
				})
			},
		},
		{
			name: "checkbox",
			build: func() g.Node {
				return checkbox.Checkbox(checkbox.Props{
					Label: "Enable shipping updates",
				})
			},
		},
		{
			name: "switch",
			build: func() g.Node {
				return switchui.Switch(switchui.Props{
					Label: "Compact mode",
				})
			},
		},
		{
			name: "radiogroup",
			build: func() g.Node {
				return radiogroup.RadioGroup(radiogroup.Props{
					Name:   "approval",
					Legend: "Approval path",
					Value:  "peer-review",
					Options: []radiogroup.Option{
						{Value: "peer-review", Label: "Peer review"},
						{Value: "change-advisory", Label: "Change advisory"},
					},
				})
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			first := testutil.Render(t, tc.build())
			second := testutil.Render(t, tc.build())
			assertRelationshipRefsResolve(t, first)
			assertRelationshipRefsResolve(t, second)

			combined := testutil.Render(t, h.Div(
				h.Div(h.Data("instance", "first"), tc.build()),
				h.Div(h.Data("instance", "second"), tc.build()),
			))
			assertUniqueIDs(t, combined)
			assertRelationshipRefsResolve(t, combined)
		})
	}
}

func assertUniqueIDs(t *testing.T, html string) {
	t.Helper()

	seen := map[string]struct{}{}
	for _, match := range idPattern.FindAllStringSubmatch(html, -1) {
		id := match[1]
		if _, ok := seen[id]; ok {
			t.Fatalf("duplicate id %q in rendered HTML:\n%s", id, html)
		}
		seen[id] = struct{}{}
	}
}

func assertRelationshipRefsResolve(t *testing.T, html string) {
	t.Helper()

	ids := map[string]struct{}{}
	for _, match := range idPattern.FindAllStringSubmatch(html, -1) {
		ids[match[1]] = struct{}{}
	}

	for _, pattern := range relationshipIDPatterns {
		for _, match := range pattern.FindAllStringSubmatch(html, -1) {
			for _, ref := range strings.Fields(match[1]) {
				if _, ok := ids[ref]; !ok {
					t.Fatalf("reference %q from %q does not resolve in rendered HTML:\n%s", ref, match[0], html)
				}
			}
		}
	}
}
