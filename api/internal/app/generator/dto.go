package generator

type Kind string

const (
	ExtensionDotCom           Kind = "extension_dot_com"
	ExtensionCommon           Kind = "extension_common"
	ExtensionCommonHyphenated Kind = "extension_common_hyphenated"
	ExtensionShort            Kind = "extensions_short"
	SimilarSynonym            Kind = "similar_synonym"
)

type Result struct {
	Domain   string `json:"domain"`
	Kind     Kind   `json:"kind"`
	Position uint   `json:"position"`
}

type Results []Result

func (r Results) Len() int { return len(r) }
func (r Results) Less(i, j int) bool {
	return r[j].Kind != ExtensionDotCom && (r[i].Kind == ExtensionDotCom || len(r[i].Domain) < len(r[j].Domain))
}
func (r Results) Swap(i, j int) { r[i], r[j] = r[j], r[i] }