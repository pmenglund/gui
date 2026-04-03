# gui

`github.com/pmenglund/gui` is a gomponents-native component library for server-rendered Go applications. It provides typed, HTMX-friendly UI components, a Tailwind-backed token system, and a small runtime for the interactive widgets that need keyboard and focus behavior.

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

## Validation

- `go test ./...`
- `npm run build:css`
- `npm run test:e2e`

