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
		projectile := &Projectile{
			Type:     "laser",
			Position: sdlutils.Vector3ToFVector3Center(item.Mob.Position),
			Source:   item.Mob,
		}
		AddProjectileToMap(GameInstance.Map, projectile)
		item.Mob.Fatigue += 60

		force := NewForce(projectile.Position, sdlutils.PointToFPointCenter(tile.Position.Base), item.Mob.Id, projectile)
		GameInstance.ForceController.Projectiles[projectile] = force

		return true
	}

	return template
}
