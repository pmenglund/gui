# Build the full v1 `github.com/pmenglund/goth` gomponents UI library with first-class HTMX support

This ExecPlan is a living document. Keep `Progress`, `Surprises & Discoveries`, `Decision Log`, and `Outcomes & Retrospective` current as implementation proceeds.

## Purpose / Big Picture

This repository currently contains only the specification and module file. The goal of this work is to turn it into a usable gomponents-native UI library for Go web applications. After this change, a developer should be able to import leaf component packages, style them with Tailwind-backed theme tokens, wire them to HTMX without ad hoc attribute assembly, and verify the result in a local showcase plus automated tests.

## Progress

- [x] Bootstrap repository directories, Go dependency, and npm toolchain.
- [x] Add the shared theme, runtime, and helper packages.
- [x] Implement the stable primitive components.
- [x] Implement the stable composite and interactive components.
- [x] Add the showcase app, golden tests, and Playwright coverage.

## Surprises & Discoveries

- Observation: the latest gomponents module uses the import path `maragu.dev/gomponents`, not the older GitHub path.
  Evidence: `go get github.com/maragudk/gomponents@v1.2.0` failed because the module declares `maragu.dev/gomponents`.
- Observation: Tailwind 4 requires Node 20 in this environment, but the workspace has Node 18.20.8.
  Evidence: installing `@tailwindcss/cli@4.2.2` emitted an engine warning for `node >= 20`.

## Decision Log

- Decision: use leaf component packages rather than a single umbrella component package.
  Rationale: it keeps imports explicit and matches the specification.
  Date/Author: 2026-04-02, user.
- Decision: build first-class HTMX support through typed `hx-*` props on component props.
  Rationale: this project is intended for HTMX-heavy server-rendered applications.
  Date/Author: 2026-04-02, user.
- Decision: use `maragu.dev/gomponents v1.2.0`.
  Rationale: it is the current module path for gomponents and gives the expected API surface.
  Date/Author: 2026-04-02, Codex.
- Decision: use Tailwind 3.4.17 instead of Tailwind 4.
  Rationale: the workspace Node runtime is 18.20.8, which cannot reliably run Tailwind 4 tooling.
  Date/Author: 2026-04-02, Codex.

## Outcomes & Retrospective

The repository now contains the initial v1 library surface, a showcase app, page goldens, and Playwright coverage for interactive and HTMX-driven behavior. Validation passes with `go test ./...`, `npm run build:css`, and `npm run test:e2e`. The main follow-up risk is breadth rather than basic viability: the component catalog is implemented and exercised, but future iterations should deepen per-component unit coverage and consumer-facing documentation as the API settles.

## Context and Orientation

The library is organized around public leaf component packages under `components/`, stable HTMX props in `htmx/`, internal implementation helpers in `internal/`, theme tokens in `theme/preset.css`, the optional client runtime in `assets/ui.js`, and a runnable demo application in `examples/showcase/`. Golden tests live under `testdata/golden/`. In this repository, a “golden test” means a test that renders HTML and compares it to a checked-in reference file.

## Plan of Work

Start with the shared plumbing: class composition, accessibility helpers, HTMX attribute rendering, theme tokens, Tailwind build configuration, and the browser runtime. Build primitive components on top of those helpers, then add form/navigation composites, then finish with the interactive components. Keep the showcase app current as components land so every stable package has a visible example and the HTMX integration can be exercised through real HTTP responses instead of isolated snippets.

## Concrete Steps

Run all commands from `/Users/pme/src/pmenglund/goth`.

1. Install the toolchain and dependencies.
2. Implement the shared helpers and assets.
3. Implement the component packages in `components/`.
4. Build the showcase app in `examples/showcase/`.
5. Add Go tests and browser tests.
6. Run:

   `npm run build:css`

   `go test ./...`

   `npm run test:e2e`

## Validation and Acceptance

Success means a developer can start the showcase, open the HTMX page, trigger partial updates and overlays, and see the components continue to work without a full page reload. `go test ./...` must pass, and `npm run test:e2e` must validate keyboard and HTMX-driven behavior for the interactive widgets.

## Idempotence and Recovery

The build and test commands are safe to rerun. If generated CSS or Playwright artifacts become stale, remove only generated output such as `examples/showcase/static/ui.css`, `playwright-report`, and `test-results`, then rebuild. Golden HTML files should only change when a rendering change is intentional.

## Artifacts and Notes

Keep this file current as the implementation proceeds. When design or environment decisions change, update the `Decision Log` and `Surprises & Discoveries` sections in the same change so the next contributor can work from this file alone.

Revision note: initial implementation-oriented plan created during execution on 2026-04-02.
