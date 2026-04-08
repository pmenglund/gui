package htmx_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	g "maragu.dev/gomponents"

	"github.com/pmenglund/goth/htmx"
	"github.com/pmenglund/goth/internal/testutil"
)

func TestScriptRendersDefaultPath(t *testing.T) {
	t.Parallel()

	got := testutil.Render(t, htmx.Script(htmx.ScriptProps{}))
	want := `<script src="/goth/htmx-2.0.8.min.js"></script>`
	if got != want {
		t.Fatalf("script rendered unexpected markup\nwant:\n%s\n\ngot:\n%s", want, got)
	}
}

func TestScriptRendersCustomAttributesInStableOrder(t *testing.T) {
	t.Parallel()

	got := testutil.Render(t, htmx.Script(htmx.ScriptProps{
		Src:   "/assets/htmx.js",
		Defer: true,
		Attributes: []g.Node{
			g.Attr("integrity", "sha256-test"),
			g.Attr("crossorigin", "anonymous"),
		},
	}))
	want := `<script src="/assets/htmx.js" defer integrity="sha256-test" crossorigin="anonymous"></script>`
	if got != want {
		t.Fatalf("script rendered unexpected markup\nwant:\n%s\n\ngot:\n%s", want, got)
	}
}

func TestHandlerServesEmbeddedRuntime(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(http.MethodGet, htmx.ScriptPath, nil)
	rr := httptest.NewRecorder()
	htmx.Handler().ServeHTTP(rr, req)
	res := rr.Result()
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("read handler body: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("handler returned status %d, body: %s", res.StatusCode, string(body))
	}
	if got := res.Header.Get("Content-Type"); got != "application/javascript; charset=utf-8" {
		t.Fatalf("handler returned content type %q", got)
	}
	if got := res.Header.Get("Cache-Control"); got != "public, max-age=31536000, immutable" {
		t.Fatalf("handler returned cache control %q", got)
	}
	if !strings.Contains(string(body), `version:"`+htmx.Version+`"`) {
		t.Fatalf("handler body did not contain bundled htmx version %q", htmx.Version)
	}
}
