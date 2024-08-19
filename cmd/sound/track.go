package sound

import "C"
import (
	"fmt"
	"path"

	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/veandco/go-sdl2/mix"
)

type Track struct {
	Path string
	Data *mix.Chunk
}

func NewTrack(trackDirectory string, trackId string) (*Track, error) {
	track := Track{
		Path: path.Join(constants.SoundsPath, trackDirectory, fmt.Sprintf("%s.ogg", trackId)),
	}

	var err error
	track.Data, err = mix.LoadWAV(track.Path)
	if err != nil {
		return &track, nil
	}

	return &track, nil
}
