package formfield

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/pmenglund/gui/components/fielddescription"
	"github.com/pmenglund/gui/components/fielderror"
	"github.com/pmenglund/gui/components/label"
	public "github.com/pmenglund/gui/htmx"
	"github.com/pmenglund/gui/internal/a11y"
	"github.com/pmenglund/gui/internal/render"
	"github.com/pmenglund/gui/internal/tw"
)

type IDs struct {
	ControlID     string
	DescriptionID string
	ErrorID       string
}

type Props struct {
	ID          string
	Class       string
	Attributes  []g.Node
	DataTestID  string
	Label       string
	Description string
	Error       string
	Required    bool
	// Builder is the escape hatch for custom controls. Standard input-like controls
	// should prefer the typed helpers in this package so FormField can apply the
	// generated accessibility wiring automatically.
	Builder func(IDs) g.Node
	HTMX    public.Props
}

func resolveIDs(p Props) IDs {
	controlID := a11y.ID("field", p.ID, p.Label, p.DataTestID)

	ids := IDs{ControlID: controlID}
	if p.Description != "" {
		ids.DescriptionID = controlID + "-description"
	}
	if p.Error != "" {
		ids.ErrorID = controlID + "-error"
	}

	return ids
}

// FormField renders the field shell around a control. For standard inputs,
// textareas, and selects, prefer the typed helpers in this package so the
// control receives the generated ID and ARIA wiring automatically.
func FormField(p Props, children ...g.Node) g.Node {
	ids := resolveIDs(p)

	var control g.Node = g.Group(children)
	if p.Builder != nil {
		control = p.Builder(ids)
	}

	return h.Div(
		append(
			render.Attrs(
				"",
				tw.Join("grid gap-2", p.Class),
				p.DataTestID,
				p.HTMX,
				p.Attributes,
			),
			label.Label(label.Props{For: ids.ControlID, Required: p.Required}, g.Text(p.Label)),
			control,
			g.If(p.Description != "", fielddescription.FieldDescription(fielddescription.Props{ID: ids.DescriptionID}, g.Text(p.Description))),
			g.If(p.Error != "", fielderror.FieldError(fielderror.Props{ID: ids.ErrorID}, g.Text(p.Error))),
		)...,
	)
}
