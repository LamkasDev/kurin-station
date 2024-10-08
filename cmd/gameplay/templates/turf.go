package templates

type TurfTemplate struct {
	Id          uint8   `json:"id"`
	Name        string  `json:"name"`
	Path        *string `json:"path"`
	Rotate      *bool   `json:"rotate"`
	Transparent bool    `json:"transparent"`
	States      *int    `json:"states"`
}
