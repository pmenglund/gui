package goth_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"math"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/pmenglund/goth/examples/showcase/app"
	"github.com/pmenglund/goth/internal/testutil"
)

func TestShowcaseExampleGoldens(t *testing.T) {
	for _, example := range app.TestableExamples() {
		t.Run(example.Slug, func(t *testing.T) {
			got := testutil.Render(t, example.Build())
			testutil.CompareGolden(t, exampleGoldenName(example.Slug), got)
		})
	}
}

func TestTestableExamplesCoverAtLeastFortyPercentOfPublicComponentEntryPoints(t *testing.T) {
	examples := app.TestableExamples()
	if len(examples) < 2 {
		t.Fatalf("expected at least two testable examples, got %d", len(examples))
	}

	entrypoints := componentEntryPoints(t)
	covered := map[string]struct{}{}
	missing := []string{}
	for _, example := range examples {
		for _, name := range example.Covered {
			if _, ok := entrypoints[name]; !ok {
				missing = append(missing, name)
				continue
			}
			covered[name] = struct{}{}
		}
	}

	if len(missing) > 0 {
		sort.Strings(missing)
		t.Fatalf("example registry references unknown component entrypoints: %s", strings.Join(missing, ", "))
	}

	required := int(math.Ceil(float64(len(entrypoints)) * 0.4))
	t.Logf("covered %d/%d public component entrypoints across %d examples; required >= %d", len(covered), len(entrypoints), len(examples), required)
	if len(covered) < required {
		t.Fatalf("testable examples cover %d entrypoints; need at least %d", len(covered), required)
	}
}

func componentEntryPoints(t *testing.T) map[string]struct{} {
	t.Helper()

	matches, err := filepath.Glob(filepath.Join("components", "*", "*.go"))
	if err != nil {
		t.Fatalf("glob component files: %v", err)
	}

	fset := token.NewFileSet()
	entrypoints := map[string]struct{}{}
	for _, match := range matches {
		if strings.HasSuffix(match, "_test.go") {
			continue
		}

		file, err := parser.ParseFile(fset, match, nil, parser.SkipObjectResolution)
		if err != nil {
			t.Fatalf("parse %s: %v", match, err)
		}

		componentName := filepath.Base(filepath.Dir(match))
		for _, decl := range file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Recv != nil || !ast.IsExported(fn.Name.Name) {
				continue
			}
			entrypoints[componentName+"."+fn.Name.Name] = struct{}{}
		}
	}

	return entrypoints
}

func exampleGoldenName(slug string) string {
	slug = strings.ReplaceAll(slug, "-", "_")
	return "showcase_example_" + slug + ".html"
}
