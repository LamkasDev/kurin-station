package gameplay

type KurinSoundController struct {
	Pending []*KurinSound
}

func NewKurinSoundController() KurinSoundController {
	return KurinSoundController{
		Pending: []*KurinSound{},
	}
}

func PlaySound(controller *KurinSoundController, soundType string) {
	controller.Pending = append(controller.Pending, NewKurinSound(soundType, 1))
}

func PlaySoundVolume(controller *KurinSoundController, soundType string, volume float32) {
	controller.Pending = append(controller.Pending, NewKurinSound(soundType, volume))
}
