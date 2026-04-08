package formfield_test

import (
	"strings"
	"testing"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/pmenglund/goth/components/formfield"
	"github.com/pmenglund/goth/internal/testutil"
)

func TestFormFieldBuilderReceivesGeneratedWiringIDs(t *testing.T) {
	t.Parallel()

	got := testutil.Render(t, formfield.FormField(formfield.Props{
		ID:          "work-email",
		Label:       "Work email",
		Description: "Use your company address.",
		Error:       "Email is required.",
		Builder: func(ids formfield.IDs) g.Node {
			return h.Input(
				h.ID(ids.ControlID),
				h.Aria("describedby", ids.DescriptionID+" "+ids.ErrorID),
			)
		},
	}))

	if !strings.Contains(got, `for="work-email"`) {
		t.Fatalf("form field label missing generated control id: %s", got)
	}
	if !strings.Contains(got, `<input id="work-email" aria-describedby="work-email-description work-email-error">`) {
		t.Fatalf("form field builder control missing generated aria wiring: %s", got)
	}
}
