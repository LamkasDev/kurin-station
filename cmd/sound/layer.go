package sound

import "github.com/LamkasDev/kurin/cmd/gameplay"

type KurinSoundLayer struct {
	Load    KurinSoundLayerLoad
	Process KurinSoundLayerProcess
	Data    interface{}
}

type KurinSoundLayerLoad func(manager *KurinSoundManager, layer *KurinSoundLayer) *error

type KurinSoundLayerProcess func(manager *KurinSoundManager, layer *KurinSoundLayer, game *gameplay.KurinGame) *error
