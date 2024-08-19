package music

import (
	"io/fs"
	"path"
	"path/filepath"
	"strings"

	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/LamkasDev/kurin/cmd/sound"
	"github.com/veandco/go-sdl2/mix"
)

type SoundLayerMusicData struct {
	Tracks map[string]*sound.Track
}

func NewSoundLayerMusic() *sound.SoundLayer {
	return &sound.SoundLayer{
		Load:    LoadSoundLayerMusic,
		Process: ProcessSoundLayerMusic,
		Data: SoundLayerMusicData{
			Tracks: map[string]*sound.Track{},
		},
	}
}

func LoadSoundLayerMusic(layer *sound.SoundLayer) error {
	mix.ReserveChannels(0)

	return filepath.WalkDir(path.Join(constants.SoundsPath, "music"), func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		name := strings.TrimSuffix(d.Name(), filepath.Ext(d.Name()))
		if layer.Data.(SoundLayerMusicData).Tracks[name], err = sound.NewTrack("music", name); err != nil {
			return err
		}

		return nil
	})
}

func ProcessSoundLayerMusic(layer *sound.SoundLayer) error {
	if mix.Playing(0) != 1 {
		track := layer.Data.(SoundLayerMusicData).Tracks["ambiicetheme"]
		c, err := track.Data.Play(0, 0)
		if err != nil {
			return err
		}
		mix.Volume(c, int(0.5*128))
	}

	return nil
}
