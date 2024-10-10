package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

func InteractObject(mob *Mob, object *Object) {
	MobHitObject(mob, object)
}

func HitObject(object *Object) {
	PlaySound(&GameInstance.SoundController, "grillehit")
	particle := NewParticleCross(
		sdlutils.Vector3ToFVector3Center(object.Tile.Position),
		0.75,
		sdl.Color{R: 210, G: 210, B: 210},
	)
	CreateParticle(&GameInstance.ParticleController, particle)
	object.Health--
	if object.Health <= 0 {
		DestroyObject(object)
	}
}
