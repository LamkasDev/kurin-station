package templates

type ItemTemplate struct {
	Id         string `json:"id"`
	Hand       *bool  `json:"hand"`
	States     *int   `json:"states"`
	StatesHand *int   `json:"statesHand"`
}
