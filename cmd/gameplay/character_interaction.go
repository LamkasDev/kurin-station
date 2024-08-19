package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

func CanCharacterInteractWithTile(character *Character, tile *Tile) bool {
	return sdlutils.GetDistanceSimple(character.Position.Base, tile.Position.Base) <= 1
}

func CanCharacterInteractWithItem(character *Character, item *Item) bool {
	return sdlutils.GetDistanceSimple(character.Position.Base, sdlutils.FPointToPoint(item.Transform.Position.Base)) <= 1
}

func CanCharacterInteractWithCharacter(character *Character, other *Character) bool {
	return sdlutils.GetDistanceSimple(character.Position.Base, other.Position.Base) <= 1
}

func InteractCharacter(character *Character, position sdl.Point) {
	if character.MovementTicks == 0 {
		TurnCharacterTo(character, position)
	}
	if character.Fatigue > 0 {
		return
	}

	tile := GetTileAt(&GameInstance.Map, sdlutils.Vector3{Base: position, Z: character.Position.Z})
	if tile == nil || !CanCharacterInteractWithTile(character, tile) {
		return
	}
	object := GetObjectAtTile(tile)
	item := character.Inventory.Hands[character.ActiveHand]
	if object != nil {
		hit := true
		if item != nil {
			hit = item.Template.CanHit
		}
		if object.Template.OnInteraction(object, item) {
			hit = false
		}
		if hit {
			CharacterHitObject(character, object)
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

func CharacterHitObject(character *Character, object *Object) {
	PlayCharacterAnimation(character, "hit")
	HitObject(object)
	character.Fatigue += 60
}
