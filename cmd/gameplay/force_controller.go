package gameplay

type ForceController struct {
	Items   map[*Item]*Force
	Bullets []*Force
}

func NewForceController() ForceController {
	return ForceController{
		Items:   map[*Item]*Force{},
		Bullets: []*Force{},
	}
}
