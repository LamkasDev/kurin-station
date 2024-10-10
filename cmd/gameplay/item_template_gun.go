package gameplay

import "github.com/LamkasDev/kurin/cmd/common/sdlutils"

type ItemGunData struct{}

func NewItemTemplateGun() *ItemTemplate {
	template := NewItemTemplate[*ItemWelderData]("gun", 1, 5)
	template.CanHit = false
	template.GetDefaultData = func() interface{} {
		return &ItemGunData{}
	}
	template.OnTileInteraction = func(item *Item, tile *Tile) bool {
		PlaySound(&GameInstance.SoundController, "laser_gun")
		force := NewForce(sdlutils.Vector3ToFVector3Center(item.Mob.Position), sdlutils.PointToFPointCenter(tile.Position.Base), item.Mob.Id, nil)
		result, rawCollider := RushForce(force)
		switch result {
		case ForceResultCollided:
			switch collider := rawCollider.(type) {
			case *Object:
				HitObject(collider)
			case *Mob:
				HitMob(collider)
			}
		}
		item.Mob.Fatigue += 60

		return true
	}

	return template
}
