package templates

type StructureTemplate struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Rotate      *bool   `json:"rotate"`
	Smooth      *bool   `json:"smooth"`
	States      *int    `json:"states"`
}
