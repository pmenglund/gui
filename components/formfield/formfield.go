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
	Builder     func(IDs) g.Node
	HTMX        public.Props
}

func FormField(p Props, children ...g.Node) g.Node {
	controlID := a11y.ID("field", p.ID, p.Label, p.DataTestID)
	descriptionID := ""
	errorID := ""
	if p.Description != "" {
		descriptionID = controlID + "-description"
	}
	if p.Error != "" {
		errorID = controlID + "-error"
	}

	var control g.Node = g.Group(children)
	if p.Builder != nil {
		control = p.Builder(IDs{
			ControlID:     controlID,
			DescriptionID: descriptionID,
			ErrorID:       errorID,
		})
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
			label.Label(label.Props{For: controlID, Required: p.Required}, g.Text(p.Label)),
			control,
			g.If(p.Description != "", fielddescription.FieldDescription(fielddescription.Props{ID: descriptionID}, g.Text(p.Description))),
			g.If(p.Error != "", fielderror.FieldError(fielderror.Props{ID: errorID}, g.Text(p.Error))),
		)...,
	)
}
