package templates

type KurinItemTemplate struct {
	Id     string `json:"id"`
	Hand   *bool  `json:"hand"`
	States *int   `json:"states"`
}
