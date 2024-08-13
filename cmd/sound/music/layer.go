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

type KurinSoundLayerMusicData struct {
	Tracks map[string]*sound.KurinTrack
}

func NewKurinSoundLayerMusic() *sound.KurinSoundLayer {
	return &sound.KurinSoundLayer{
		Load:    LoadKurinSoundLayerMusic,
		Process: ProcessKurinSoundLayerMusic,
		Data: KurinSoundLayerMusicData{
			Tracks: map[string]*sound.KurinTrack{},
		},
	}
}

func LoadKurinSoundLayerMusic(manager *sound.KurinSoundManager, layer *sound.KurinSoundLayer) error {
	mix.ReserveChannels(0)

	return filepath.WalkDir(path.Join(constants.SoundsPath, "music"), func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		name := strings.TrimSuffix(d.Name(), filepath.Ext(d.Name()))
		if layer.Data.(KurinSoundLayerMusicData).Tracks[name], err = sound.NewKurinTrack(manager, "music", name); err != nil {
			return err
		}

		return nil
	})
}

func ProcessKurinSoundLayerMusic(manager *sound.KurinSoundManager, layer *sound.KurinSoundLayer) error {
	if mix.Playing(0) != 1 {
		track := layer.Data.(KurinSoundLayerMusicData).Tracks["ambiicetheme"]
		c, err := track.Data.Play(0, 0)
		if err != nil {
			return err
		}
		mix.Volume(c, int(0.5 * 128))
	}

	return nil
}
