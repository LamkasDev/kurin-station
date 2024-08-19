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

type SoundLayerAmbientData struct {
	Tracks map[string]*sound.Track
}

func NewSoundLayerAmbient() *sound.SoundLayer {
	return &sound.SoundLayer{
		Load:    LoadSoundLayerAmbient,
		Process: ProcessSoundLayerAmbient,
		Data: SoundLayerAmbientData{
			Tracks: map[string]*sound.Track{},
		},
	}
}

func LoadSoundLayerAmbient(layer *sound.SoundLayer) error {
	return filepath.WalkDir(path.Join(constants.SoundsPath, "effects"), func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		name := strings.TrimSuffix(d.Name(), filepath.Ext(d.Name()))
		if layer.Data.(SoundLayerAmbientData).Tracks[name], err = sound.NewTrack("effects", name); err != nil {
			return err
		}

		return nil
	})
}

func ProcessSoundLayerAmbient(layer *sound.SoundLayer) error {
	if len(gameplay.GameInstance.SoundController.Pending) > 0 {
		for _, sound := range gameplay.GameInstance.SoundController.Pending {
			track := layer.Data.(SoundLayerAmbientData).Tracks[sound.Type]
			c, err := track.Data.Play(-1, 0)
			if err != nil {
				return err
			}
			mix.Volume(c, int(sound.Volume*128))
		}
		gameplay.GameInstance.SoundController.Pending = []*gameplay.Sound{}
	}

	return nil
}
