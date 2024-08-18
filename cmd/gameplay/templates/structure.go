package templates

type KurinStructureTemplate struct {
	Id           string                          `json:"id"`
	Name         string                          `json:"name"`
	Description  *string                         `json:"description"`
	Requirements *[]KurinItemRequirementTemplate `json:"requirements"`
	Rotate       *bool                           `json:"rotate"`
	Smooth       *bool                           `json:"smooth"`
	States       *int                            `json:"states"`
}
