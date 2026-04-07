package badge

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	public "github.com/pmenglund/gui/htmx"
	"github.com/pmenglund/gui/internal/render"
	"github.com/pmenglund/gui/internal/tw"
)

type Variant string

const (
	VariantDefault Variant = "default"
	VariantMuted   Variant = "muted"
	VariantSuccess Variant = "success"
	VariantDanger  Variant = "danger"
)

type Props struct {
	ID         string
	Class      string
	Attributes []g.Node
	DataTestID string
	Variant    Variant
}

// Badge renders a compact status label with the provided content.
func Badge(p Props, children ...g.Node) g.Node {
	variant := map[Variant]string{
		VariantDefault: "border-[rgb(var(--ui-primary))] bg-[rgb(var(--ui-primary))] text-[rgb(var(--ui-primary-foreground))]",
		VariantMuted:   "border-[rgb(var(--ui-muted))] bg-[rgb(var(--ui-surface-strong))] text-[rgb(var(--ui-muted-foreground))]",
		VariantSuccess: "border-[rgb(var(--ui-success))] bg-[rgb(var(--ui-success))] text-[rgb(var(--ui-success-foreground))]",
		VariantDanger:  "border-[rgb(var(--ui-danger))] bg-[rgb(var(--ui-danger))] text-[rgb(var(--ui-danger-foreground))]",
	}[p.Variant]
	if variant == "" {
		variant = map[Variant]string{
			VariantDefault: "border-[rgb(var(--ui-primary))] bg-[rgb(var(--ui-primary))] text-[rgb(var(--ui-primary-foreground))]",
		}[VariantDefault]
	}
	return h.Span(append(
		render.Attrs(
			p.ID,
			tw.Join("inline-flex items-center rounded-full border px-2.5 py-1 text-xs font-semibold uppercase tracking-[0.15em]", variant, p.Class),
			p.DataTestID,
			public.Props{},
			p.Attributes,
		),
		children...,
	)...)
}
