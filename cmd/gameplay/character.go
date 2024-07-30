package gameplay

import (
	"math/rand"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

const KurinDefaultSpecies = "human"
const KurinDefaultType = "f"

type KurinCharacter struct {
	Map     *KurinMap
	Species string
	Type    string

	ActiveHand KurinHand
	Fatigue    int32

	Position       sdlutils.Vector3
	PositionRender sdl.FPoint
	Movement       sdl.FPoint
	Moving         bool
	Direction      KurinDirection

	Inventory           KurinInventory
	Thinktree           KurinCharacterThinktree
	JobTracker          KurinJobTracker
	AnimationController KurinAnimationController
}

func NewKurinCharacterRandom(kmap *KurinMap) *KurinCharacter {
	character := &KurinCharacter{
		Map:                 kmap,
		Species:             KurinDefaultSpecies,
		Type:                KurinDefaultType,
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
	character.Inventory.Hands[KurinHandLeft] = NewKurinItem("survivalknife", nil)
	character.Inventory.Hands[KurinHandRight] = NewKurinItem("welder", nil)

	for {
		position := sdlutils.Vector3{Base: sdl.Point{X: int32(rand.Float32() * float32(kmap.Size.Base.X)), Y: int32(rand.Float32() * float32(kmap.Size.Base.Y))}, Z: 0}
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
	if !CanEnterPosition(character.Map, position) {
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

func InteractKurinCharacter(character *KurinCharacter, game *KurinGame, position sdl.Point) {
	if !character.Moving {
		TurnKurinCharacterTo(character, position)
	}
	if character.Fatigue > 0 {
		return
	}

	tile := GetTileAt(character.Map, sdlutils.Vector3{Base: position, Z: character.Position.Z})
	if tile == nil || !CanKurinCharacterInteractWithTile(character, tile) {
		return
	}
	if len(tile.Objects) > 0 {
		HitObject(game, character, tile, tile.Objects[0])
		return
	}

	if game.HoveredItem != nil {
		if RawTransferKurinItemToCharacter(game.HoveredItem, &game.Map, character) {
			character.Fatigue += 20
		}
	}
}

func HitObject(game *KurinGame, character *KurinCharacter, tile *KurinTile, object *KurinObject) {
	PlayKurinCharacterAnimation(character, "hit")
	PlaySound(&game.SoundController, "grillehit")
	CreateKurinParticle(&game.ParticleController, NewKurinParticleCross(game, sdlutils.Vector3ToFVector3Center(tile.Position)))
	character.Fatigue += 60
}

func ProcessKurinCharacter(game *KurinGame, character *KurinCharacter) {
	if character.Fatigue > 0 {
		character.Fatigue--
	}
	if game.SelectedCharacter != character {
		if !ProcessKurinJobTracker(game, character) {
			ProcessKurinCharacterThinktree(character)
		}
	}
}
