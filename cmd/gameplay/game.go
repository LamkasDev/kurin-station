package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

var GameInstance *KurinGame

type KurinGame struct {
	Map               KurinMap
	Ticks             uint64
	Credits           uint32
	Characters        []*KurinCharacter
	SelectedCharacter *KurinCharacter

	JobController      *KurinJobController
	ParticleController KurinParticleController
	RunechatController KurinRunechatController
	SoundController    KurinSoundController
	ForceController    KurinForceController
	DialogController   KurinDialogController
	Narrator           *KurinNarrator

	HoveredCharacter *KurinCharacter
	HoveredItem      *KurinItem
}

func NewKurinGame() KurinGame {
	game := KurinGame{
		Map:                NewKurinMap(sdlutils.Vector3{Base: sdl.Point{X: 32, Y: 32}, Z: 1}),
		Ticks:              0,
		Characters:         []*KurinCharacter{},
		JobController:      NewKurinJobController(),
		ParticleController: NewKurinParticleController(),
		RunechatController: NewKurinRunechatController(),
		SoundController:    NewKurinSoundController(),
		ForceController:    NewKurinForceController(),
		DialogController:   NewKurinDialogController(),
		Narrator:           NewKurinNarrator(),
	}
	GameInstance = &game
	PopulateKurinMap(&game.Map)

	playerCharacter := NewKurinCharacter()
	PopulateKurinCharacter(playerCharacter)
	TeleportRandomlyKurinCharacter(playerCharacter)
	game.Characters = append(game.Characters, playerCharacter)
	game.SelectedCharacter = playerCharacter

	npcCharacter := NewKurinCharacter()
	TeleportRandomlyKurinCharacter(npcCharacter)
	game.Characters = append(game.Characters, npcCharacter)

	return game
}

func ProcessKurinGame() {
	for _, object := range GameInstance.Map.Objects {
		object.Process(object)
	}
	for _, character := range GameInstance.Characters {
		ProcessKurinCharacter(character)
	}
	ProcessKurinNarrator()
	GameInstance.Ticks++
}

func TransferKurinItemToCharacter(item *KurinItem, character *KurinCharacter) bool {
	if TransferKurinItemToCharacterRaw(item, &GameInstance.Map, character) {
		delete(GameInstance.ForceController.Forces, item)
		return true
	}

	return false
}

func TransferKurinItemFromCharacter(item *KurinItem, character *KurinCharacter) bool {
	return TransferKurinItemFromCharacterRaw(item, &GameInstance.Map, character)
}

func DropKurinItemFromCharacter(character *KurinCharacter) bool {
	return TransferKurinItemFromCharacter(character.Inventory.Hands[character.ActiveHand], character)
}

func CreateKurinTile(position sdlutils.Vector3, tileType string) *KurinTile {
	if !CanBuildKurinTileAtMapPosition(&GameInstance.Map, position) {
		return nil
	}
	tile := CreateKurinTileRaw(&GameInstance.Map, position, tileType)

	return tile
}

func CreateKurinObject(tile *KurinTile, objectType string) *KurinObject {
	if !CanBuildKurinObjectAtMapPosition(&GameInstance.Map, tile.Position) {
		return nil
	}
	obj := CreateKurinObjectRaw(&GameInstance.Map, tile, objectType)
	obj.OnCreate(obj)
	KurinNarratorOnCreateObject(obj)

	return obj
}

func DestroyKurinObject(obj *KurinObject) {
	DestroyKurinObjectRaw(&GameInstance.Map, obj)
	obj.OnDestroy(obj)
	KurinNarratorOnDestroyObject(obj)
}
