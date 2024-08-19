package force

import (
	"math/rand/v2"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/event"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/veandco/go-sdl2/sdl"
)

type EventLayerForceData struct{}

func NewEventLayerForce() *event.EventLayer {
	return &event.EventLayer{
		Load:    LoadEventLayerForce,
		Process: ProcessEventLayerForce,
		Data:    &EventLayerForceData{},
	}
}

func LoadEventLayerForce(layer *event.EventLayer) error {
	return nil
}

func ProcessEventLayerForce(layer *event.EventLayer) error {
	for _, force := range gameplay.GameInstance.ForceController.Forces {
		if force.Item == nil {
			continue
		}
		newPosition := sdlutils.AddFPoints(force.Item.Transform.Position.Base, force.Delta)
		newVector := sdlutils.Vector3{Base: sdlutils.FPointToPointFloored(newPosition), Z: force.Item.Transform.Position.Z}
		if gameplay.IsMapPositionOutOfBounds(&gameplay.GameInstance.Map, newVector) {
			gameplay.RemoveItemFromMapRaw(&gameplay.GameInstance.Map, force.Item)
			delete(gameplay.GameInstance.ForceController.Forces, force.Item)
			continue
		}
		if gameplay.GetTileAt(&gameplay.GameInstance.Map, newVector) != nil && gameplay.CanEnterMapPosition(&gameplay.GameInstance.Map, newVector) != gameplay.EnteranceStatusYes {
			gameplay.PlaySound(&gameplay.GameInstance.SoundController, "grillehit")
			gameplay.CreateParticle(&gameplay.GameInstance.ParticleController, gameplay.NewParticleCross(force.Item.Transform.Position, 0.75, sdl.Color{R: 210, G: 210, B: 210}))
			force.Item.Transform.Rotation = rand.Float64() * 360
			delete(gameplay.GameInstance.ForceController.Forces, force.Item)
			continue
		}
		force.Item.Transform.Position.Base = newPosition
		if sdlutils.GetDistanceF(force.Item.Transform.Position.Base, force.Target) < 0.01 {
			delete(gameplay.GameInstance.ForceController.Forces, force.Item)
			continue
		}
	}

	return nil
}
