package classmode

// ClassMode controls how component default classes combine with caller classes.
type ClassMode string

const (
	// ClassMerge appends caller classes to component defaults.
	ClassMerge ClassMode = ""
	// ClassReplace renders only caller classes for the component root.
	ClassReplace ClassMode = "replace"
)
