package gameplay

import (
	"github.com/veandco/go-sdl2/sdl"
)

func InteractCharacter(character *Mob, position sdl.Point) {
	if character.MovementTicks == 0 {
		TurnMobTo(character, position)
	}
	if character.Fatigue > 0 || character.Health.Dead {
		return
	}
	if GameInstance.HoveredTile == nil {
		return
	}

	if GameInstance.HoveredMob != nil && CanMobInteractWithMob(character, GameInstance.HoveredMob) {
		MobHitMob(character, GameInstance.HoveredMob)
		return
	}

	if GameInstance.HoveredItem != nil && CanMobInteractWithItem(character, GameInstance.HoveredItem) && GameInstance.HoveredItem.Template.CanPickup {
		if TransferItemToCharacter(GameInstance.HoveredItem, character) {
			character.Fatigue += 20
		}
		return
	}

	item := GetHeldItem(character)
	if GameInstance.HoveredObject != nil && CanMobInteractWithTile(character, GameInstance.HoveredTile) {
		hit := true
		if item != nil {
			hit = item.Template.CanHit
		}
		if GameInstance.HoveredObject.Template.OnInteraction(GameInstance.HoveredObject, item) {
			hit = false
		}
		if hit {
			MobHitObject(character, GameInstance.HoveredObject)
		}
	} else if item != nil {
		item.Template.OnTileInteraction(item, GameInstance.HoveredTile)
	}
}
