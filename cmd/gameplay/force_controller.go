package gameplay

type ForceController struct {
	Forces map[*Item]*Force
}

func NewForceController() ForceController {
	return ForceController{
		Forces: map[*Item]*Force{},
	}
}
