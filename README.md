# gui

`github.com/pmenglund/gui` is a gomponents-native component library for server-rendered Go applications. It provides typed, HTMX-friendly UI components, a Tailwind-backed token system, and a small runtime for the interactive widgets that need keyboard and focus behavior.

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
