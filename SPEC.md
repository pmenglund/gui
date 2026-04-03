# Gomponents Tailwind UI Library Specification

Status: Draft v1 (Go-specific)

Purpose: Define a reusable UI component library for Go web applications built with gomponents and Tailwind CSS. The library provides higher-level, production-oriented components similar in spirit to templUI, while preserving gomponents' pure-Go authoring model.

## 1. Problem Statement

Go teams using gomponents can already build reusable HTML in pure Go, and gomponents pairs naturally with Tailwind CSS. What is missing in many projects is not the ability to emit HTML, but a consistent, higher-level UI layer that provides:

- ready-made visual primitives with coherent defaults
- typed component APIs instead of ad hoc class strings everywhere
- a stable styling contract for variants, sizes, spacing, and color tokens
- accessible patterns for forms, navigation, and overlays
- a disciplined interactivity model for components that need JavaScript
- a shared package layout and testing strategy that can scale across multiple apps

Without such a library, teams often accumulate one-off wrappers around `Div`, `Button`, `Input`, and `A`, with duplicated Tailwind classes, inconsistent accessibility, and no clear boundary between design-system code and app-specific UI code.

This specification defines a library that solves those problems while keeping the implementation grounded in gomponents:

- Components are authored in Go.
- Public constructors return `gomponents.Node`.
- Styling is expressed primarily through Tailwind utility classes plus a small token layer.
- Interactive behavior is progressive-enhancement-oriented and optional.
- Consumers can compose, wrap, or fork components without adopting a new template language or code generator.

Important boundary:

- This project is a UI component library and design-system runtime for server-rendered Go apps.
- It is not a SPA framework, virtual DOM, or general-purpose frontend build system.
- It may ship small JavaScript helpers for specific interactive components, but JavaScript is an implementation detail rather than the core programming model.

## 2. Goals and Non-Goals

### 2.1 Goals

- Provide a coherent catalog of reusable UI components for gomponents.
- Standardize typed props, variants, sizes, states, and escape hatches.
- Establish a Tailwind-based design token system that is easy to override.
- Make accessible markup and keyboard behavior the default rather than an afterthought.
- Support server-rendered HTML with optional progressive enhancement.
- Keep the public API ergonomic for Go developers and compatible with standard Go tooling.
- Minimize runtime dependencies and avoid introducing a template compiler.
- Enable incremental adoption inside existing gomponents applications.
- Support component-level testing, accessibility verification, and visual review.

### 2.2 Non-Goals

- Reproducing templUI component-for-component or API-for-API.
- Providing a React-like client state model.
- Requiring a JavaScript framework such as React, Vue, Alpine, or similar.
- Requiring a code-copying CLI in v1.
- Serving as a general CSS framework independent of Go.
- Solving product-specific layout, branding, or information architecture for consuming apps.
- Abstracting away all HTML knowledge from consumers.

## 3. Design Principles

1. **Pure Go first**
   - Component authoring and composition happen in Go.
   - The library should feel like an idiomatic extension of gomponents rather than a second templating system.

2. **Styling by contract**
   - Tailwind utility composition is centralized inside the library.
   - Tokens, variants, and sizes are explicit, documented, and testable.

3. **Accessible by default**
   - Components must emit semantic HTML and correct ARIA relationships where applicable.
   - Interactive widgets must define keyboard, focus, and dismissal behavior.

4. **Progressive enhancement**
   - Static components must work with no JavaScript.
   - Interactive components may enhance server-rendered HTML with a small JS runtime.

5. **Escape hatches without chaos**
   - Every component should expose controlled extensibility via `Class`, `Attributes`, slot content, and composition.
   - Escape hatches should not undermine accessibility or core layout invariants.

6. **Composable over magical**
   - Prefer explicit props and subcomponents over hidden global behavior.
   - Avoid code generation and hidden runtime registration.

## 4. System Overview

### 4.1 Main Components

1. `Theme Layer`
   - Defines CSS variables and Tailwind theme extensions.
   - Owns semantic tokens such as surface, foreground, border, accent, destructive, and muted.

2. `Class Composition Layer`
   - Builds deterministic class lists from base styles, variants, sizes, states, and overrides.
   - Prevents scattered duplication of Tailwind strings across components.

3. `Component Primitives`
   - Stateless or minimally stateful building blocks such as button, input, card, badge, label, separator, and spinner.

4. `Composite Components`
   - Higher-level compositions such as form field, navbar, table, tabs, pagination, and breadcrumbs.

5. `Interactive Runtime`
   - Small optional JavaScript bundle for dialogs, dropdowns, popovers, toasts, and other overlay-style widgets.
   - Uses data attributes and event delegation rather than framework-specific hydration.

6. `Accessibility Utilities`
   - Helpers for IDs, ARIA relationships, visually hidden content, and focus management hooks.

7. `Testing Surface`
   - Golden HTML tests, runtime behavior tests, accessibility checks, and example/showcase pages.

### 4.2 Abstraction Levels

The library should remain maintainable by separating these layers:

1. `Public API Layer`
   - Stable Go constructors, props types, enums, and helper functions.

2. `Styling Layer`
   - Token definitions, class recipes, and Tailwind integration.

3. `Rendering Layer`
   - gomponents-based HTML assembly.

4. `Behavior Layer`
   - Optional JavaScript and state attributes for interactive widgets.

5. `Documentation Layer`
   - Example pages, API docs, component usage, and migration guidance.

### 4.3 External Dependencies

Required:

- Go toolchain
- gomponents
- Tailwind CSS build tooling

Optional:

- htmx or other server-driven enhancement libraries used by consuming apps
- a tiny bundled JS file for interactive components
- visual regression tooling and browser automation for CI

## 5. Core Domain Model

### 5.1 Entities

#### 5.1.1 Component

A reusable UI constructor that returns `gomponents.Node`.

Fields, logically:

- `name` (string)
- `category` (primitive, composite, interactive)
- `props_type` (Go type)
- `slots` (set of named content insertion points, if any)
- `requires_js` (bool)
- `a11y_contract` (documented required semantics)

#### 5.1.2 Props

A typed configuration object for a component.

Common fields available to many components:

- `ID` (string)
- `Class` (string)
- `Attributes` (`[]gomponents.Node`)
- `Variant` (enum-like string type)
- `Size` (enum-like string type)
- `Disabled` (bool)
- `DataTestID` (string, optional convenience)

Component-specific props may extend this set.

#### 5.1.3 Slot

A structured content region within a component.

Examples:

- `LeadingIcon`
- `TrailingIcon`
- `Header`
- `Footer`
- `Title`
- `Description`
- `Body`
- `Trigger`
- `Content`

Slots may be represented as explicit props, subcomponent helpers, or variadic child content depending on the component.

#### 5.1.4 Variant Recipe

A deterministic mapping from semantic choices to Tailwind classes.

Fields:

- `base_classes` (string)
- `variant_classes` (map)
- `size_classes` (map)
- `state_classes` (map)
- `compound_rules` (optional list)

#### 5.1.5 Theme Token

A semantic design token backed by CSS variables and Tailwind usage.

Examples:

- `--ui-background`
- `--ui-foreground`
- `--ui-muted`
- `--ui-border`
- `--ui-primary`
- `--ui-primary-foreground`
- `--ui-danger`
- `--ui-ring`
- `--ui-radius`

#### 5.1.6 Interactive Controller

The runtime behavior definition for a JS-enhanced component.

Fields:

- `controller_name` (string)
- `data_attributes` (set)
- `custom_events` (set, optional)
- `keyboard_contract` (documented behavior)
- `focus_contract` (documented behavior)
- `dismissal_contract` (documented behavior)

### 5.2 Normalization Rules

- Variant and size values must be lowercase string constants.
- Unknown variant or size values must resolve to documented defaults rather than silently breaking output.
- `Class` values are appended after internal classes so consumers can extend or override when necessary.
- `Attributes` are appended last unless a component documents a stronger precedence rule for safety.
- Components should generate deterministic IDs only when the caller does not provide an `ID` and deterministic generation is needed for accessibility relationships.

## 6. Public API Specification

### 6.1 Constructor Shape

The preferred v1 constructor pattern is:

```go
func Button(p ButtonProps, children ...g.Node) g.Node
```

General rules:

- Public constructors must return `gomponents.Node`.
- Constructors should accept a single props struct plus optional child content.
- Constructors must be safe to call with the zero value of the props struct.
- Constructors must not panic on invalid user input; they should fall back to documented defaults where practical.

### 6.2 Props Conventions

Each component props type should follow these conventions:

- use explicit Go fields instead of `map[string]any`
- use enum-like custom string types for `Variant`, `Size`, and similar constrained values
- reserve `Class` for utility-class overrides
- reserve `Attributes` for arbitrary HTML attrs and `data-*` hooks
- prefer booleans for binary behavior flags
- prefer dedicated slot props over magic child indexing when structure matters

Example:

```go
type ButtonVariant string

type ButtonSize string

const (
    ButtonVariantDefault ButtonVariant = "default"
    ButtonVariantOutline ButtonVariant = "outline"
    ButtonVariantGhost   ButtonVariant = "ghost"
)

const (
    ButtonSizeSM ButtonSize = "sm"
    ButtonSizeMD ButtonSize = "md"
    ButtonSizeLG ButtonSize = "lg"
)

type ButtonProps struct {
    ID         string
    Class      string
    Attributes []g.Node
    Variant    ButtonVariant
    Size       ButtonSize
    Type       string
    Disabled   bool
    Href       string
}
```

### 6.3 Escaping and HTML Safety

- Components must default to safe, escaped content.
- Raw HTML insertion must only occur through explicit gomponents raw-node usage supplied by the caller.
- Components must not internally concatenate unsafe HTML strings.

### 6.4 Rendering Semantics

- Components should emit minimal, semantic HTML.
- Components must avoid unnecessary wrapper elements unless they are required for semantics, styling, or behavior.
- Anchor-vs-button behavior must be explicit. For example, a button component with `Href` should render an anchor-like control only if that behavior is documented.
- Form controls must preserve native HTML behavior as much as possible.

## 7. Styling Specification

### 7.1 Tailwind Contract

The library uses Tailwind CSS as the primary styling mechanism.

Requirements:

- Tailwind content scanning must include Go source files that contain component class strings.
- Internal class composition should prefer complete class tokens rather than string fragments that Tailwind cannot detect reliably.
- Semantic theme values should map to CSS variables and be referenced through the Tailwind theme configuration where appropriate.

### 7.2 Theme Contract

The project must ship a documented base theme with:

- light mode tokens
- dark mode tokens
- border radius scale
- spacing expectations for component sizes
- focus ring treatment
- destructive and success emphasis colors

The theme contract must support consumer overrides through one or both of:

- CSS variable overrides
- a published Tailwind preset or config snippet

### 7.3 Class Composition

The library should define an internal helper package for deterministic class assembly.

Capabilities should include:

- joining class lists
- conditional class inclusion
- base + variant + size + state composition
- compound variant rules
- deduplication where practical

The class composition layer is internal and not part of the stable public API in v1.

### 7.4 Override Policy

Consumers must be able to customize components without forking by using:

- `Class`
- `Attributes`
- wrapper components in their own app code

Consumers should not be expected to depend on internal class names or DOM nesting beyond what the documentation marks as stable.

## 8. Accessibility Specification

### 8.1 General Requirements

All shipped components must have documented accessibility behavior.

Minimum requirements:

- semantic HTML whenever native elements can express the behavior
- visible focus styles
- sufficient color contrast in the base theme
- proper `label`/`input` association
- ARIA only where native semantics are insufficient
- keyboard support for all interactive controls

### 8.2 Forms

Form components must support:

- labels
- descriptions or helper text
- invalid/error state messaging
- disabled state
- required state
- accessible association via `for`, `id`, `aria-describedby`, and `aria-invalid` where applicable

### 8.3 Overlay Components

Dialog, popover, dropdown, and similar components must define:

- trigger semantics
- open/close state attributes
- ESC behavior
- click-outside behavior where applicable
- focus management and restoration
- inert or hidden treatment for offscreen content

If a component cannot meet an acceptable accessibility bar in v1, it must not ship as stable.

## 9. Interactivity Specification

### 9.1 Runtime Model

Interactive components may rely on a small optional JS bundle.

Requirements:

- JavaScript must be framework-agnostic.
- Controllers should bind using `data-ui-*` attributes.
- The bundle should use event delegation where practical.
- Components must degrade gracefully when JS is absent, unless the component is inherently interactive and clearly documented as JS-required.

### 9.2 Stable Runtime Hooks

The following are stable hook categories in v1:

- `data-ui-controller`
- `data-ui-trigger`
- `data-ui-content`
- `data-ui-state`
- `data-ui-target`

Specific components may define additional documented attributes.

### 9.3 Event Model

Interactive controllers may emit custom DOM events for observability and integration, for example:

- `ui:open`
- `ui:close`
- `ui:toggle`

Custom events are optional but, if shipped, must be documented and semantically stable.

### 9.4 Preferred v1 Interactive Set

The initial stable interactive set should be narrow:

- Dialog
- Dropdown Menu
- Tabs
- Toast

Popover, combobox, datepicker, and command palette should remain experimental until accessibility and maintenance costs are understood.

## 10. Component Catalog

### 10.1 Required v1 Primitive Components

The library should ship the following as stable in v1:

- Button
- IconButton
- Input
- Textarea
- Label
- Checkbox
- RadioGroup
- Switch
- Select
- Card
- Badge
- Alert
- Separator
- Avatar
- Spinner
- Skeleton

### 10.2 Required v1 Composite Components

The library should ship the following as stable in v1:

- FormField
- FieldError
- FieldDescription
- Navbar
- Breadcrumbs
- Tabs
- Table
- EmptyState
- Pagination

### 10.3 Interactive Components

The following may ship in v1 if they meet the accessibility and runtime requirements:

- Dialog
- DropdownMenu
- Toast
- Sheet

### 10.4 Experimental Components

These should remain explicitly experimental until proven in production:

- Popover
- Combobox
- DatePicker
- CommandPalette
- RichTextEditor wrappers
- Complex charts

## 11. Repository and Package Layout

Recommended layout:

```text
/
  SPEC.md
  go.mod
  tailwind.config.js
  package.json                # only if needed for Tailwind/tooling
  assets/
    ui.js                     # optional interactive runtime
    ui.css                    # base tokens / layer additions
  components/
    button/
    card/
    input/
    label/
    dialog/
    dropdown/
    form/
    table/
    toast/
  theme/
    tokens.go                 # typed token metadata if useful
    preset.css                # CSS variables and base theme
  internal/
    tw/
    a11y/
    runtime/
    render/
  examples/
    showcase/
  testdata/
    golden/
```

Package rules:

- Each component package should expose a small public surface.
- Internal helpers must stay under `internal/`.
- The `examples/showcase` app should exercise all stable components.
- The project may provide a convenience umbrella package later, but leaf packages are preferred in v1 for explicit imports.

## 12. Testing and Quality Requirements

### 12.1 Unit and Golden Tests

Each stable component must have:

- constructor behavior tests
- HTML golden tests for representative variants and states
- tests covering accessibility attributes and ID wiring

### 12.2 Browser and Behavior Tests

Interactive components must have browser-level tests covering:

- keyboard navigation
- focus movement and restoration
- open/close behavior
- click-away handling
- ESC dismissal

### 12.3 Accessibility Review

The project should run automated accessibility checks in CI where practical and require manual review for complex widgets.

### 12.4 Visual Review

A showcase or story-like example surface must exist so maintainers can inspect components across themes, variants, and states.

## 13. Versioning and Stability Policy

### 13.1 Semantic Versioning

The project must use semantic versioning.

### 13.2 Stable API Surface

The following are considered public API in v1:

- exported constructors
- exported props types
- exported variant and size constants
- documented data attributes and custom events
- documented theme tokens intended for consumer override

### 13.3 Unstable Surface

The following are not stable in v1 unless explicitly documented otherwise:

- internal package layout
- internal class composition functions
- exact DOM nesting not required by the documented contract
- internal helper CSS class names
- experimental component APIs

## 14. Consumer Integration Contract

### 14.1 Minimum Consumer Requirements

Consumers must:

- include the library's Tailwind content paths in their Tailwind scan configuration or otherwise ensure classes are not purged
- include the base theme CSS variables
- include the JS runtime when using interactive components that require it

### 14.2 Incremental Adoption

The library must be usable incrementally.

That means:

- components can coexist with raw gomponents code
- consumers can wrap components inside app-specific functions
- consumers do not need to restructure their entire rendering layer to adopt a few components

### 14.3 Framework Interop

The library should remain compatible with:

- standard `net/http`
- server-side rendering stacks built around gomponents
- htmx-enhanced server flows
- custom CSS extensions layered on top of Tailwind

## 15. Security and Trust Posture

### 15.1 HTML Safety

The library must default to escaped HTML output and avoid introducing raw HTML injection points beyond the explicit mechanisms already exposed by gomponents.

### 15.2 CSP Posture

The preferred runtime posture is CSP-friendly:

- avoid inline script requirements where possible
- keep runtime JS bundle external and deterministic
- document any required nonce or policy considerations if inline behavior is ever introduced

### 15.3 Dependency Discipline

The project should minimize frontend dependencies.

- Tailwind is expected.
- A tiny JS runtime is acceptable.
- Heavy client frameworks are out of scope.

## 16. Reference Implementation Guidance

A reference implementation for v1 should prioritize the following order:

### Phase 1: Foundation

- theme tokens
- class composition helpers
- Button, Input, Textarea, Label, Card, Badge, Alert, Separator
- example showcase
- golden test harness

### Phase 2: Forms and Navigation

- Checkbox, RadioGroup, Switch, Select
- FormField, FieldError, FieldDescription
- Navbar, Breadcrumbs, Table, Pagination, EmptyState

### Phase 3: Narrow Interactive Layer

- JS runtime foundation
- Dialog
- DropdownMenu
- Toast
- browser behavior tests

### Phase 4: Hardening

- dark mode polish
- accessibility audit
- versioned migration notes
- API cleanup before v1.0.0

## 17. Open Questions

These items should be resolved before declaring the specification complete:

1. Should the project publish a single umbrella import path in addition to per-component packages?
2. Should `Href` on button-like components be supported directly or separated into explicit link/button primitives?
3. Should the project ship icons or integrate with an external icon package only?
4. Should the runtime JS be authored by hand or generated from a tiny controller registry?
5. Should toasts and dialogs be portaled or rendered in-place in v1?
6. How much DOM structure should be considered stable for consumer CSS targeting?
7. Should a code-copying CLI ever exist, or should the project remain import-only?

## 18. Bottom Line

This project should be treated as a **gomponents-native design system and component library**, not just a bag of styled helpers.

Its success depends less on whether HTML can be rendered from Go and more on whether it can establish:

- a stable public API
- a coherent Tailwind styling contract
- accessible defaults
- a disciplined interactive runtime
- a testing and documentation culture strong enough to keep the library trustworthy over time

The recommended v1 scope is deliberately narrow. A smaller, well-specified set of primitives and a few high-quality interactive components is more valuable than a large but inconsistent catalog.
