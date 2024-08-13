package sound

import "C"
import (
	"bytes"
	"fmt"
	"path"

	"github.com/LamkasDev/kurin/cmd/common/constants"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinTrackComplex struct {
	Base        *KurinTrack
	Buffer      *bytes.Buffer
	Stream      *ffmpeg_go.Stream
	FinalStream *sdl.RWops
}

func NewKurinTrackComplex(manager *KurinSoundManager, trackId string, pitch float32) (*KurinTrackComplex, error) {
	track := KurinTrackComplex{
		Base: &KurinTrack{
			Path: path.Join(constants.SoundsPath, "effects", fmt.Sprintf("%s.ogg", trackId)),
		},
	}

	track.Buffer = bytes.NewBuffer(nil)
	track.Stream = ffmpeg_go.Input(track.Base.Path).Output("pipe:", ffmpeg_go.KwArgs{
		"format": "ogg",
		"af":     fmt.Sprintf("asetrate=44100*%f,aresample=44100,atempo=1/%f", pitch, pitch),
	}).WithOutput(track.Buffer)
	err := track.Stream.Run()
	if err != nil {
		return &track, err
	}
	track.FinalStream, err = sdl.RWFromMem(track.Buffer.Bytes())
	if err != nil {
		return &track, err
	}
	track.Base.Data, err = mix.LoadWAVRW(track.FinalStream, false)
	if err != nil {
		return &track, err
	}

	return &track, nil
}

func ConcatKurinTrackComplex(manager *KurinSoundManager, paths []string) (*KurinTrackComplex, error) {
	track := KurinTrackComplex{
		Base: &KurinTrack{},
	}

	streams := []*ffmpeg_go.Stream{}
	for _, path := range paths {
		streams = append(streams, ffmpeg_go.Input(path))
	}

	track.Buffer = bytes.NewBuffer(nil)
	track.Stream = ffmpeg_go.Concat(streams, ffmpeg_go.KwArgs{
		"v": 0,
		"a": 1,
	}).Output("pipe:", ffmpeg_go.KwArgs{
		"format": "ogg",
	}).WithOutput(track.Buffer)
	err := track.Stream.Run()
	if err != nil {
		return &track, err
	}
	track.FinalStream, err = sdl.RWFromMem(track.Buffer.Bytes())
	if err != nil {
		return &track, err
	}
	track.Base.Data, err = mix.LoadWAVRW(track.FinalStream, false)
	if err != nil {
		return &track, err
	}

	return &track, nil
}
