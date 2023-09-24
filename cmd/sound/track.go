package sound

import "C"
import (
	"fmt"
	"path"

	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/veandco/go-sdl2/mix"
)

type KurinTrack struct {
	Path string
	Data *mix.Music
}

func NewKurinTrack(manager *KurinSoundManager, trackId string) (*KurinTrack, *error) {
	track := KurinTrack{
		Path: path.Join(constants.SoundsPath, "effects", fmt.Sprintf("%s.ogg", trackId)),
	}

	var err error
	track.Data, err = mix.LoadMUS(track.Path)
	if err != nil {
		return &track, nil
	}

	return &track, nil
}
