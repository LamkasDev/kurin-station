package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
	"github.com/veandco/go-sdl2/sdl"
)

func TeleportRandomlyKurinCharacter(character *KurinCharacter) {
	for {
		position := GetRandomMapPosition(&GameInstance.Map)
		if MoveKurinCharacter(character, position) {
			character.PositionRender = sdlutils.PointToFPoint(position.Base)
			break
		}
	}
}

func MoveKurinCharacterDirection(character *KurinCharacter, direction common.KurinDirection) bool {
	tile := GetKurinTileInDirection(&GameInstance.Map, GetKurinTileAt(&GameInstance.Map, character.Position), direction)
	if tile == nil {
		return false
	}

	return MoveKurinCharacter(character, tile.Position)
}

func MoveKurinCharacter(character *KurinCharacter, position sdlutils.Vector3) bool {
	if !CanEnterMapPosition(&GameInstance.Map, position) {
		return false
	}
	TurnKurinCharacterTo(character, position.Base)
	character.Position = position

	return true
}

func TurnKurinCharacterTo(character *KurinCharacter, position sdl.Point) {
	character.Direction = common.GetFacingDirection(character.Position.Base, position)
}

func FollowKurinPath(character *KurinCharacter, path *KurinPath) bool {
	if path.Index == len(path.Nodes) {
		return true
	}
	path.Ticks++
	if path.Ticks > KurinCharacterMovementTicks {
		MoveKurinCharacter(character, path.Nodes[path.Index].Position)
		path.Ticks = 0
		path.Index++
	}

	return false
}
