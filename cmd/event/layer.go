package event

import "github.com/LamkasDev/kurin/cmd/gameplay"

type KurinEventLayer struct {
	Load    KurinEventLayerLoad
	Process KurinEventLayerProcess
	Data    interface{}
}

type KurinEventLayerLoad func(manager *KurinEventManager, layer *KurinEventLayer) *error

type KurinEventLayerProcess func(manager *KurinEventManager, layer *KurinEventLayer, game *gameplay.KurinGame) *error
