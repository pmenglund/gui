package tw

import (
	"strings"

	"github.com/pmenglund/goth/components/classmode"
)

// Join deduplicates and joins Tailwind class fragments.
func Join(parts ...string) string {
	seen := map[string]bool{}
	out := make([]string, 0, len(parts))
	for _, part := range parts {
		for _, className := range strings.Fields(part) {
			if !seen[className] {
				seen[className] = true
				out = append(out, className)
			}
		}
	}
	return strings.Join(out, " ")
}

// Classes composes root classes according to the public class mode.
func Classes(mode classmode.ClassMode, defaults, caller string) string {
	if mode == classmode.ClassReplace {
		return caller
	}
	return Join(defaults, caller)
}

// When returns the class string only when the condition is true.
func When(ok bool, className string) string {
	if ok {
		return className
	}
	return ""
}

// Recipe assembles stable class recipes from base, variant, size, and state flags.
type Recipe struct {
	Base           string
	Variants       map[string]string
	Sizes          map[string]string
	DefaultVariant string
	DefaultSize    string
	StateClasses   map[string]string
}

// Class resolves a class list from the recipe.
func (r Recipe) Class(variant, size string, flags map[string]bool, extra ...string) string {
	if variant == "" || r.Variants[variant] == "" {
		variant = r.DefaultVariant
	}
	if size == "" || r.Sizes[size] == "" {
		size = r.DefaultSize
	}

	parts := []string{r.Base, r.Variants[variant], r.Sizes[size]}
	for name, on := range flags {
		if on {
			parts = append(parts, r.StateClasses[name])
		}
	}
	parts = append(parts, extra...)
	return Join(parts...)
}
