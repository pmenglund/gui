package a11y

import (
	"fmt"
	"hash/fnv"
	"strings"
	"unicode"
)

// ID returns the provided id, or a deterministic id derived from the seeds.
func ID(prefix, provided string, seeds ...string) string {
	if provided != "" {
		return provided
	}
	base := Slug(seeds...)
	if base == "" {
		h := fnv.New32a()
		_, _ = h.Write([]byte(strings.Join(seeds, "|")))
		base = fmt.Sprintf("%x", h.Sum32())
	}
	return prefix + "-" + base
}

// Slug builds a lowercase identifier fragment from human-readable text.
func Slug(parts ...string) string {
	var b strings.Builder
	lastDash := false
	for _, part := range parts {
		for _, r := range strings.ToLower(strings.TrimSpace(part)) {
			switch {
			case unicode.IsLetter(r) || unicode.IsDigit(r):
				b.WriteRune(r)
				lastDash = false
			case !lastDash:
				b.WriteByte('-')
				lastDash = true
			}
		}
		if b.Len() > 0 && !lastDash {
			b.WriteByte('-')
			lastDash = true
		}
	}
	return strings.Trim(b.String(), "-")
}

// DescribedBy joins non-empty ids for aria-describedby.
func DescribedBy(ids ...string) string {
	out := make([]string, 0, len(ids))
	for _, id := range ids {
		id = strings.TrimSpace(id)
		if id != "" {
			out = append(out, id)
		}
	}
	return strings.Join(out, " ")
}
