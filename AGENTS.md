# Repository Guidelines

## Project Structure & Module Organization

`components/` holds the public leaf packages for the UI library. Each component lives in its own directory, for example `components/button` or `components/dialog`. `htmx/` exposes stable HTMX props used by public component APIs. `internal/` contains shared implementation helpers for accessibility, attribute rendering, class composition, and test utilities; do not import these from outside the module. `assets/` contains the client runtime and Tailwind entry CSS, `theme/` defines token CSS, and `examples/showcase/` is the runnable demo app. Golden snapshots live in `testdata/golden/`, and browser tests live in `tests/e2e/`.

## Build, Test, and Development Commands

- `npm run build:css` builds `examples/showcase/static/ui.css` from `assets/ui.css`.
- `go test ./...` runs the Go unit and golden tests.
- `npm run test:e2e` runs the Playwright suite against the showcase app.
- `go run ./examples/showcase` starts the local demo server on `http://localhost:8080`.

Run commands from the repository root.

## Coding Style & Naming Conventions

Use standard Go formatting with `gofmt -w`. Keep public component APIs typed and zero-value-safe. Follow the existing constructor shape: `func ComponentName(p Props, children ...g.Node) g.Node`. Use clear leaf package names matching the directory where possible; note that `components/select` and `components/switch` use package names `selectui` and `switchui` because of Go keywords. Keep shared implementation details under `internal/`.

For CSS and runtime code, prefer small, explicit changes. Tailwind classes should stay centralized in component code and `assets/ui.css`, and runtime hooks should use the documented `data-ui-*` attributes.

## Testing Guidelines

Add or update Go tests for markup changes and keep the golden files in `testdata/golden/` in sync. Browser behavior belongs in `tests/e2e/showcase.spec.ts`. Name Go tests with the usual `TestXxx` pattern and keep Playwright scenarios focused on observable behavior such as focus handling, HTMX swaps, and dismissal flows.

## Commit & Pull Request Guidelines

Use short, imperative commit messages in the style of the existing history, for example `Implement HTMX-friendly gomponents UI library`. PRs should summarize the user-visible change, list validation commands run, and include screenshots or short notes when updating showcase behavior.

## Agent-Specific Notes

Check the worktree before editing; this repository may change between turns. Do not commit generated artifacts outside the tracked golden files and source assets. `node_modules/`, Playwright output, and generated CSS are build products, not source.
