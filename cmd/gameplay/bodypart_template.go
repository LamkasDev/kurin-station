package gameplay

type BodypartTemplate struct {
	Type      string
	MaxPoints uint16
}

func NewBodypartTemplate[D any](bodypartType string, maxPoints uint16) *BodypartTemplate {
	return &BodypartTemplate{
		Type:      bodypartType,
		MaxPoints: maxPoints,
	}
}
