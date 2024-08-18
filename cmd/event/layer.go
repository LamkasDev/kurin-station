package event

type EventLayer struct {
	Load    EventLayerLoad
	Process EventLayerProcess
	Data    interface{}
}

type EventLayerLoad func(layer *EventLayer) error

type EventLayerProcess func(layer *EventLayer) error
