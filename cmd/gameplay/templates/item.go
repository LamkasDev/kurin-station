package templates

type KurinItemTemplate struct {
	Id         string `json:"id"`
	Hand       *bool  `json:"hand"`
	States     *int   `json:"states"`
	StatesHand *int   `json:"statesHand"`
}

type KurinItemRequirementTemplate struct {
	Type  string `json:"type"`
	Count uint16 `json:"count"`
}
