package sound

type SoundLayer struct {
	Load    SoundLayerLoad
	Process SoundLayerProcess
	Data    interface{}
}

type SoundLayerLoad func(layer *SoundLayer) error

type SoundLayerProcess func(layer *SoundLayer) error
