package templates

type KurinStructureTemplate struct {
	Id          string  `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Rotate      *bool   `json:"rotate"`
}
