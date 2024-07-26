package force

import (
	"math/rand/v2"

	"github.com/LamkasDev/kurin/cmd/common/mathutils"
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/event"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinEventLayerForceData struct {
}

func NewKurinEventLayerForce() *event.KurinEventLayer {
	return &event.KurinEventLayer{
		Load:    LoadKurinEventLayerForce,
		Process: ProcessKurinEventLayerForce,
		Data:    KurinEventLayerForceData{},
	}
}

func LoadKurinEventLayerForce(manager *event.KurinEventManager, layer *event.KurinEventLayer) *error {
	return nil
}

func ProcessKurinEventLayerForce(manager *event.KurinEventManager, layer *event.KurinEventLayer, game *gameplay.KurinGame) *error {
	for _, force := range game.ForceController.Forces {
		if force.Item != nil {
			base := mathutils.LerpFPoint(force.Item.Transform.Position.Base, force.Target, 0.2)
			if !gameplay.CanEnterPosition(&game.Map, sdlutils.Vector3{Base: sdlutils.FPointToPointFloored(base), Z: force.Item.Transform.Position.Z}) {
				gameplay.PlaySound(&game.SoundController, "grillehit")
				gameplay.CreateKurinParticle(&game.ParticleController, gameplay.NewKurinParticleCross(game, sdlutils.FVector3{Base: sdl.FPoint{X: force.Item.Transform.Position.Base.X, Y: force.Item.Transform.Position.Base.Y}, Z: force.Item.Transform.Position.Z}))
				force.Item.Transform.Rotation = rand.Float64() * 360
				delete(game.ForceController.Forces, force.Item)
				continue
			}
			force.Item.Transform.Position.Base = base
			if sdlutils.GetDistanceF(force.Item.Transform.Position.Base, force.Target) < 0.01 {
				delete(game.ForceController.Forces, force.Item)
				continue
			}
		}
	}

	return nil
}
