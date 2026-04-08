package alert

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/pmenglund/goth/components/classmode"
	public "github.com/pmenglund/goth/htmx"
	"github.com/pmenglund/goth/internal/render"
	"github.com/pmenglund/goth/internal/tw"
)

// Variant controls the alert color treatment.
type Variant string

const (
	// VariantInfo renders the default informational alert styling.
	VariantInfo Variant = "info"
	// VariantSuccess renders success alert styling.
	VariantSuccess Variant = "success"
	// VariantDanger renders danger alert styling.
	VariantDanger Variant = "danger"
)

// Props configures Alert rendering.
type Props struct {
	ID               string
	Class            string
	ClassMode        classmode.ClassMode
	Attributes       []g.Node
	DataTestID       string
	Variant          Variant
	Title            string
	Description      string
	TitleClass       string
	DescriptionClass string
	HTMX             public.Props
}

// Alert renders a status message with the provided variant and content.
func Alert(p Props, children ...g.Node) g.Node {
	variant := map[Variant]string{
		VariantInfo:    "border-[rgb(var(--ui-primary))] bg-[rgb(var(--ui-primary-foreground))]",
		VariantSuccess: "border-[rgb(var(--ui-success))] bg-[rgb(var(--ui-success-foreground))]",
		VariantDanger:  "border-[rgb(var(--ui-danger))] bg-[rgb(var(--ui-danger-foreground))]",
	}[p.Variant]
	if variant == "" {
		variant = "border-[rgb(var(--ui-primary))] bg-[rgb(var(--ui-primary-foreground))]"
	}

	content := []g.Node{}
	if p.Title != "" {
		content = append(content, h.H4(h.Class(tw.Join("font-semibold", p.TitleClass)), g.Text(p.Title)))
	}
	if p.Description != "" {
		content = append(content, h.P(h.Class(tw.Join("text-sm text-[rgb(var(--ui-muted-foreground))]", p.DescriptionClass)), g.Text(p.Description)))
	}
	content = append(content, children...)

	return h.Div(append(
		render.Attrs(
			p.ID,
			tw.Classes(p.ClassMode, tw.Join("grid gap-2 rounded-[var(--ui-radius)] border p-4", variant), p.Class),
			p.DataTestID,
			p.HTMX,
			p.Attributes,
			h.Role("alert"),
		),
		content...,
	)...)
}
