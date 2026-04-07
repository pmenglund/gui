package formfield

import (
	g "maragu.dev/gomponents"

	inputui "github.com/pmenglund/gui/components/input"
	selectui "github.com/pmenglund/gui/components/select"
	textareaui "github.com/pmenglund/gui/components/textarea"
	"github.com/pmenglund/gui/internal/a11y"
)

// Input renders a FormField-backed text input with generated ID and ARIA wiring.
func Input(field Props, control inputui.Props) g.Node {
	field.Builder = func(ids IDs) g.Node {
		control.ID = ids.ControlID
		control.DescribedBy = a11y.DescribedBy(ids.DescriptionID, ids.ErrorID, control.DescribedBy)
		control.Invalid = control.Invalid || ids.ErrorID != ""
		return inputui.Input(control)
	}

	return FormField(field)
}

// Textarea renders a FormField-backed textarea with generated ID and ARIA wiring.
func Textarea(field Props, control textareaui.Props) g.Node {
	field.Builder = func(ids IDs) g.Node {
		control.ID = ids.ControlID
		control.DescribedBy = a11y.DescribedBy(ids.DescriptionID, ids.ErrorID, control.DescribedBy)
		control.Invalid = control.Invalid || ids.ErrorID != ""
		return textareaui.Textarea(control)
	}

	return FormField(field)
}

// Select renders a FormField-backed select with generated ID and ARIA wiring.
func Select(field Props, control selectui.Props) g.Node {
	field.Builder = func(ids IDs) g.Node {
		control.ID = ids.ControlID
		control.DescribedBy = a11y.DescribedBy(ids.DescriptionID, ids.ErrorID, control.DescribedBy)
		control.Invalid = control.Invalid || ids.ErrorID != ""
		return selectui.Select(control)
	}

	return FormField(field)
}
