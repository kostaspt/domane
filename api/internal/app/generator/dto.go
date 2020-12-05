package generator

type Kind string

const (
	DotComExtension           Kind = "dot_com_extension"
	CommonExtension           Kind = "common_extension"
	CommonExtensionHyphenated Kind = "common_extension_hyphenated"
	ShortExtension            Kind = "short_extensions"
)

type Result struct {
	Domain   string `json:"domain"`
	Kind     Kind   `json:"kind"`
	Position uint   `json:"position"`
}

type Results []Result

func (r Results) Len() int { return len(r) }
func (r Results) Less(i, j int) bool {
	return r[j].Kind != DotComExtension && (r[i].Kind == DotComExtension || len(r[i].Domain) < len(r[j].Domain))
}
func (r Results) Swap(i, j int) { r[i], r[j] = r[j], r[i] }
