package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

func InteractCharacter(character *Mob, position sdl.Point) {
	if character.MovementTicks == 0 {
		TurnMobTo(character, position)
	}
	if character.Fatigue > 0 {
		return
	}

	tile := GetTileAt(&GameInstance.Map, sdlutils.Vector3{Base: position, Z: character.Position.Z})
	if tile == nil || !CanMobInteractWithTile(character, tile) {
		return
	}
	object := GetObjectAtTile(tile)
	item := GetHeldItem(character)
	if object != nil {
		hit := true
		if item != nil {
			hit = item.Template.CanHit
		}
		if object.Template.OnInteraction(object, item) {
			hit = false
		}
		if hit {
			MobHitObject(character, object)
		}
	} else if item != nil {
		item.Template.OnTileInteraction(item, tile)
	}

	if GameInstance.HoveredItem != nil {
		if TransferItemToCharacter(GameInstance.HoveredItem, character) {
			character.Fatigue += 20
		}
	}
}
