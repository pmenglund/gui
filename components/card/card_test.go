package card_test

import (
	"strings"
	"testing"

	"github.com/pmenglund/goth/components/card"
	"github.com/pmenglund/goth/internal/testutil"
)

func TestCardSlotClassesRenderOnGeneratedTitleAndDescription(t *testing.T) {
	t.Parallel()

	got := testutil.Render(t, card.Card(card.Props{
		Title:            "Release queue",
		Description:      "Pending rollouts.",
		TitleClass:       "host-title",
		DescriptionClass: "host-description",
	}))
	if !strings.Contains(got, `<h3 class="text-lg font-semibold host-title">Release queue</h3>`) {
		t.Fatalf("card title missing slot class: %s", got)
	}
	if !strings.Contains(got, `<p class="text-sm text-[rgb(var(--ui-muted-foreground))] host-description">Pending rollouts.</p>`) {
		t.Fatalf("card description missing slot class: %s", got)
	}
}
