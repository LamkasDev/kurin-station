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

func LoadKurinEventLayerForce(manager *event.KurinEventManager, layer *event.KurinEventLayer) error {
	return nil
}

func ProcessKurinEventLayerForce(manager *event.KurinEventManager, layer *event.KurinEventLayer) error {
	for _, force := range gameplay.KurinGameInstance.ForceController.Forces {
		if force.Item == nil {
			continue
		}
		base := mathutils.LerpFPoint(force.Item.Transform.Position.Base, force.Target, 0.2)
		if !gameplay.CanEnterPosition(&gameplay.KurinGameInstance.Map, sdlutils.Vector3{Base: sdlutils.FPointToPointFloored(base), Z: force.Item.Transform.Position.Z}) {
			gameplay.PlaySound(&gameplay.KurinGameInstance.SoundController, "grillehit")
			gameplay.CreateKurinParticle(&gameplay.KurinGameInstance.ParticleController, gameplay.NewKurinParticleCross(force.Item.Transform.Position, 0.75, sdl.Color{R: 210, G: 210, B: 210}))
			force.Item.Transform.Rotation = rand.Float64() * 360
			delete(gameplay.KurinGameInstance.ForceController.Forces, force.Item)
			continue
		}
		force.Item.Transform.Position.Base = base
		if sdlutils.GetDistanceF(force.Item.Transform.Position.Base, force.Target) < 0.01 {
			delete(gameplay.KurinGameInstance.ForceController.Forces, force.Item)
			continue
		}
	}

	return nil
}
