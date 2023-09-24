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
)

var KurinAvailablePitches = []float32{0.8, 0.9, 1, 1.1, 1.2, 1.3, 1.4}

type KurinSoundLayerVoiceData struct {
	Filesystem vfs.FileSystem
	Silence    *sound.KurinTrackComplex
	Normal     map[float32]*sound.KurinTrackComplex
}

func NewKurinSoundLayerVoice() *sound.KurinSoundLayer {
	return &sound.KurinSoundLayer{
		Load:    LoadKurinSoundLayerVoice,
		Process: ProcessKurinSoundLayerVoice,
		Data: KurinSoundLayerVoiceData{
			Filesystem: memoryfs.New(),
			Normal:     map[float32]*sound.KurinTrackComplex{},
		},
	}
}

func GetCachePitchedPath(trackId string, pitch float32) string {
	return path.Join(constants.TempAudioPath, fmt.Sprintf("%s_%d.ogg", trackId, int(pitch*100)))
}

func LoadKurinSoundLayerVoice(manager *sound.KurinSoundManager, layer *sound.KurinSoundLayer) *error {
	data := layer.Data.(KurinSoundLayerVoiceData)
	var err *error
	if data.Silence, err = sound.NewKurinTrackComplex(manager, "silence", 1); err != nil {
		return err
	}

	for _, pitch := range KurinAvailablePitches {
		path := GetCachePitchedPath("owl", pitch)
		if _, err := os.Stat(path); err == nil {
			continue
		}
		if data.Normal[pitch], err = sound.NewKurinTrackComplex(manager, "owl", pitch); err != nil {
			return err
		}
		os.WriteFile(path, data.Normal[pitch].Buffer.Bytes(), 777)
	}

	layer.Data = data
	return nil
}

func ProcessKurinSoundLayerVoice(manager *sound.KurinSoundManager, layer *sound.KurinSoundLayer, game *gameplay.KurinGame) *error {
	data := layer.Data.(KurinSoundLayerVoiceData)
	if len(game.RunechatController.Sounds) > 0 {
		for _, runechatSound := range game.RunechatController.Sounds {
			paths := []string{}
			for _, rune := range runechatSound.Runechat.Message {
				switch rune {
				case ' ', ',', '.':
					paths = append(paths, data.Silence.Base.Path)
				default:
					pitch := KurinAvailablePitches[rand.Intn(len(KurinAvailablePitches))]
					paths = append(paths, GetCachePitchedPath("owl", pitch))
				}
			}

			track, err := sound.ConcatKurinTrackComplex(manager, paths)
			if err != nil {
				return err
			}
			if err := track.Base.Data.Play(1); err != nil {
				return &err
			}
		}
		game.RunechatController.Sounds = []*gameplay.KurinRunechatSound{}
	}

	return nil
}
