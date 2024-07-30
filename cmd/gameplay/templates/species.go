package templates

type KurinSpeciesTemplate struct {
	Id    string                         `json:"id"`
	Parts []KurinSpeciesTemplateBodypart `json:"parts"`
}

type KurinSpeciesTemplateBodypart struct {
	Id     string                              `json:"id"`
	Type   *bool                               `json:"type"`
	Path   *string                             `json:"path"`
	Offset *KurinSpeciesTemplateBodypartOffset `json:"offset"`
}

type KurinSpeciesTemplateBodypartOffset struct {
	X int32 `json:"x"`
	Y int32 `json:"y"`
}
