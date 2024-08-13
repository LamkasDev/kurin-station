package event

type KurinEventLayer struct {
	Load    KurinEventLayerLoad
	Process KurinEventLayerProcess
	Data    interface{}
}

type KurinEventLayerLoad func(manager *KurinEventManager, layer *KurinEventLayer) error

type KurinEventLayerProcess func(manager *KurinEventManager, layer *KurinEventLayer) error
