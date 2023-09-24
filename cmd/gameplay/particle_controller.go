package gameplay

type KurinParticleController struct {
	Pending []*KurinParticle
}

func NewKurinParticleController() KurinParticleController {
	return KurinParticleController{
		Pending: []*KurinParticle{},
	}
}

func CreateKurinParticle(controller *KurinParticleController, particle *KurinParticle) {
	controller.Pending = append(controller.Pending, particle)
}
