package alert

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	public "github.com/pmenglund/goth/htmx"
	"github.com/pmenglund/goth/internal/render"
	"github.com/pmenglund/goth/internal/tw"
)

type Variant string

const (
	VariantInfo    Variant = "info"
	VariantSuccess Variant = "success"
	VariantDanger  Variant = "danger"
)

type Props struct {
	ID          string
	Class       string
	Attributes  []g.Node
	DataTestID  string
	Variant     Variant
	Title       string
	Description string
	HTMX        public.Props
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
		content = append(content, h.H4(h.Class("font-semibold"), g.Text(p.Title)))
	}
	if p.Description != "" {
		content = append(content, h.P(h.Class("text-sm text-[rgb(var(--ui-muted-foreground))]"), g.Text(p.Description)))
	}
	content = append(content, children...)

	return h.Div(append(
		render.Attrs(
			p.ID,
			tw.Join("grid gap-2 rounded-[var(--ui-radius)] border p-4", variant, p.Class),
			p.DataTestID,
			p.HTMX,
			p.Attributes,
			h.Role("alert"),
		),
		content...,
	)...)
}
