package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

func CanKurinCharacterInteractWithTile(character *KurinCharacter, tile *KurinTile) bool {
	return sdlutils.GetDistanceSimple(character.Position.Base, tile.Position.Base) <= 1
}

func CanKurinCharacterInteractWithItem(character *KurinCharacter, item *KurinItem) bool {
	return sdlutils.GetDistanceSimple(character.Position.Base, sdlutils.FPointToPoint(item.Transform.Position.Base)) <= 1
}

func CanKurinCharacterInteractWithCharacter(character *KurinCharacter, other *KurinCharacter) bool {
	return sdlutils.GetDistanceSimple(character.Position.Base, other.Position.Base) <= 1
}

func InteractKurinCharacter(character *KurinCharacter, position sdl.Point) {
	if !character.Moving {
		TurnKurinCharacterTo(character, position)
	}
	if character.Fatigue > 0 {
		return
	}

	tile := GetKurinTileAt(&GameInstance.Map, sdlutils.Vector3{Base: position, Z: character.Position.Z})
	if tile == nil || !CanKurinCharacterInteractWithTile(character, tile) {
		return
	}

	object := GetKurinObjectAtTile(tile)
	item := character.Inventory.Hands[character.ActiveHand]
	if object != nil {
		hit := true
		if item != nil {
			hit = item.CanHit
		}
		if object.OnInteraction(object, item) {
			hit = false
		}
		if hit {
			KurinCharacterHitObject(character, object)
		}
	} else if item != nil {
		item.OnTileInteraction(item, tile)
	}

	if GameInstance.HoveredItem != nil {
		if TransferKurinItemToCharacter(GameInstance.HoveredItem, character) {
			character.Fatigue += 20
		}
	}
}

func KurinCharacterHitObject(character *KurinCharacter, object *KurinObject) {
	PlayKurinCharacterAnimation(character, "hit")
	HitKurinObject(object)
	character.Fatigue += 60
}
