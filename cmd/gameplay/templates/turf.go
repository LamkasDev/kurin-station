package templates

type KurinTurfTemplate struct {
	Id           string                          `json:"id"`
	Name         string                          `json:"name"`
	Requirements *[]KurinItemRequirementTemplate `json:"requirements"`
	Path         *string                         `json:"path"`
	Rotate       *bool                           `json:"rotate"`
}
