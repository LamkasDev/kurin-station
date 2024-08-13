package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

var KurinGameInstance *KurinGame 

type KurinGame struct {
	Map KurinMap
	Ticks uint64
	Credits uint32
	Characters        []*KurinCharacter
	SelectedCharacter *KurinCharacter

	JobController      KurinJobController
	ParticleController KurinParticleController
	RunechatController KurinRunechatController
	SoundController    KurinSoundController
	ForceController KurinForceController
	Narrator KurinNarrator
	
	HoveredCharacter  *KurinCharacter
	HoveredItem *KurinItem
}

func NewKurinGame() KurinGame {
	game := KurinGame{
		Map:                NewKurinMap(sdlutils.Vector3{Base: sdl.Point{X: 25, Y: 25}, Z: 1}),
		Ticks: 0,
		Characters:         []*KurinCharacter{},
		JobController:      NewKurinJobController(),
		ParticleController: NewKurinParticleController(),
		RunechatController: NewKurinRunechatController(),
		SoundController:    NewKurinSoundController(),
		ForceController: NewKurinForceController(),
		Narrator: NewKurinNarrator(),
	}
	KurinGameInstance = &game
	PopulateKurinMap(&game.Map)
	for i := 0; i < 2; i++ {
		character := NewKurinCharacterRandom()
		game.Characters = append(game.Characters, character)
		game.SelectedCharacter = character
	}

	return game
}

func ProcessKurinGame() {
	for _, character := range KurinGameInstance.Characters {
		ProcessKurinCharacter(character)
	}
	ProcessKurinNarrator()
	KurinGameInstance.Ticks++
}

func TransferKurinItemToCharacter(item *KurinItem, character *KurinCharacter) bool {
	if IsKurinCharacterHandEmptyRaw(character) {
		return false
	}
	if TransferKurinItemFromCharacterRaw(item, &KurinGameInstance.Map, character) {
		delete(KurinGameInstance.ForceController.Forces, item)
		return true
	}

	return false
}

func TransferKurinItemFromCharacter(item *KurinItem, character *KurinCharacter) bool {
	if IsKurinCharacterHandEmptyRaw(character) {
		return false
	}

	return TransferKurinItemFromCharacterRaw(item, &KurinGameInstance.Map, character)
}

func DropKurinItemFromCharacter(character *KurinCharacter) bool {
	return TransferKurinItemFromCharacter(character.Inventory.Hands[character.ActiveHand], character)
}

func CreateKurinObject(tile *KurinTile, objectType string) {
	obj := CreateKurinObjectRaw(&KurinGameInstance.Map, tile, objectType)
	obj.OnCreate(obj)
	KurinNarratorOnCreateObject(obj)
}

func DestroyKurinObject(obj *KurinObject) {
	DestroyKurinObjectRaw(&KurinGameInstance.Map, obj)
	obj.OnDestroy(obj)
	KurinNarratorOnDestroyObject(obj)
}
