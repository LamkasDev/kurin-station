package gameplay

import (
	"math/rand"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

const KurinDefaultSpecies = "human"
const KurinDefaultType = "f"

type KurinCharacter struct {
	Id uint32
	Type    string
	Species string
	Position       sdlutils.Vector3
	Direction      KurinDirection
	Fatigue    int32
	ActiveHand KurinHand
	Inventory           KurinInventory

	PositionRender sdl.FPoint
	Movement       sdl.FPoint
	Moving         bool
	Thinktree           KurinCharacterThinktree
	JobTracker          KurinJobTracker
	AnimationController KurinAnimationController
}

func NewKurinCharacter() *KurinCharacter {
	return &KurinCharacter{
		Id: GetNextId(),
		Type:                KurinDefaultType,
		Species:             KurinDefaultSpecies,
		ActiveHand:          KurinHandLeft,
		Fatigue:             0,
		Position:            sdlutils.Vector3{},
		PositionRender:      sdl.FPoint{},
		Movement:            sdl.FPoint{},
		Direction:           KurinDirectionEast,
		Inventory:           NewKurinInventory(),
		Thinktree:           NewKurinCharacterThinktree(),
		JobTracker:          NewKurinJobTracker(),
		AnimationController: NewKurinAnimationController(),
	}
}

func NewKurinCharacterRandom() *KurinCharacter {
	character := NewKurinCharacter()
	character.Inventory.Hands[KurinHandLeft] = NewKurinItem("survivalknife")
	character.Inventory.Hands[KurinHandRight] = NewKurinItem("welder")

	for {
		position := sdlutils.Vector3{Base: sdl.Point{X: int32(rand.Float32() * float32(KurinGameInstance.Map.Size.Base.X)), Y: int32(rand.Float32() * float32(KurinGameInstance.Map.Size.Base.Y))}, Z: 0}
		if MoveKurinCharacter(character, position) {
			character.PositionRender = sdlutils.PointToFPoint(position.Base)
			break
		}
	}

	return character
}

func MoveKurinCharacterDirection(character *KurinCharacter, direction KurinDirection) bool {
	switch direction {
	case KurinDirectionNorth:
		return MoveKurinCharacter(character, sdlutils.Vector3{
			Base: sdl.Point{
				X: character.Position.Base.X,
				Y: character.Position.Base.Y - 1,
			},
			Z: character.Position.Z,
		})
	case KurinDirectionEast:
		return MoveKurinCharacter(character, sdlutils.Vector3{
			Base: sdl.Point{
				X: character.Position.Base.X + 1,
				Y: character.Position.Base.Y,
			},
			Z: character.Position.Z,
		})
	case KurinDirectionSouth:
		return MoveKurinCharacter(character, sdlutils.Vector3{
			Base: sdl.Point{
				X: character.Position.Base.X,
				Y: character.Position.Base.Y + 1,
			},
			Z: character.Position.Z,
		})
	case KurinDirectionWest:
		return MoveKurinCharacter(character, sdlutils.Vector3{
			Base: sdl.Point{
				X: character.Position.Base.X - 1,
				Y: character.Position.Base.Y,
			},
			Z: character.Position.Z,
		})
	}

	return false
}

func MoveKurinCharacter(character *KurinCharacter, position sdlutils.Vector3) bool {
	if !CanEnterPosition(&KurinGameInstance.Map, position) {
		return false
	}

	TurnKurinCharacterTo(character, position.Base)
	character.Position = position
	return true
}

func TurnKurinCharacterTo(character *KurinCharacter, position sdl.Point) {
	character.Direction = GetFacingDirection(sdlutils.PointToFPoint(character.Position.Base), sdlutils.PointToFPoint(position))
}

func FollowKurinPath(character *KurinCharacter, path *KurinPath) bool {
	if path.Index == len(path.Nodes) {
		return true
	}
	path.Ticks++
	if path.Ticks > 30 {
		MoveKurinCharacter(character, path.Nodes[path.Index].Tile.Position)
		path.Ticks = 0
		path.Index++
	}

	return false
}

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

	tile := GetTileAt(&KurinGameInstance.Map, sdlutils.Vector3{Base: position, Z: character.Position.Z})
	if tile == nil || !CanKurinCharacterInteractWithTile(character, tile) {
		return
	}
	item := character.Inventory.Hands[character.ActiveHand]

	hit := true
	if item != nil {
		hit = item.CanHit
		if len(tile.Objects) > 0 {
			object := tile.Objects[0]
			if object.OnItemInteraction(object, item) {
				hit = false
			}
		} else if item.OnTileInteraction(item, tile) {
			hit = false
		}
	}
	if hit && len(tile.Objects) > 0 {
		KurinCharacterHitObject(character, tile.Objects[0])
	}

	if KurinGameInstance.HoveredItem != nil {
		if TransferKurinItemToCharacterRaw(KurinGameInstance.HoveredItem, &KurinGameInstance.Map, character) {
			character.Fatigue += 20
		}
	}
}

func KurinCharacterHitObject(character *KurinCharacter, object *KurinObject) {
	PlayKurinCharacterAnimation(character, "hit")
	character.Fatigue += 60
	HitKurinObject(object)
}

func ProcessKurinCharacter(character *KurinCharacter) {
	if character.Fatigue > 0 {
		character.Fatigue--
	}
	if KurinGameInstance.SelectedCharacter != character {
		if !ProcessKurinJobTracker(character) {
			ProcessKurinCharacterThinktree(character)
		}
	}
}
