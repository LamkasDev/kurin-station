package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	DefaultSpecies         = "human"
	DefaultGender          = "f"
	CharacterMovementTicks = 10
)

type Character struct {
	Id         uint32
	Species    string
	Gender     string
	Faction    Faction
	Position   sdlutils.Vector3
	Direction  common.Direction
	Fatigue    int32
	ActiveHand Hand
	Inventory  Inventory
	JobTracker *JobTracker

	PositionRender      sdl.FPoint
	Movement            sdl.Point
	MovementTicks       uint8
	Thinktree           CharacterThinktree
	AnimationController AnimationController
}

func NewCharacter(faction Faction) *Character {
	character := &Character{
		Id:                  GetNextId(),
		Species:             DefaultSpecies,
		Gender:              DefaultGender,
		Faction:             faction,
		Position:            sdlutils.Vector3{},
		Direction:           common.DirectionEast,
		Fatigue:             0,
		ActiveHand:          HandLeft,
		Inventory:           NewInventory(),
		PositionRender:      sdl.FPoint{},
		Movement:            sdl.Point{},
		MovementTicks:       0,
		Thinktree:           NewCharacterThinktree(),
		AnimationController: NewAnimationController(),
	}
	character.JobTracker = NewJobTracker(character)

	return character
}

func PopulateCharacter(character *Character) {
	character.Inventory.Hands[HandLeft] = NewItem("survivalknife", 1)
	character.Inventory.Hands[HandRight] = NewItem("welder", 1)
}

func ProcessCharacter(character *Character) {
	if character.Fatigue > 0 {
		character.Fatigue--
	}
	if GameInstance.SelectedCharacter != character {
		if !ProcessJobTracker(character.JobTracker) {
			ProcessCharacterThinktree(character)
		}
	}
}
