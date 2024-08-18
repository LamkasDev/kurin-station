package ambient

import (
	"io/fs"
	"path"
	"path/filepath"
	"strings"

	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/sound"
	"github.com/veandco/go-sdl2/mix"
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

func LoadKurinSoundLayerAmbient(manager *sound.KurinSoundManager, layer *sound.KurinSoundLayer) error {
	return filepath.WalkDir(path.Join(constants.SoundsPath, "effects"), func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		name := strings.TrimSuffix(d.Name(), filepath.Ext(d.Name()))
		if layer.Data.(KurinSoundLayerAmbientData).Tracks[name], err = sound.NewKurinTrack(manager, "effects", name); err != nil {
			return err
		}

		return nil
	})
}

func ProcessKurinSoundLayerAmbient(manager *sound.KurinSoundManager, layer *sound.KurinSoundLayer) error {
	if len(gameplay.GameInstance.SoundController.Pending) > 0 {
		for _, sound := range gameplay.GameInstance.SoundController.Pending {
			track := layer.Data.(KurinSoundLayerAmbientData).Tracks[sound.Type]
			c, err := track.Data.Play(-1, 0)
			if err != nil {
				return err
			}
			mix.Volume(c, int(sound.Volume*128))
		}
		gameplay.GameInstance.SoundController.Pending = []*gameplay.KurinSound{}
	}

	return nil
}
