package templates

type SpeciesTemplate struct {
	Id    string                    `json:"id"`
	Parts []SpeciesTemplateBodypart `json:"parts"`
}

type SpeciesTemplateBodypart struct {
	Id     string                         `json:"id"`
	Type   *bool                          `json:"type"`
	Path   *string                        `json:"path"`
	Offset *SpeciesTemplateBodypartOffset `json:"offset"`
}

type SpeciesTemplateBodypartOffset struct {
	X int32 `json:"x"`
	Y int32 `json:"y"`
}
