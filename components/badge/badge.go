package badge

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/pmenglund/goth/components/classmode"
	public "github.com/pmenglund/goth/htmx"
	"github.com/pmenglund/goth/internal/render"
	"github.com/pmenglund/goth/internal/tw"
)

// Variant controls the badge color treatment.
type Variant string

const (
	// VariantDefault renders primary badge styling.
	VariantDefault Variant = "default"
	// VariantMuted renders muted badge styling.
	VariantMuted Variant = "muted"
	// VariantSuccess renders success badge styling.
	VariantSuccess Variant = "success"
	// VariantDanger renders danger badge styling.
	VariantDanger Variant = "danger"
)

// Props configures Badge rendering.
type Props struct {
	ID         string
	Class      string
	ClassMode  classmode.ClassMode
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
			tw.Classes(p.ClassMode, tw.Join("inline-flex items-center rounded-full border px-2.5 py-1 text-xs font-semibold uppercase tracking-[0.15em]", variant), p.Class),
			p.DataTestID,
			public.Props{},
			p.Attributes,
		),
		children...,
	)...)
}
