package voice

import (
	"fmt"
	"math/rand"
	"os"
	"path"

	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/sound"
	"github.com/mandelsoft/vfs/pkg/memoryfs"
	"github.com/mandelsoft/vfs/pkg/vfs"
	"github.com/veandco/go-sdl2/mix"
)

var AvailablePitches = []float32{0.8, 0.9, 1, 1.1, 1.2, 1.3, 1.4}

type SoundLayerVoiceData struct {
	Filesystem vfs.FileSystem
	Silence    *sound.TrackComplex
	Normal     map[float32]*sound.TrackComplex
}

func NewSoundLayerVoice() *sound.SoundLayer {
	return &sound.SoundLayer{
		Load:    LoadSoundLayerVoice,
		Process: ProcessSoundLayerVoice,
		Data: SoundLayerVoiceData{
			Filesystem: memoryfs.New(),
			Normal:     map[float32]*sound.TrackComplex{},
		},
	}
}

func GetCachePitchedPath(trackId string, pitch float32) string {
	return path.Join(constants.TempAudioPath, fmt.Sprintf("%s_%d.ogg", trackId, int(pitch*100)))
}

func LoadSoundLayerVoice(layer *sound.SoundLayer) error {
	data := layer.Data.(SoundLayerVoiceData)
	var err error
	if data.Silence, err = sound.NewTrackComplex("silence", 1); err != nil {
		return err
	}

	for _, pitch := range AvailablePitches {
		path := GetCachePitchedPath("owl", pitch)
		if _, err := os.Stat(path); err == nil {
			continue
		}
		if data.Normal[pitch], err = sound.NewTrackComplex("owl", pitch); err != nil {
			return err
		}
		os.WriteFile(path, data.Normal[pitch].Buffer.Bytes(), 777)
	}

	layer.Data = data
	return nil
}

func ProcessSoundLayerVoice(layer *sound.SoundLayer) error {
	data := layer.Data.(SoundLayerVoiceData)
	if len(gameplay.GameInstance.RunechatController.Sounds) > 0 {
		for _, runechatSound := range gameplay.GameInstance.RunechatController.Sounds {
			paths := []string{}
			for _, rune := range runechatSound.Runechat.Message {
				switch rune {
				case ' ', ',', '.':
					paths = append(paths, data.Silence.Base.Path)
				default:
					pitch := AvailablePitches[rand.Intn(len(AvailablePitches))]
					paths = append(paths, GetCachePitchedPath("owl", pitch))
				}
			}

			track, err := sound.ConcatTrackComplex(paths)
			if err != nil {
				return err
			}
			c, err := track.Base.Data.Play(-1, 0)
			if err != nil {
				return err
			}
			mix.Volume(c, 128)
		}
		gameplay.GameInstance.RunechatController.Sounds = []*gameplay.RunechatSound{}
	}

	return nil
}
