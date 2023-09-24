package sound

import "C"
import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinSoundManager struct {
	Layers []*KurinSoundLayer
	Device sdl.AudioDeviceID
}

func NewKurinSoundManager() (KurinSoundManager, *error) {
	manager := KurinSoundManager{
		Layers: []*KurinSoundLayer{},
	}

	if err := mix.Init(mix.INIT_OGG); err != nil {
		return manager, &err
	}
	if err := mix.OpenAudio(mix.DEFAULT_FREQUENCY, mix.DEFAULT_FORMAT, mix.DEFAULT_CHANNELS, mix.DEFAULT_CHUNKSIZE); err != nil {
		return manager, &err
	}

	return manager, nil
}

func LoadKurinSoundManager(manager *KurinSoundManager) *error {
	for _, layer := range manager.Layers {
		if err := layer.Load(manager, layer); err != nil {
			return err
		}
	}

	return nil
}

func ProcessKurinSoundManager(manager *KurinSoundManager, game *gameplay.KurinGame) *error {
	for _, layer := range manager.Layers {
		if err := layer.Process(manager, layer, game); err != nil {
			return err
		}
	}

	return nil
}

// TODO: make layers free themselves.
func FreeKurinSoundManager(manager *KurinSoundManager) *error {
	mix.Quit()
	mix.CloseAudio()

	return nil
}
