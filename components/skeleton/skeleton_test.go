package skeleton_test

import (
	"strings"
	"testing"

	"github.com/pmenglund/goth/components/skeleton"
	"github.com/pmenglund/goth/internal/testutil"
)

func TestSkeletonNegativeCountDefaultsToOne(t *testing.T) {
	t.Parallel()

	got := testutil.Render(t, skeleton.Skeleton(skeleton.Props{Count: -3}))
	if !strings.Contains(got, `data-count="1"`) {
		t.Fatalf("negative skeleton count did not default to one: %s", got)
	}
}
