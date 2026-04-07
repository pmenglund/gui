package testutil

import (
	"bytes"
	"flag"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	g "maragu.dev/gomponents"
)

var updateGolden = flag.Bool("update", false, "update golden files")

// Render turns a gomponents node into trimmed HTML.
func Render(t *testing.T, node g.Node) string {
	t.Helper()

	var b bytes.Buffer
	if err := node.Render(&b); err != nil {
		t.Fatalf("render failed: %v", err)
	}
	return strings.TrimSpace(b.String())
}

// CompareGolden compares the rendered output against a golden file.
func CompareGolden(t *testing.T, name, got string) {
	t.Helper()

	path := filepath.Join(repoRoot(t), "testdata", "golden", name)
	if *updateGolden {
		if err := os.WriteFile(path, []byte(got+"\n"), 0o644); err != nil {
			t.Fatalf("write golden %s: %v", path, err)
		}
		return
	}

	wantBytes, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read golden %s: %v", path, err)
	}

	want := strings.TrimSpace(string(wantBytes))
	if got != want {
		t.Fatalf("golden mismatch for %s\nwant:\n%s\n\ngot:\n%s", name, want, got)
	}
}

func repoRoot(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("unable to determine caller path")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(file), "..", ".."))
}
