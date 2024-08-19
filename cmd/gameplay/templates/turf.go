package templates

type TurfTemplate struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Path   *string `json:"path"`
	Rotate *bool   `json:"rotate"`
}
