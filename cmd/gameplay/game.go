package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

var GameInstance *Game

type Game struct {
	Map               Map
	Ticks             uint64
	Credits           uint32
	Mobs              []*Mob
	SelectedCharacter *Mob

	JobController         map[Faction]*JobController
	ParticleController    ParticleController
	RunechatController    RunechatController
	SoundController       SoundController
	ForceController       ForceController
	DialogController      DialogController
	ReservationController ReservationController
	Narrator              *Narrator

	HoveredTile   *Tile
	HoveredObject *Object
	HoveredMob    *Mob
	HoveredItem   *Item
	Godmode       bool
}

func InitializeGame() {
	GameInstance = &Game{
		Map:   NewMap(sdlutils.Vector3{Base: sdl.Point{X: 50, Y: 50}, Z: 1}),
		Ticks: 0,
		Mobs:  []*Mob{},
		JobController: map[Faction]*JobController{
			FactionPlayer: NewJobController(),
			FactionTrader: NewJobController(),
		},
		ParticleController:    NewParticleController(),
		RunechatController:    NewRunechatController(),
		SoundController:       NewSoundController(),
		ForceController:       NewForceController(),
		DialogController:      NewDialogController(),
		ReservationController: NewReservationController(),
		Narrator:              NewNarrator(),
	}
	RegisterItems()
	RegisterObjects()
	RegisterMobs()
	RegisterJobToils()
	RegisterJobDrivers()
	RegisterObjectiveRequirements()
	PopulateMap(&GameInstance.Map)

	playerCharacter := NewMob("character", FactionPlayer)
	PopulateCharacter(playerCharacter)
	TeleportMobRandom(playerCharacter)
	GameInstance.Mobs = append(GameInstance.Mobs, playerCharacter)
	GameInstance.SelectedCharacter = playerCharacter

	npcCharacter := NewMob("character", FactionPlayer)
	TeleportMobRandom(npcCharacter)
	GameInstance.Mobs = append(GameInstance.Mobs, npcCharacter)

	cat := NewMob("cat", FactionPlayer)
	TeleportMobRandom(cat)
	GameInstance.Mobs = append(GameInstance.Mobs, cat)
}

func ProcessGame() {
	for _, object := range GameInstance.Map.Objects {
		object.Template.Process(object)
	}
	for _, mob := range GameInstance.Mobs {
		mob.Template.Process(mob)
	}
	ProcessNarrator()
	GameInstance.Ticks++
}

func AddCredits(amount uint32) {
	GameInstance.Credits += amount
	NarratorOnAddCredits(amount)
}

func TransferItemToCharacter(item *Item, character *Mob) bool {
	if TransferItemToCharacterRaw(item, &GameInstance.Map, character) {
		delete(GameInstance.ForceController.Forces, item)
		return true
	}

	return false
}

func TransferItemFromCharacter(item *Item, character *Mob) bool {
	return TransferItemFromCharacterRaw(item, &GameInstance.Map, character)
}

func DropItemFromCharacter(character *Mob) bool {
	item := GetHeldItem(character)
	if item == nil {
		return true
	}

	return TransferItemFromCharacter(item, character)
}

func CreateTile(position sdlutils.Vector3, tileType string) *Tile {
	if !CanBuildTileAtMapPosition(&GameInstance.Map, position) {
		return nil
	}
	tile := CreateTileRaw(&GameInstance.Map, position, tileType)

	return tile
}

func DestroyTile(tile *Tile) {
	DestroyTileRaw(&GameInstance.Map, tile)
}

func CreateObject(tile *Tile, objectType string) *Object {
	if !CanBuildObjectAtMapPosition(&GameInstance.Map, tile.Position) {
		return nil
	}
	obj := CreateObjectRaw(&GameInstance.Map, tile, objectType)
	obj.Template.OnCreate(obj)
	NarratorOnCreateObject(obj)

	return obj
}

func DestroyObject(obj *Object) {
	DestroyObjectRaw(&GameInstance.Map, obj)
	obj.Template.OnDestroy(obj)
	NarratorOnDestroyObject(obj)
}
