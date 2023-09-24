package ambient

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/sound"
)

type KurinSoundLayerAmbientData struct {
	Tracks map[string]*sound.KurinTrack
}

func NewKurinSoundLayerAmbient() *sound.KurinSoundLayer {
	return &sound.KurinSoundLayer{
		Load:    LoadKurinSoundLayerAmbient,
		Process: ProcessKurinSoundLayerAmbient,
		Data: KurinSoundLayerAmbientData{
			Tracks: map[string]*sound.KurinTrack{},
		},
	}
}

func LoadKurinSoundLayerAmbient(manager *sound.KurinSoundManager, layer *sound.KurinSoundLayer) *error {
	var err *error
	if layer.Data.(KurinSoundLayerAmbientData).Tracks["grillehit"], err = sound.NewKurinTrack(manager, "grillehit"); err != nil {
		return err
	}

	return nil
}

func ProcessKurinSoundLayerAmbient(manager *sound.KurinSoundManager, layer *sound.KurinSoundLayer, game *gameplay.KurinGame) *error {
	if len(game.SoundController.Pending) > 0 {
		for _, sound := range game.SoundController.Pending {
			track := layer.Data.(KurinSoundLayerAmbientData).Tracks[sound.Type]
			if err := track.Data.Play(1); err != nil {
				return &err
			}
		}
		game.SoundController.Pending = []*gameplay.KurinSound{}
	}

	return nil
}
