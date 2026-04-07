package app

import (
	"strings"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/pmenglund/gui/components/alert"
	"github.com/pmenglund/gui/components/badge"
	"github.com/pmenglund/gui/components/button"
	"github.com/pmenglund/gui/components/card"
	"github.com/pmenglund/gui/components/checkbox"
	"github.com/pmenglund/gui/components/dialog"
	"github.com/pmenglund/gui/components/dropdownmenu"
	"github.com/pmenglund/gui/components/fielddescription"
	"github.com/pmenglund/gui/components/fielderror"
	"github.com/pmenglund/gui/components/formfield"
	"github.com/pmenglund/gui/components/input"
	"github.com/pmenglund/gui/components/label"
	"github.com/pmenglund/gui/components/navbar"
	"github.com/pmenglund/gui/components/pagination"
	"github.com/pmenglund/gui/components/radiogroup"
	selectui "github.com/pmenglund/gui/components/select"
	"github.com/pmenglund/gui/components/sheet"
	switchui "github.com/pmenglund/gui/components/switch"
	"github.com/pmenglund/gui/components/table"
	"github.com/pmenglund/gui/components/tabs"
	"github.com/pmenglund/gui/components/textarea"
	"github.com/pmenglund/gui/components/toast"
)

type ExampleSpec struct {
	Slug    string
	Title   string
	Summary string
	Covered []string
	Build   func() g.Node
}

func TestableExamples() []ExampleSpec {
	return []ExampleSpec{
		{
			Slug:    "form-workflow",
			Title:   "Form workflow",
			Summary: "A release request form that combines explicit field wiring with typed form controls and review actions.",
			Covered: []string{
				"alert.Alert",
				"button.Button",
				"checkbox.Checkbox",
				"fielddescription.FieldDescription",
				"fielderror.FieldError",
				"formfield.FormField",
				"input.Input",
				"label.Label",
				"radiogroup.RadioGroup",
				"select.Select",
				"switch.Switch",
				"textarea.Textarea",
			},
			Build: FormWorkflowExample,
		},
		{
			Slug:    "runtime-workbench",
			Title:   "Runtime workbench",
			Summary: "An operator workbench with overlays, navigation, tabs, and feedback widgets that can be exercised in browser tests.",
			Covered: []string{
				"badge.Badge",
				"card.Card",
				"dialog.Dialog",
				"dropdownmenu.DropdownMenu",
				"navbar.Navbar",
				"pagination.Pagination",
				"sheet.Sheet",
				"table.Table",
				"tabs.Tabs",
				"toast.Toast",
			},
			Build: RuntimeWorkbenchExample,
		},
	}
}

func ExamplePath(slug string) string {
	return "/examples/" + slug
}

func FormWorkflowExample() g.Node {
	return shell("Form workflow", "",
		hero("Testable form workflow", "A realistic release request example that exercises explicit labels, helper text, validation copy, and typed selection controls."),
		cardGrid(
			sectionCard("Release guardrail", "The alert and action row frame the workflow before a user starts filling it in.",
				alert.Alert(alert.Props{
					Variant:     alert.VariantInfo,
					Title:       "Review before deploy",
					Description: "Production changes require an incident rollback plan and an approver from the owning team.",
				},
					h.Div(h.Class("flex flex-wrap gap-3 pt-2"),
						button.Button(button.Props{}, g.Text("Open checklist")),
						button.Button(button.Props{Variant: button.VariantOutline}, g.Text("View recent deploys")),
					),
				),
			),
			sectionCard("Deployment request", "FormField coordinates common field wiring, while the custom ticket block shows direct use of label, helper text, and error text.",
				formfield.FormField(formfield.Props{
					ID:          "form-workflow-service-name",
					Label:       "Service name",
					Description: "Shown in Slack notifications and the deployment feed.",
					Required:    true,
					Builder: func(ids formfield.IDs) g.Node {
						return input.Input(input.Props{
							ID:          ids.ControlID,
							Name:        "service",
							Value:       "search-api",
							DescribedBy: ids.DescriptionID,
						})
					},
				}),
				formfield.FormField(formfield.Props{
					ID:          "form-workflow-deployment-notes",
					Label:       "Deployment notes",
					Description: "Summarize the blast radius, rollback steps, and what changed.",
					Builder: func(ids formfield.IDs) g.Node {
						return textarea.Textarea(textarea.Props{
							ID:          ids.ControlID,
							Name:        "notes",
							Value:       "Ship the new query planner behind a 20% traffic split and monitor latency for ten minutes.",
							DescribedBy: ids.DescriptionID,
						})
					},
				}),
				releaseTicketField(),
			),
			sectionCard("Review settings", "These controls are rendered directly so tests can prove individual component behavior without relying only on composed wrappers.",
				checkbox.Checkbox(checkbox.Props{
					ID:      "form-workflow-announce",
					Name:    "announce",
					Label:   "Post an announcement to #deployments when this ships",
					Checked: true,
				}),
				switchui.Switch(switchui.Props{
					ID:      "form-workflow-canary",
					Name:    "canary",
					Label:   "Start in canary mode",
					Checked: true,
				}),
				radiogroup.RadioGroup(radiogroup.Props{
					ID:     "form-workflow-approval",
					Name:   "approval",
					Legend: "Approval path",
					Value:  "peer-review",
					Options: []radiogroup.Option{
						{Value: "peer-review", Label: "Peer review", Description: "Another engineer verifies the rollout notes and rollback plan."},
						{Value: "change-advisory", Label: "Change advisory", Description: "Required for high-risk changes outside the normal window."},
					},
				}),
				formfield.FormField(formfield.Props{
					ID:          "form-workflow-window",
					Label:       "Deployment window",
					Description: "Pick the window that matches the primary on-call rotation.",
					Builder: func(ids formfield.IDs) g.Node {
						return selectui.Select(selectui.Props{
							ID:          ids.ControlID,
							Name:        "window",
							Value:       "eu-morning",
							DescribedBy: ids.DescriptionID,
							Options: []selectui.Option{
								{Value: "us-early", Label: "US early morning"},
								{Value: "eu-morning", Label: "EU business morning"},
								{Value: "global-offhours", Label: "Global off-hours"},
							},
						})
					},
				}),
			),
		),
	)
}

func RuntimeWorkbenchExample() g.Node {
	toolbar := navbar.Navbar(navbar.Props{
		Brand: h.Div(
			h.Class("grid gap-1"),
			h.Span(h.Class("ui-kicker"), g.Text("Ops")),
			h.Strong(g.Text("Runtime workbench")),
		),
		Items: []navbar.Item{
			{Label: "Queue", Href: "#queue", Current: true},
			{Label: "Incidents", Href: "#incidents"},
			{Label: "Runbooks", Href: "#runbooks"},
		},
		Actions: []g.Node{
			badge.Badge(badge.Props{Variant: badge.VariantSuccess}, g.Text("3 healthy")),
		},
	})

	dialogExample := dialog.Dialog(dialog.Props{
		ID:          "runtime-workbench-dialog",
		Title:       "Approve rollout",
		Description: "The dialog stays testable because all behavior is attached through data attributes on server-rendered markup.",
		Trigger:     button.Button(button.Props{}, g.Text("Open approval dialog")),
		Footer:      button.Button(button.Props{Variant: button.VariantOutline}, g.Text("Approve rollout")),
	}, h.P(g.Text("Confirm the canary has remained healthy for five minutes before promoting the release.")))

	sheetExample := sheet.Sheet(sheet.Props{
		ID:          "runtime-workbench-sheet",
		Title:       "Operator notes",
		Description: "The sheet is useful for secondary details that should not crowd the main grid.",
		Trigger:     button.Button(button.Props{Variant: button.VariantOutline}, g.Text("Open operator sheet")),
	},
		card.Card(card.Props{Title: "Escalation path", Description: "Use the paging rotation if the canary error budget drops below target."},
			h.Ul(h.Class("list-disc pl-5 text-sm"), h.Li(g.Text("Page platform on-call")), h.Li(g.Text("Pause the rollout")), h.Li(g.Text("Post an update in Slack"))),
		),
	)

	dropdownExample := dropdownmenu.DropdownMenu(dropdownmenu.Props{
		Trigger: button.Button(button.Props{Variant: button.VariantOutline}, g.Text("Quick actions")),
		Items: []dropdownmenu.Item{
			{Label: "View service logs", Href: "#logs"},
			{Label: "Pause rollout", Dangerous: true},
		},
	})

	toastExample := toast.Toast(toast.Props{
		Title:       "Incident note saved",
		Description: "The toast is a lightweight acknowledgement for non-blocking actions.",
		Trigger:     button.Button(button.Props{Variant: button.VariantOutline}, g.Text("Show save toast")),
	})

	tabExample := tabs.Tabs(tabs.Props{
		ID: "runtime-workbench-tabs",
		Items: []tabs.Item{
			{Key: "overview", Label: "Overview", Panel: h.P(g.Text("The overview panel highlights the steady-state health of the rollout."))},
			{Key: "activity", Label: "Activity", Panel: h.P(g.Text("Keyboard arrow keys move focus between tabs without a client-side framework."))},
			{Key: "handoff", Label: "Handoff", Panel: h.P(g.Text("The handoff panel keeps operator notes server-rendered and easy to snapshot test."))},
		},
	})

	pipelineTable := card.Card(card.Props{
		Title:       "Release queue",
		Description: "The table and pagination make the workbench easy to exercise in both golden and browser tests.",
	},
		table.Table(table.Props{
			Columns: []table.Column{{Header: "Service"}, {Header: "Stage"}, {Header: "Status"}},
			Rows: []table.Row{
				{Cells: []g.Node{g.Text("search-api"), g.Text("Canary"), badge.Badge(badge.Props{Variant: badge.VariantSuccess}, g.Text("Healthy"))}},
				{Cells: []g.Node{g.Text("billing"), g.Text("Queued"), badge.Badge(badge.Props{Variant: badge.VariantMuted}, g.Text("Waiting"))}},
			},
		}),
		pagination.Pagination(pagination.Props{
			Items: []pagination.Item{
				{Label: "1", Href: "#", Current: true},
				{Label: "2", Href: "#"},
				{Label: "3", Href: "#"},
			},
		}),
	)

	return shell("Runtime workbench", "",
		hero("Testable runtime workbench", "A focused example page for overlays, menus, tabs, and operator feedback controls."),
		toolbar,
		cardGrid(
			sectionCard("Operational queue", "The workbench keeps the primary data view and the action affordances close together.", pipelineTable),
			sectionCard("Interactive controls", "These widgets are the browser-test targets for this issue.", dialogExample, sheetExample, dropdownExample, toastExample, tabExample),
		),
	)
}

func releaseTicketField() g.Node {
	const fieldID = "release-ticket"
	const helpID = fieldID + "-help"
	const errorID = fieldID + "-error"

	return h.Div(
		h.Class("grid gap-2"),
		label.Label(label.Props{For: fieldID, Required: true}, g.Text("Change ticket")),
		input.Input(input.Props{
			ID:          fieldID,
			Name:        "ticket",
			Value:       "CHG-1842",
			Invalid:     true,
			DescribedBy: strings.Join([]string{helpID, errorID}, " "),
		}),
		fielddescription.FieldDescription(fielddescription.Props{ID: helpID}, g.Text("Use the CAB ticket for high-risk changes and link it in the rollout note.")),
		fielderror.FieldError(fielderror.Props{ID: errorID}, g.Text("Ticket ownership has not been assigned yet.")),
	)
}
