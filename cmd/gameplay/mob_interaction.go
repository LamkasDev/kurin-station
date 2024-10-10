package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

func CanMobInteractWithTile(mob *Mob, tile *Tile) bool {
	return sdlutils.GetDistanceSimple(mob.Position.Base, tile.Position.Base) <= 1
}

func CanMobInteractWithItem(mob *Mob, item *Item) bool {
	return sdlutils.GetDistanceSimple(mob.Position.Base, sdlutils.FPointToPoint(item.Transform.Position.Base)) <= 1
}

func CanMobInteractWithMob(mob *Mob, other *Mob) bool {
	return sdlutils.GetDistanceSimple(mob.Position.Base, other.Position.Base) <= 1
}

func HitMob(mob *Mob) {
	if mob.Health.Dead {
		return
	}
	PlaySound(&GameInstance.SoundController, "grillehit")
	particle := NewParticleCross(
		sdlutils.Vector3ToFVector3Center(mob.Position),
		0.75,
		sdl.Color{R: 210, G: 40, B: 40},
	)
	CreateParticle(&GameInstance.ParticleController, particle)
	mob.Health.Points--
	if mob.Health.Points <= 0 {
		KillMob(mob)
		mob.Health.Dead = true
	}
}

func MobHitObject(mob *Mob, target *Object) {
	PlayMobAnimation(mob, "hit")
	HitObject(target)
	mob.Fatigue += 60
}

func MobHitMob(mob *Mob, target *Mob) {
	PlayMobAnimation(mob, "hit")
	HitMob(target)
	mob.Fatigue += 60
}
