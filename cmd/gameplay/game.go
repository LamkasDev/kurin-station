package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinGame struct {
	Map KurinMap
	Ticks uint64
	Characters        []*KurinCharacter

	SelectedCharacter *KurinCharacter
	HoveredCharacter  *KurinCharacter
	HoveredItem *KurinItem

	JobController      KurinJobController
	ParticleController KurinParticleController
	RunechatController KurinRunechatController
	SoundController    KurinSoundController
	ForceController KurinForceController
}

type KurinGameMarshal struct {
	Map KurinMapMarshal
}

func NewKurinGame() KurinGame {
	kmap := NewKurinMap(sdlutils.Vector3{Base: sdl.Point{X: 25, Y: 25}, Z: 1})
	game := KurinGame{
		Map:                kmap,
		Ticks: 0,
		Characters:         []*KurinCharacter{},
		JobController:      NewKurinJobController(),
		ParticleController: NewKurinParticleController(),
		SoundController:    NewKurinSoundController(),
		ForceController: NewKurinForceController(),
	}
	for i := 0; i < 2; i++ {
		character := NewKurinCharacterRandom(&kmap)
		game.Characters = append(game.Characters, character)
		game.SelectedCharacter = character
	}

	return game
}

func ProcessKurinGame(game *KurinGame) {
	for _, character := range game.Characters {
		ProcessKurinCharacter(game, character)
	}
	game.Ticks++
}

func TransferKurinItemToCharacter(game *KurinGame, item *KurinItem, character *KurinCharacter) bool {
	if RawIsKurinCharacterHandEmpty(character) {
		return false
	}
	if RawTransferKurinItemFromCharacter(item, &game.Map, character) {
		delete(game.ForceController.Forces, item)
		return true
	}

	return false
}

func TransferKurinItemFromCharacter(game *KurinGame, item *KurinItem, character *KurinCharacter) bool {
	if RawIsKurinCharacterHandEmpty(character) {
		return false
	}

	return RawTransferKurinItemFromCharacter(item, &game.Map, character)
}

func DropKurinItemFromCharacter(game *KurinGame, character *KurinCharacter) bool {
	return TransferKurinItemFromCharacter(game, character.Inventory.Hands[character.ActiveHand], character)
}

func MarshalKurinGame(game *KurinGame) (KurinGameMarshal, *error) {
	mgame := KurinGameMarshal{}
	var err *error
	if mgame.Map, err = MarshalKurinMap(&game.Map); err != nil {
		return mgame, err
	}

	return mgame, nil
}
