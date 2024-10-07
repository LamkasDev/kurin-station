package gameplay

import "github.com/LamkasDev/kurin/cmd/common/sdlutils"

func CanMobInteractWithTile(mob *Mob, tile *Tile) bool {
	return sdlutils.GetDistanceSimple(mob.Position.Base, tile.Position.Base) <= 1
}

func CanMobInteractWithItem(mob *Mob, item *Item) bool {
	return sdlutils.GetDistanceSimple(mob.Position.Base, sdlutils.FPointToPoint(item.Transform.Position.Base)) <= 1
}

func CanMobInteractWithMob(mob *Mob, other *Mob) bool {
	return sdlutils.GetDistanceSimple(mob.Position.Base, other.Position.Base) <= 1
}

func MobHitObject(mob *Mob, object *Object) {
	PlayMobAnimation(mob, "hit")
	HitObject(object)
	mob.Fatigue += 60
}
