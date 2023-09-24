package gameplay

type KurinSoundController struct {
	Pending []*KurinSound
}

func NewKurinSoundController() KurinSoundController {
	return KurinSoundController{
		Pending: []*KurinSound{},
	}
}

func PlaySound(controller *KurinSoundController, stype string) {
	controller.Pending = append(controller.Pending, NewKurinSound(stype))
}
