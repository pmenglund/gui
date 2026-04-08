package tw

import (
	"testing"

	"github.com/pmenglund/goth/components/classmode"
)

func TestClassesMergesByDefault(t *testing.T) {
	t.Parallel()

	got := Classes(classmode.ClassMerge, "rounded p-4", "p-4 text-sm")
	want := "rounded p-4 text-sm"
	if got != want {
		t.Fatalf("Classes merge = %q, want %q", got, want)
	}
}

func TestClassesReplacesWithCallerClasses(t *testing.T) {
	t.Parallel()

	got := Classes(classmode.ClassReplace, "rounded p-4", "host-card")
	want := "host-card"
	if got != want {
		t.Fatalf("Classes replace = %q, want %q", got, want)
	}
}

func TestClassesReplaceAllowsEmptyClass(t *testing.T) {
	t.Parallel()

	got := Classes(classmode.ClassReplace, "rounded p-4", "")
	if got != "" {
		t.Fatalf("Classes replace with empty caller = %q, want empty", got)
	}
}
