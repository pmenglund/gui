package formfield

import (
	"strings"
	"testing"

	inputui "github.com/pmenglund/gui/components/input"
	selectui "github.com/pmenglund/gui/components/select"
	textareaui "github.com/pmenglund/gui/components/textarea"
	"github.com/pmenglund/gui/internal/testutil"
)

func TestInputAppliesFieldAccessibilityWiring(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		field    Props
		control  inputui.Props
		contains []string
		excludes []string
	}{
		{
			name: "description only",
			field: Props{
				Label:       "Work email",
				Description: "Used for deployment notifications.",
			},
			control: inputui.Props{Type: "email"},
			contains: []string{
				`for="field-work-email"`,
				`id="field-work-email"`,
				`aria-describedby="field-work-email-description"`,
				`id="field-work-email-description"`,
			},
			excludes: []string{`aria-invalid="true"`},
		},
		{
			name: "error only",
			field: Props{
				Label: "Work email",
				Error: "Please use a company email address.",
			},
			control: inputui.Props{Type: "email"},
			contains: []string{
				`for="field-work-email"`,
				`id="field-work-email"`,
				`aria-describedby="field-work-email-error"`,
				`aria-invalid="true"`,
				`id="field-work-email-error"`,
			},
		},
		{
			name: "description and error with extra described by",
			field: Props{
				Label:       "Work email",
				Description: "Used for deployment notifications.",
				Error:       "Please use a company email address.",
			},
			control: inputui.Props{
				Type:        "email",
				DescribedBy: "email-hint",
			},
			contains: []string{
				`for="field-work-email"`,
				`id="field-work-email"`,
				`aria-describedby="field-work-email-description field-work-email-error email-hint"`,
				`aria-invalid="true"`,
				`id="field-work-email-description"`,
				`id="field-work-email-error"`,
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := testutil.Render(t, Input(tc.field, tc.control))
			for _, want := range tc.contains {
				if !strings.Contains(got, want) {
					t.Fatalf("expected rendered field to contain %q\n%s", want, got)
				}
			}
			for _, unwanted := range tc.excludes {
				if strings.Contains(got, unwanted) {
					t.Fatalf("expected rendered field not to contain %q\n%s", unwanted, got)
				}
			}
		})
	}
}

func TestTextareaAppliesFieldAccessibilityWiring(t *testing.T) {
	t.Parallel()

	got := testutil.Render(t, Textarea(
		Props{
			Label:       "Deployment notes",
			Description: "Rendered as a textarea with preserved native behavior.",
		},
		textareaui.Props{Value: "Roll out to the canary ring first."},
	))

	for _, want := range []string{
		`for="field-deployment-notes"`,
		`id="field-deployment-notes"`,
		`aria-describedby="field-deployment-notes-description"`,
		`id="field-deployment-notes-description"`,
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("expected rendered textarea field to contain %q\n%s", want, got)
		}
	}
}

func TestSelectAppliesFieldAccessibilityWiring(t *testing.T) {
	t.Parallel()

	got := testutil.Render(t, Select(
		Props{
			Label:       "Deployment window",
			Description: "Pick the window that matches the primary on-call rotation.",
		},
		selectui.Props{
			Value: "eu-morning",
			Options: []selectui.Option{
				{Value: "us-early", Label: "US early morning"},
				{Value: "eu-morning", Label: "EU business morning"},
			},
		},
	))

	for _, want := range []string{
		`for="field-deployment-window"`,
		`id="field-deployment-window"`,
		`aria-describedby="field-deployment-window-description"`,
		`id="field-deployment-window-description"`,
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("expected rendered select field to contain %q\n%s", want, got)
		}
	}
}
