package templates

type KurinItemTemplate struct {
	Id   string  `json:"id"`
	Path *string `json:"path"`
	Hand *bool   `json:"hand"`
}
