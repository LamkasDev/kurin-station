package templates

type KurinStructureTemplate struct {
	Id          string  `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Path        *string `json:"path"`
	Rotate      *bool   `json:"rotate"`
}
