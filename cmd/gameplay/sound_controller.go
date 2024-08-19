package gameplay

type SoundController struct {
	Pending []*Sound
}

func NewSoundController() SoundController {
	return SoundController{
		Pending: []*Sound{},
	}
}

func PlaySound(controller *SoundController, soundType string) {
	controller.Pending = append(controller.Pending, NewSound(soundType, 1))
}

func PlaySoundVolume(controller *SoundController, soundType string, volume float32) {
	controller.Pending = append(controller.Pending, NewSound(soundType, volume))
}
