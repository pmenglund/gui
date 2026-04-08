package badge_test

import (
	"strings"
	"testing"

	g "maragu.dev/gomponents"

	"github.com/pmenglund/goth/components/badge"
	"github.com/pmenglund/goth/components/classmode"
	"github.com/pmenglund/goth/internal/testutil"
)

func TestBadgeRootClassReplaceAllowsNoClassAttribute(t *testing.T) {
	t.Parallel()

	got := testutil.Render(t, badge.Badge(badge.Props{
		ClassMode: classmode.ClassReplace,
	}, g.Text("Stable")))
	if strings.Contains(got, `class=`) {
		t.Fatalf("badge replacement with empty caller rendered class attribute: %s", got)
	}
}
