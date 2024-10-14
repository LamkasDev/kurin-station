package gameplay

type ForceController struct {
	Items       map[*Item]*Force
	Projectiles map[*Projectile]*Force
}

func NewForceController() ForceController {
	return ForceController{
		Items:       map[*Item]*Force{},
		Projectiles: map[*Projectile]*Force{},
	}
}
