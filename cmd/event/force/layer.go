package force

import (
	"math/rand/v2"

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
	for _, force := range gameplay.GameInstance.ForceController.Items {
		item := force.Data.(*gameplay.Item)
		result, _ := gameplay.AdvanceForce(force)
		switch result {
		case gameplay.ForceResultOutOfBounds:
			gameplay.RemoveItemFromMapRaw(gameplay.GameInstance.Map, item)
			delete(gameplay.GameInstance.ForceController.Items, item)
		case gameplay.ForceResultCollided:
			gameplay.PlaySound(&gameplay.GameInstance.SoundController, "grillehit")
			gameplay.CreateParticle(&gameplay.GameInstance.ParticleController, gameplay.NewParticleCross(item.Transform.Position, 0.75, sdl.Color{R: 210, G: 210, B: 210}))
			item.Transform.Rotation = rand.Float64() * 360
			delete(gameplay.GameInstance.ForceController.Items, item)
		case gameplay.ForceResultReached:
			item.Transform.Position.Base = force.Position.Base
			delete(gameplay.GameInstance.ForceController.Items, item)
		case gameplay.ForceResultNone:
			item.Transform.Position.Base = force.Position.Base
		}
	}
	for _, force := range gameplay.GameInstance.ForceController.Projectiles {
		projectile := force.Data.(*gameplay.Projectile)
		result, rawCollider := gameplay.AdvanceForce(force)
		switch result {
		case gameplay.ForceResultOutOfBounds:
		case gameplay.ForceResultReached:
			gameplay.RemoveProjectileFromMap(gameplay.GameInstance.Map, projectile)
			delete(gameplay.GameInstance.ForceController.Projectiles, projectile)
		case gameplay.ForceResultCollided:
			switch collider := rawCollider.(type) {
			case *gameplay.Object:
				gameplay.HitObject(collider)
			case *gameplay.Mob:
				gameplay.HitMob(collider)
				collider.Health.LastDamageSource = projectile.Source
			}
			gameplay.RemoveProjectileFromMap(gameplay.GameInstance.Map, projectile)
			delete(gameplay.GameInstance.ForceController.Projectiles, projectile)
		case gameplay.ForceResultNone:
			projectile.Position.Base = force.Position.Base
		}
	}

	return nil
}
