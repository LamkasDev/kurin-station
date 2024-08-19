package sound

import "C"
import (
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

var SoundManagerInstance *SoundManager

type SoundManager struct {
	Layers []*SoundLayer
	Device sdl.AudioDeviceID
}

func InitializeSoundManager() error {
	SoundManagerInstance = &SoundManager{
		Layers: []*SoundLayer{},
	}

	if err := mix.Init(mix.INIT_OGG); err != nil {
		return err
	}
	if err := mix.OpenAudio(mix.DEFAULT_FREQUENCY, mix.DEFAULT_FORMAT, mix.DEFAULT_CHANNELS, mix.DEFAULT_CHUNKSIZE); err != nil {
		return err
	}
	mix.AllocateChannels(32)

	return nil
}

func LoadSoundManager() error {
	for _, layer := range SoundManagerInstance.Layers {
		if err := layer.Load(layer); err != nil {
			return err
		}
	}

	return nil
}

func ProcessSoundManager() error {
	for _, layer := range SoundManagerInstance.Layers {
		if err := layer.Process(layer); err != nil {
			return err
		}
	}

	return nil
}

// TODO: make layers free themselves.
func FreeSoundManager() error {
	mix.Quit()
	mix.CloseAudio()

	return nil
}
