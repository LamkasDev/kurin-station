package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
	"github.com/veandco/go-sdl2/sdl"
)

func TeleportRandomlyCharacter(character *Character) {
	for {
		position := GetRandomMapPosition(&GameInstance.Map)
		if MoveCharacter(character, position) {
			character.PositionRender = sdlutils.PointToFPoint(position.Base)
			break
		}
	}
}

func MoveCharacterDirection(character *Character, direction common.Direction) bool {
	return MoveCharacter(character, common.GetPositionInDirectionV(character.Position, direction))
}

func MoveCharacter(character *Character, position sdlutils.Vector3) bool {
	if CanEnterMapPosition(&GameInstance.Map, position) == EnteranceStatusNo {
		return false
	}
	TurnCharacterTo(character, position.Base)
	tile := GetTileAt(&GameInstance.Map, position)
	object := GetObjectAtTile(tile)
	if object != nil {
		switch data := object.Data.(type) {
		case *ObjectAirlockData:
			if !data.Open {
				object.Template.OnInteraction(object, nil)
				return false
			}
		}
	}
	character.Position = position

	return true
}

func TurnCharacterTo(character *Character, position sdl.Point) {
	character.Direction = common.GetFacingDirection(character.Position.Base, position)
}

func FollowPath(character *Character, path *Path) bool {
	if path.Index == len(path.Nodes) {
		return true
	}
	character.MovementTicks++
	if character.MovementTicks >= CharacterMovementTicks {
		node := path.Nodes[path.Index]
		character.MovementTicks = 0
		if !MoveCharacter(character, node.Position) {
			return false
		}
		path.Index++
	}

	return false
}
