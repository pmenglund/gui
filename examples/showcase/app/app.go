package app

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/pmenglund/gui/components/alert"
	"github.com/pmenglund/gui/components/avatar"
	"github.com/pmenglund/gui/components/badge"
	"github.com/pmenglund/gui/components/breadcrumbs"
	"github.com/pmenglund/gui/components/button"
	"github.com/pmenglund/gui/components/card"
	"github.com/pmenglund/gui/components/checkbox"
	"github.com/pmenglund/gui/components/dialog"
	"github.com/pmenglund/gui/components/dropdownmenu"
	"github.com/pmenglund/gui/components/emptystate"
	"github.com/pmenglund/gui/components/formfield"
	"github.com/pmenglund/gui/components/iconbutton"
	"github.com/pmenglund/gui/components/input"
	"github.com/pmenglund/gui/components/navbar"
	"github.com/pmenglund/gui/components/pagination"
	"github.com/pmenglund/gui/components/radiogroup"
	selectui "github.com/pmenglund/gui/components/select"
	"github.com/pmenglund/gui/components/separator"
	"github.com/pmenglund/gui/components/sheet"
	"github.com/pmenglund/gui/components/skeleton"
	"github.com/pmenglund/gui/components/spinner"
	switchui "github.com/pmenglund/gui/components/switch"
	"github.com/pmenglund/gui/components/table"
	"github.com/pmenglund/gui/components/tabs"
	"github.com/pmenglund/gui/components/textarea"
	"github.com/pmenglund/gui/components/toast"
	public "github.com/pmenglund/gui/htmx"
)

func NewMux(root string) http.Handler {
	mux := http.NewServeMux()

	staticDir := filepath.Join(root, "examples", "showcase", "static")
	assetsDir := filepath.Join(root, "assets")
	themeDir := filepath.Join(root, "theme")
	vendorDir := filepath.Join(root, "node_modules", "htmx.org", "dist")

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(assetsDir))))
	mux.Handle("/theme/", http.StripPrefix("/theme/", http.FileServer(http.Dir(themeDir))))
	mux.Handle("/vendor/", http.StripPrefix("/vendor/", http.FileServer(http.Dir(vendorDir))))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		renderPage(w, Overview())
	})
	mux.HandleFunc("/primitives", func(w http.ResponseWriter, r *http.Request) { renderPage(w, Primitives()) })
	mux.HandleFunc("/forms", func(w http.ResponseWriter, r *http.Request) { renderPage(w, Forms()) })
	mux.HandleFunc("/interactive", func(w http.ResponseWriter, r *http.Request) { renderPage(w, Interactive()) })
	mux.HandleFunc("/htmx", func(w http.ResponseWriter, r *http.Request) { renderPage(w, HTMXPage()) })
	for _, example := range TestableExamples() {
		example := example
		mux.HandleFunc(ExamplePath(example.Slug), func(w http.ResponseWriter, r *http.Request) {
			renderPage(w, example.Build())
		})
	}
	mux.HandleFunc("/partials/counter", func(w http.ResponseWriter, r *http.Request) {
		value, _ := strconv.Atoi(r.URL.Query().Get("value"))
		if value == 0 {
			value = 1
		}
		renderNode(w, CounterFragment(value))
	})
	mux.HandleFunc("/partials/validate", func(w http.ResponseWriter, r *http.Request) {
		renderNode(w, ValidationFragment(r.URL.Query().Get("email")))
	})
	mux.HandleFunc("/partials/activity", func(w http.ResponseWriter, r *http.Request) {
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		if page == 0 {
			page = 1
		}
		renderNode(w, ActivityFragment(page))
	})
	mux.HandleFunc("/partials/overlay-trigger", func(w http.ResponseWriter, r *http.Request) {
		renderNode(w, OverlayFragment())
	})

	return mux
}

func Overview() g.Node {
	exampleCards := make([]g.Node, 0, len(TestableExamples()))
	for _, example := range TestableExamples() {
		exampleCards = append(exampleCards,
			sectionCard(example.Title, example.Summary,
				h.A(h.Href(ExamplePath(example.Slug)), h.Class("font-medium hover:underline"), g.Text("Open example")),
			),
		)
	}

	return shell("Overview", "/",
		hero("gomponents UI for HTMX-heavy Go apps", "Typed components, semantic defaults, and progressive enhancement in one place."),
		cardGrid(
			sectionCard("Primitive catalog", "Buttons, form controls, badges, alerts, and loading states.", h.A(h.Href("/primitives"), h.Class("font-medium hover:underline"), g.Text("Browse primitives"))),
			sectionCard("Forms and navigation", "Accessible label wiring, table patterns, tabs, breadcrumbs, and pagination.", h.A(h.Href("/forms"), h.Class("font-medium hover:underline"), g.Text("Inspect forms"))),
			sectionCard("Interactive runtime", "Dialog, sheet, dropdown, toast, and tabs powered by a tiny delegated runtime.", h.A(h.Href("/interactive"), h.Class("font-medium hover:underline"), g.Text("Try interactions"))),
			sectionCard("HTMX integration", "Built-in `hx-*` support, partial updates, and swapped fragments that stay interactive.", h.A(h.Href("/htmx"), h.Class("font-medium hover:underline"), g.Text("Open HTMX flows"))),
			g.Group(exampleCards),
		),
	)
}

func Primitives() g.Node {
	return shell("Primitives", "/primitives",
		hero("Primitive catalog", "The stable building blocks use native HTML semantics, typed props, and optional HTMX attributes."),
		cardGrid(
			sectionCard("Buttons and badges", "", h.Div(h.Class("flex flex-wrap gap-3"),
				button.Button(button.Props{}, g.Text("Default")),
				button.Button(button.Props{Variant: button.VariantOutline}, g.Text("Outline")),
				button.Button(button.Props{Variant: button.VariantDestructive}, g.Text("Delete")),
				iconbutton.IconButton(iconbutton.Props{Label: "Search", Variant: button.VariantGhost}, h.Span(g.Text("⌕"))),
				badge.Badge(badge.Props{}, g.Text("Stable")),
				badge.Badge(badge.Props{Variant: badge.VariantSuccess}, g.Text("Healthy")),
			)),
			sectionCard("Inputs", "", formfield.FormField(formfield.Props{
				Label:       "Project name",
				Description: "A plain text input with token-driven focus styles.",
				Builder: func(ids formfield.IDs) g.Node {
					return input.Input(input.Props{ID: ids.ControlID, Placeholder: "Aurora", DescribedBy: ids.DescriptionID})
				},
			}), formfield.FormField(formfield.Props{
				Label: "Project summary",
				Builder: func(ids formfield.IDs) g.Node {
					return textarea.Textarea(textarea.Props{ID: ids.ControlID, Value: "Server-rendered interfaces with a narrow client runtime."})
				},
			})),
			sectionCard("Selection controls", "", checkbox.Checkbox(checkbox.Props{Label: "Enable shipping updates", Checked: true}),
				radiogroup.RadioGroup(radiogroup.Props{
					Name:   "tier",
					Legend: "Support tier",
					Value:  "team",
					Options: []radiogroup.Option{
						{Value: "solo", Label: "Solo"},
						{Value: "team", Label: "Team", Description: "Shared access and audit trail."},
					},
				}),
				switchui.Switch(switchui.Props{Label: "Compact mode", Checked: true}),
				selectui.Select(selectui.Props{
					Placeholder: "Choose a region",
					Options: []selectui.Option{
						{Value: "us-east-1", Label: "US East 1"},
						{Value: "eu-west-1", Label: "EU West 1"},
					},
				}),
			),
			sectionCard("Feedback and loading", "", alert.Alert(alert.Props{Title: "Heads up", Description: "This alert keeps the copy short and the semantics obvious."}),
				h.Div(h.Class("flex items-center gap-4"), spinner.Spinner(spinner.Props{Label: "Loading metrics"}), skeleton.Skeleton(skeleton.Props{Height: "h-10", Count: 2}), separator.Separator(separator.Props{Vertical: true}), avatar.Avatar(avatar.Props{Name: "Priya Mehta"}))),
		),
	)
}

func Forms() g.Node {
	return shell("Forms", "/forms",
		hero("Forms and navigation", "FormField coordinates IDs, descriptions, and errors while the navigation components stay plain HTML."),
		cardGrid(
			sectionCard("Accessible field composition", "",
				formfield.FormField(formfield.Props{
					Label:       "Work email",
					Description: "Used for deployment notifications.",
					Error:       "Please use a company email address.",
					Required:    true,
					Builder: func(ids formfield.IDs) g.Node {
						return input.Input(input.Props{
							ID:          ids.ControlID,
							Type:        "email",
							Value:       "person@example.com",
							Invalid:     true,
							DescribedBy: strings.TrimSpace(strings.Join([]string{ids.DescriptionID, ids.ErrorID}, " ")),
						})
					},
				}),
				formfield.FormField(formfield.Props{
					Label:       "Deployment notes",
					Description: "Rendered as a textarea with preserved native behavior.",
					Builder: func(ids formfield.IDs) g.Node {
						return textarea.Textarea(textarea.Props{
							ID:          ids.ControlID,
							Value:       "Roll out to the canary ring first.",
							DescribedBy: ids.DescriptionID,
						})
					},
				}),
			),
			sectionCard("Navigation patterns", "",
				breadcrumbs.Breadcrumbs(breadcrumbs.Props{Items: []breadcrumbs.Item{
					{Label: "Home", Href: "/"},
					{Label: "Deployments", Href: "/forms"},
					{Label: "Current", Current: true},
				}}),
				table.Table(table.Props{
					Columns: []table.Column{{Header: "Service"}, {Header: "Status"}, {Header: "Owner"}},
					Rows: []table.Row{
						{Cells: []g.Node{g.Text("api"), badge.Badge(badge.Props{Variant: badge.VariantSuccess}, g.Text("Healthy")), g.Text("Platform")}},
						{Cells: []g.Node{g.Text("worker"), badge.Badge(badge.Props{Variant: badge.VariantMuted}, g.Text("Queued")), g.Text("Operations")}},
					},
				}),
				pagination.Pagination(pagination.Props{Items: []pagination.Item{
					{Label: "1", Href: "#", Current: true},
					{Label: "2", Href: "#"},
					{Label: "3", Href: "#"},
				}}),
			),
			sectionCard("Empty state", "",
				emptystate.EmptyState(emptystate.Props{
					Eyebrow:     "No incidents",
					Title:       "Everything is quiet",
					Description: "Use this when a panel needs a strong empty state without inventing a one-off wrapper.",
					Action:      button.Button(button.Props{Variant: button.VariantOutline}, g.Text("Create runbook")),
				}),
			),
		),
	)
}

func Interactive() g.Node {
	dialogExample := dialog.Dialog(dialog.Props{
		Title:       "Review rollout",
		Description: "This dialog is opened by the delegated runtime and closes on Escape.",
		Trigger:     button.Button(button.Props{}, g.Text("Open dialog")),
		Footer:      button.Button(button.Props{Variant: button.VariantOutline}, g.Text("Confirm")),
	}, h.P(g.Text("The dialog content stays in place in the DOM and relies on focus restoration instead of portals.")))

	sheetExample := sheet.Sheet(sheet.Props{
		Title:       "Deployment details",
		Description: "The sheet shares the same runtime state model as the dialog.",
		Trigger:     button.Button(button.Props{Variant: button.VariantOutline}, g.Text("Open sheet")),
	}, h.P(g.Text("This panel uses the right edge of the viewport and the same close affordances.")))

	dropdownExample := dropdownmenu.DropdownMenu(dropdownmenu.Props{
		Trigger: button.Button(button.Props{Variant: button.VariantOutline}, g.Text("Quick actions")),
		Items: []dropdownmenu.Item{
			{Label: "View logs", Href: "/interactive"},
			{Label: "Restart canary", Dangerous: true},
		},
	})

	toastExample := toast.Toast(toast.Props{
		Title:       "Settings saved",
		Description: "The toast uses the same delegated runtime hooks as the other widgets.",
		Trigger:     button.Button(button.Props{Variant: button.VariantOutline}, g.Text("Show toast")),
	})

	tabExample := tabs.Tabs(tabs.Props{
		Items: []tabs.Item{
			{Key: "overview", Label: "Overview", Panel: h.P(g.Text("The overview panel is visible on first paint."))},
			{Key: "activity", Label: "Activity", Panel: h.P(g.Text("The tabs runtime handles keyboard navigation with arrow keys."))},
			{Key: "logs", Label: "Logs", Panel: h.P(g.Text("Panels remain server-rendered HTML; the runtime only manages state."))},
		},
	})

	return shell("Interactive", "/interactive",
		hero("Interactive runtime", "The widgets keep their markup server-rendered and attach behavior with stable `data-ui-*` hooks."),
		cardGrid(
			sectionCard("Overlays", "", dialogExample, sheetExample),
			sectionCard("Menus and tabs", "", dropdownExample, tabExample),
			sectionCard("Toast", "", toastExample),
		),
	)
}

func HTMXPage() g.Node {
	return shell("HTMX", "/htmx",
		hero("HTMX support", "Components accept typed HTMX props and continue working after swaps because the runtime is delegated."),
		cardGrid(
			sectionCard("Counter refresh", "",
				button.Button(button.Props{
					HTMX: public.Props{
						Get:    "/partials/counter?value=2",
						Target: "#counter-target",
						Swap:   "innerHTML",
					},
				}, g.Text("Refresh counter fragment")),
				h.Div(h.ID("counter-target"), CounterFragment(1)),
			),
			sectionCard("Inline validation", "",
				formfield.FormField(formfield.Props{
					Label:       "Email",
					Description: "The hint below is replaced by an HTMX request as you type.",
					Builder: func(ids formfield.IDs) g.Node {
						return input.Input(input.Props{
							ID:          ids.ControlID,
							Name:        "email",
							Placeholder: "you@company.com",
							DescribedBy: ids.DescriptionID + " email-hint",
							HTMX: public.Props{
								Get:     "/partials/validate",
								Trigger: "keyup changed delay:300ms",
								Target:  "#email-hint",
								Swap:    "outerHTML",
							},
						})
					},
				}),
				ValidationFragment(""),
			),
			sectionCard("Paged table", "",
				h.Div(h.ID("activity-target"), ActivityFragment(1)),
			),
			sectionCard("Swapped interactive fragment", "",
				button.Button(button.Props{
					HTMX: public.Props{
						Get:    "/partials/overlay-trigger",
						Target: "#overlay-target",
						Swap:   "innerHTML",
					},
				}, g.Text("Load interactive fragment")),
				h.Div(h.ID("overlay-target")),
			),
		),
	)
}

func CounterFragment(value int) g.Node {
	return card.Card(card.Props{
		Title:       fmt.Sprintf("Counter update %d", value),
		Description: "This fragment is server-rendered and swapped into the page with HTMX.",
	}, badge.Badge(badge.Props{}, g.Text("Fresh partial")))
}

func ValidationFragment(email string) g.Node {
	if strings.Contains(email, "@") && strings.Contains(email, ".") {
		return h.P(h.ID("email-hint"), h.Class("text-sm font-medium text-[rgb(var(--ui-success))]"), g.Text("Looks good for a company email."))
	}
	return h.P(h.ID("email-hint"), h.Class("text-sm font-medium text-[rgb(var(--ui-danger))]"), g.Text("Enter a valid email address to continue."))
}

func ActivityFragment(page int) g.Node {
	rows := map[int][]table.Row{
		1: {
			{Cells: []g.Node{g.Text("API"), g.Text("Canary"), badge.Badge(badge.Props{Variant: badge.VariantSuccess}, g.Text("Healthy"))}},
			{Cells: []g.Node{g.Text("Worker"), g.Text("Primary"), badge.Badge(badge.Props{Variant: badge.VariantMuted}, g.Text("Queued"))}},
		},
		2: {
			{Cells: []g.Node{g.Text("Billing"), g.Text("Canary"), badge.Badge(badge.Props{Variant: badge.VariantSuccess}, g.Text("Healthy"))}},
			{Cells: []g.Node{g.Text("Search"), g.Text("Primary"), badge.Badge(badge.Props{Variant: badge.VariantDefault}, g.Text("Observing"))}},
		},
	}
	if _, ok := rows[page]; !ok {
		page = 1
	}

	return h.Div(
		h.Class("grid gap-4"),
		table.Table(table.Props{
			Columns: []table.Column{{Header: "Service"}, {Header: "Track"}, {Header: "Status"}},
			Rows:    rows[page],
		}),
		pagination.Pagination(pagination.Props{
			Items: []pagination.Item{
				{Label: "1", Href: "#", Current: page == 1, HTMX: public.Props{Get: "/partials/activity?page=1", Target: "#activity-target", Swap: "innerHTML"}},
				{Label: "2", Href: "#", Current: page == 2, HTMX: public.Props{Get: "/partials/activity?page=2", Target: "#activity-target", Swap: "innerHTML"}},
			},
		}),
	)
}

func OverlayFragment() g.Node {
	return dialog.Dialog(dialog.Props{
		Title:       "Swapped fragment",
		Description: "This dialog trigger arrived through an HTMX swap and still works.",
		Trigger:     button.Button(button.Props{Variant: button.VariantOutline}, g.Text("Open swapped dialog")),
	}, h.P(g.Text("Delegated runtime hooks mean the client code does not need a manual rebind.")))
}

func renderPage(w http.ResponseWriter, node g.Node) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	renderNode(w, node)
}

func renderNode(w http.ResponseWriter, node g.Node) {
	if err := node.Render(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func shell(title, current string, sections ...g.Node) g.Node {
	return h.Doctype(h.HTML(
		h.Lang("en"),
		h.Head(
			h.Meta(h.Charset("utf-8")),
			h.Meta(h.Name("viewport"), h.Content("width=device-width, initial-scale=1")),
			h.TitleEl(g.Text(title+" • gui showcase")),
			h.Link(h.Rel("stylesheet"), h.Href("/theme/preset.css")),
			h.Link(h.Rel("stylesheet"), h.Href("/static/ui.css")),
			h.Script(h.Src("/vendor/htmx.min.js")),
			h.Script(h.Defer(), h.Src("/assets/ui.js")),
		),
		h.Body(
			h.Div(
				h.Class("ui-shell ui-page-grid"),
				navbar.Navbar(navbar.Props{
					Brand: h.Div(
						h.Class("grid gap-1"),
						h.Span(h.Class("ui-kicker"), g.Text("gui")),
						h.Strong(g.Text("Component showcase")),
					),
					Items: []navbar.Item{
						{Label: "Overview", Href: "/", Current: current == "/"},
						{Label: "Primitives", Href: "/primitives", Current: current == "/primitives"},
						{Label: "Forms", Href: "/forms", Current: current == "/forms"},
						{Label: "Interactive", Href: "/interactive", Current: current == "/interactive"},
						{Label: "HTMX", Href: "/htmx", Current: current == "/htmx"},
					},
				}),
				g.Group(sections),
			),
		),
	))
}

func hero(title, copy string) g.Node {
	return card.Card(card.Props{}, h.P(h.Class("ui-kicker"), g.Text("Showcase")), h.H1(h.Class("ui-section-title"), g.Text(title)), h.P(h.Class("max-w-3xl text-sm text-[rgb(var(--ui-muted-foreground))]"), g.Text(copy)))
}

func cardGrid(cards ...g.Node) g.Node {
	return h.Section(h.Class("ui-card-grid"), g.Group(cards))
}

func sectionCard(title, description string, children ...g.Node) g.Node {
	content := []g.Node{
		h.H2(h.Class("text-lg font-semibold"), g.Text(title)),
	}
	if description != "" {
		content = append(content, h.P(h.Class("text-sm text-[rgb(var(--ui-muted-foreground))]"), g.Text(description)))
	}
	content = append(content, children...)
	return h.Section(h.Class("ui-demo-card grid gap-4"), g.Group(content))
}
