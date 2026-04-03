package htmx

// Props holds stable HTMX attributes that can be attached to components.
type Props struct {
	Get       string
	Post      string
	Put       string
	Patch     string
	Delete    string
	Trigger   string
	Target    string
	Swap      string
	Select    string
	Include   string
	Indicator string
	PushURL   string
	Confirm   string
	Encoding  string
	Values    string
	Disabled  bool
}
