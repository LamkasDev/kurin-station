package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	KurinDefaultSpecies         = "human"
	KurinDefaultType            = "f"
	KurinCharacterMovementTicks = 30
)

type KurinCharacter struct {
	Id         uint32
	Type       string
	Species    string
	Position   sdlutils.Vector3
	Direction  common.KurinDirection
	Fatigue    int32
	ActiveHand KurinHand
	Inventory  KurinInventory
	JobTracker *KurinJobTracker

	PositionRender      sdl.FPoint
	Movement            sdl.FPoint
	Moving              bool
	Thinktree           KurinCharacterThinktree
	AnimationController KurinAnimationController
}

func NewKurinCharacter() *KurinCharacter {
	character := &KurinCharacter{
		Id:                  GetNextId(),
		Type:                KurinDefaultType,
		Species:             KurinDefaultSpecies,
		ActiveHand:          KurinHandLeft,
		Fatigue:             0,
		Position:            sdlutils.Vector3{},
		PositionRender:      sdl.FPoint{},
		Movement:            sdl.FPoint{},
		Direction:           common.KurinDirectionEast,
		Inventory:           NewKurinInventory(),
		Thinktree:           NewKurinCharacterThinktree(),
		AnimationController: NewKurinAnimationController(),
	}
	character.JobTracker = NewKurinJobTracker(character)

	return character
}

func PopulateKurinCharacter(character *KurinCharacter) {
	character.Inventory.Hands[KurinHandLeft] = NewKurinItem("survivalknife", 1)
	character.Inventory.Hands[KurinHandRight] = NewKurinItem("welder", 1)
}

func ProcessKurinCharacter(character *KurinCharacter) {
	if character.Fatigue > 0 {
		character.Fatigue--
	}
	if GameInstance.SelectedCharacter != character {
		if !ProcessKurinJobTracker(character.JobTracker) {
			ProcessKurinCharacterThinktree(character)
		}
	}
}
