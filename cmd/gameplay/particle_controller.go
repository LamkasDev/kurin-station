package gameplay

type ParticleController struct {
	Pending []*Particle
}

func NewParticleController() ParticleController {
	return ParticleController{
		Pending: []*Particle{},
	}
}

func CreateParticle(controller *ParticleController, particle *Particle) {
	controller.Pending = append(controller.Pending, particle)
}
