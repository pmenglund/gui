package htmx

import (
	"bytes"
	"net/http"
	"time"

	_ "embed"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

const (
	// Version is the bundled HTMX runtime version.
	Version = "2.0.8"

	// ScriptPath is the default path for the bundled HTMX runtime.
	ScriptPath = "/goth/htmx-" + Version + ".min.js"
)

// ScriptProps configures the bundled HTMX script tag.
type ScriptProps struct {
	Src        string
	Defer      bool
	Attributes []g.Node
}

//go:embed htmx.min.js
var htmxMinJS []byte

// Script renders a script tag for the bundled HTMX runtime.
func Script(p ScriptProps) g.Node {
	src := p.Src
	if src == "" {
		src = ScriptPath
	}

	attrs := []g.Node{h.Src(src)}
	if p.Defer {
		attrs = append(attrs, h.Defer())
	}
	attrs = append(attrs, p.Attributes...)

	return h.Script(attrs...)
}

// Handler returns an HTTP handler that serves the bundled HTMX runtime.
func Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
		w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
		http.ServeContent(w, r, "htmx.min.js", time.Time{}, bytes.NewReader(htmxMinJS))
	})
}
