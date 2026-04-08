package button

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/pmenglund/goth/components/classmode"
	public "github.com/pmenglund/goth/htmx"
	"github.com/pmenglund/goth/internal/render"
	"github.com/pmenglund/goth/internal/tw"
)

// Variant controls the button color and border treatment.
type Variant string

// Size controls the button height and padding.
type Size string

const (
	// VariantDefault renders primary button styling.
	VariantDefault Variant = "default"
	// VariantOutline renders outlined button styling.
	VariantOutline Variant = "outline"
	// VariantGhost renders transparent button styling.
	VariantGhost Variant = "ghost"
	// VariantDestructive renders destructive action button styling.
	VariantDestructive Variant = "destructive"
	// VariantSecondary renders secondary button styling.
	VariantSecondary Variant = "secondary"
)

const (
	// SizeSM renders the small button size.
	SizeSM Size = "sm"
	// SizeMD renders the medium button size.
	SizeMD Size = "md"
	// SizeLG renders the large button size.
	SizeLG Size = "lg"
)

// Props configures Button rendering.
type Props struct {
	ID         string
	Class      string
	ClassMode  classmode.ClassMode
	Attributes []g.Node
	Variant    Variant
	Size       Size
	Type       string
	Disabled   bool
	DataTestID string
	HTMX       public.Props
}

var recipe = tw.Recipe{
	Base: "inline-flex items-center justify-center gap-2 rounded-[var(--ui-radius)] border text-sm font-medium transition focus:outline-none focus-visible:ring-2 focus-visible:ring-[rgb(var(--ui-ring))] focus-visible:ring-offset-2 focus-visible:ring-offset-[rgb(var(--ui-background))]",
	Variants: map[string]string{
		string(VariantDefault):     "border-[rgb(var(--ui-primary))] bg-[rgb(var(--ui-primary))] text-[rgb(var(--ui-primary-foreground))] hover:opacity-90",
		string(VariantOutline):     "border-[rgb(var(--ui-border))] bg-[rgb(var(--ui-surface))] text-[rgb(var(--ui-foreground))] hover:bg-[rgb(var(--ui-surface-strong))]",
		string(VariantGhost):       "border-transparent bg-transparent text-[rgb(var(--ui-foreground))] hover:bg-[rgb(var(--ui-surface-strong))]",
		string(VariantDestructive): "border-[rgb(var(--ui-danger))] bg-[rgb(var(--ui-danger))] text-[rgb(var(--ui-danger-foreground))] hover:opacity-90",
		string(VariantSecondary):   "border-[rgb(var(--ui-muted))] bg-[rgb(var(--ui-surface-strong))] text-[rgb(var(--ui-foreground))] hover:border-[rgb(var(--ui-primary))]",
	},
	Sizes: map[string]string{
		string(SizeSM): "h-9 px-3 text-sm",
		string(SizeMD): "h-11 px-4 text-sm",
		string(SizeLG): "h-12 px-5 text-base",
	},
	DefaultVariant: string(VariantDefault),
	DefaultSize:    string(SizeMD),
	StateClasses: map[string]string{
		"disabled": "cursor-not-allowed opacity-60",
	},
}

// Button renders a button element with the provided styling and content.
func Button(p Props, children ...g.Node) g.Node {
	buttonType := p.Type
	if buttonType == "" {
		buttonType = "button"
	}

	attrs := render.Attrs(
		p.ID,
		tw.Classes(p.ClassMode, recipe.Class(string(p.Variant), string(p.Size), map[string]bool{"disabled": p.Disabled}), p.Class),
		p.DataTestID,
		p.HTMX,
		p.Attributes,
		h.Type(buttonType),
		g.If(p.Disabled, h.Disabled()),
	)
	return h.Button(append(attrs, children...)...)
}
