# gui

`github.com/pmenglund/gui` is a gomponents-native component library for server-rendered Go applications. It provides typed, HTMX-friendly UI components, a Tailwind-backed token system, and a small runtime for the interactive widgets that need keyboard and focus behavior.

## Using the module

Install the module with:

```sh
go get github.com/pmenglund/gui@latest
```

The module does not expose one large root API. Import the component packages you need directly from `components/...`, and import `github.com/pmenglund/gui/htmx` when you want to attach `hx-*` attributes through typed props.

```go
package ui

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/pmenglund/gui/components/button"
	"github.com/pmenglund/gui/components/card"
)

func SettingsCard() g.Node {
	return card.Card(card.Props{},
		h.H2(h.Class("text-lg font-semibold"), g.Text("Profile")),
		h.P(g.Text("Update the account details shown to other users.")),
		button.Button(button.Props{Variant: button.VariantDefault}, g.Text("Save changes")),
	)
}
```

To get the intended styling and behavior in your own app:

- Serve `theme/preset.css` on pages that render `gui` components.
- Build a Tailwind bundle from `assets/ui.css`. The simplest integration today is to copy that file into your app and compile it with a `content` list that includes both your app's Go files and the `gui` component source files you use.
- Serve `assets/ui.js` if you use interactive components such as `dialog`, `dropdownmenu`, `sheet`, `tabs`, or `toast`.
- Include HTMX itself only if your app uses the `htmx.Props` helpers.

The runnable showcase in `examples/showcase` is the reference integration. It shows the stylesheet and script wiring, component composition patterns, and HTMX partial-update flows end to end.

Use `Class` for styling overrides and use `Attributes` only for extra non-conflicting HTML attributes such as `title`, `lang`, `dir`, or custom `data-*` hooks. Raw passthrough attributes do not override component-managed semantics, typed `hx-*` props, or delegated runtime hooks in the `data-ui-*` namespace.

## Showcase

The repository includes a runnable showcase app so you can inspect the component library before wiring it into your own application.

<p>
  <img src="docs/screenshots/showcase-overview.png" alt="Overview page of the gui showcase app" width="49%">
  <img src="docs/screenshots/showcase-primitives.png" alt="Primitives page showing buttons, alerts, avatars, and badges" width="49%">
</p>
<p>
  <img src="docs/screenshots/showcase-forms.png" alt="Forms page showing form controls, navigation, and empty states" width="49%">
  <img src="docs/screenshots/showcase-interactive.png" alt="Interactive page showing dialogs, menus, tabs, toasts, and sheets" width="49%">
</p>

## Development

Run the CSS build and the showcase app from the repository root:

- `npm install`
- `npm run build:css`
- `go run ./examples/showcase`

The showcase serves:

- `/` for the overview
- `/primitives` for the primitive catalog
- `/forms` for form composition
- `/interactive` for dialog, dropdown, toast, and sheet
- `/htmx` for partial-update patterns
- `/examples/form-workflow` for a realistic, testable release request flow
- `/examples/runtime-workbench` for a testable runtime control surface

## Validation

- `go test ./...`
- `npm run build:css`
- `npm run test:e2e`
