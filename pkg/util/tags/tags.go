package utiltags

// Tags set tags and label
type Tags struct {
	KV    map[string]string
	Label string // format: version:1.0|app:mall|env:prod
}

// NewDefaultTag returns Tags with version and appID
func NewDefaultTag(version, appID string) Tags { _ = "STUB: not implemented"; return *new(Tags) }

// String returns label of tags
func (t Tags) String() string {
	_ = "STUB: not implemented"

	// AppID returns buildinTagApp of tags
	return ""
}

func (t Tags) AppID() string { _ = "STUB: not implemented"; return "" }

// Version returns buildinTagVersion of tags
func (t Tags) Version() string { _ = "STUB: not implemented"; return "" }

// IsSubsetOf returns if tags is labels
func (t Tags) IsSubsetOf(labels map[string]string) bool { _ = "STUB: not implemented"; return false }

// TODO: remove buildinTag version

// LabelOfTags returns tags as string
func LabelOfTags(t map[string]string) (ret string) { _ = "STUB: not implemented"; return "" }
