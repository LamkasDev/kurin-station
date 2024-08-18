package force

import (
	"math/rand/v2"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/event"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinEventLayerForceData struct{}

func NewKurinEventLayerForce() *event.EventLayer {
	return &event.EventLayer{
		Load:    LoadKurinEventLayerForce,
		Process: ProcessKurinEventLayerForce,
		Data:    &KurinEventLayerForceData{},
	}
}

func LoadKurinEventLayerForce(layer *event.EventLayer) error {
	return nil
}

func ProcessKurinEventLayerForce(layer *event.EventLayer) error {
	for _, force := range gameplay.GameInstance.ForceController.Forces {
		if force.Item == nil {
			continue
		}
		newPosition := sdlutils.AddFPoints(force.Item.Transform.Position.Base, force.Delta)
		newVector := sdlutils.Vector3{Base: sdlutils.FPointToPointFloored(newPosition), Z: force.Item.Transform.Position.Z}
		if gameplay.IsMapPositionOutOfBounds(&gameplay.GameInstance.Map, newVector) {
			gameplay.RemoveKurinItemFromMapRaw(&gameplay.GameInstance.Map, force.Item)
			delete(gameplay.GameInstance.ForceController.Forces, force.Item)
			continue
		}
		if gameplay.GetKurinTileAt(&gameplay.GameInstance.Map, newVector) != nil && !gameplay.CanEnterMapPosition(&gameplay.GameInstance.Map, newVector) {
			gameplay.PlaySound(&gameplay.GameInstance.SoundController, "grillehit")
			gameplay.CreateKurinParticle(&gameplay.GameInstance.ParticleController, gameplay.NewKurinParticleCross(force.Item.Transform.Position, 0.75, sdl.Color{R: 210, G: 210, B: 210}))
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
