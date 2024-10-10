package serialization

import (
	"slices"

	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/kelindar/binary"
)

type GameData struct {
	NextId uint32

	Map               MapData
	Ticks             uint64
	Credits           uint32
	SelectedCharacter uint32
	SelectedZ         uint8
	Narrator          NarratorData
	JobController     map[gameplay.Faction]JobControllerData
}

func EncodeGame() []byte {
	data := &GameData{
		NextId:        gameplay.NextId,
		Map:           EncodeMap(gameplay.GameInstance.Map),
		Ticks:         gameplay.GameInstance.Ticks,
		Credits:       gameplay.GameInstance.Credits,
		SelectedZ:     gameplay.GameInstance.SelectedZ,
		Narrator:      EncodeNarrator(gameplay.GameInstance.Narrator),
		JobController: map[gameplay.Faction]JobControllerData{},
	}
	if gameplay.GameInstance.SelectedCharacter != nil {
		data.SelectedCharacter = gameplay.GameInstance.SelectedCharacter.Id
	}
	for faction, controller := range gameplay.GameInstance.JobController {
		data.JobController[faction] = EncodeJobController(controller)
	}

	buffer, err := binary.Marshal(data)
	if err != nil {
		panic(err)
	}

	return buffer
}

func PredecodeGame(data GameData) {
	gameplay.NextId = 0
	gameplay.GameInstance.Map = PredecodeMap(data.Map)
	gameplay.GameInstance.Ticks = data.Ticks
	gameplay.GameInstance.Credits = data.Credits
	if data.SelectedCharacter != 0 {
		i := slices.IndexFunc(gameplay.GameInstance.Map.Mobs, func(mob *gameplay.Mob) bool {
			return mob.Id == data.SelectedCharacter
		})
		gameplay.GameInstance.SelectedCharacter = gameplay.GameInstance.Map.Mobs[i]
	}
	gameplay.GameInstance.SelectedZ = data.SelectedZ
	gameplay.GameInstance.HoveredMob = nil
	gameplay.GameInstance.HoveredItem = nil
	for faction, controllerData := range data.JobController {
		gameplay.GameInstance.JobController[faction] = DecodeJobController(gameplay.GameInstance.Map, controllerData)
	}
	gameplay.GameInstance.ParticleController = gameplay.NewParticleController()
	gameplay.GameInstance.RunechatController = gameplay.NewRunechatController()
	gameplay.GameInstance.SoundController = gameplay.NewSoundController()
	gameplay.GameInstance.ForceController = gameplay.NewForceController()
	gameplay.GameInstance.Narrator = DecodeNarrator(data.Narrator)
	gameplay.NextId = data.NextId
}

func DecodeGame(data GameData) {
	DecodeMap(gameplay.GameInstance.Map, data.Map)
}
