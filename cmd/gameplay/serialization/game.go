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
	Characters        []CharacterData
	SelectedCharacter uint32
	Narrator          NarratorData
	JobController     map[gameplay.Faction]JobControllerData
}

func EncodeGame(game *gameplay.Game) []byte {
	data := &GameData{
		NextId:        gameplay.NextId,
		Map:           EncodeMap(&game.Map),
		Ticks:         game.Ticks,
		Credits:       game.Credits,
		Characters:    []CharacterData{},
		Narrator:      EncodeNarrator(game.Narrator),
		JobController: map[gameplay.Faction]JobControllerData{},
	}
	for _, character := range game.Characters {
		data.Characters = append(data.Characters, EncodeCharacter(character))
	}
	if game.SelectedCharacter != nil {
		data.SelectedCharacter = game.SelectedCharacter.Id
	}
	for faction, controller := range game.JobController {
		data.JobController[faction] = EncodeJobController(controller)
	}

	buffer, err := binary.Marshal(data)
	if err != nil {
		panic(err)
	}

	return buffer
}

func DecodeGame(buffer []byte, game *gameplay.Game) {
	var data GameData
	if err := binary.Unmarshal(buffer, &data); err != nil {
		panic(err)
	}

	gameplay.NextId = 0
	game.Map = DecodeMap(data.Map)
	game.Ticks = data.Ticks
	game.Credits = data.Credits
	game.Characters = []*gameplay.Character{}
	for _, characterData := range data.Characters {
		game.Characters = append(game.Characters, DecodeCharacter(characterData))
	}
	if data.SelectedCharacter != 0 {
		i := slices.IndexFunc(game.Characters, func(character *gameplay.Character) bool {
			return character.Id == data.SelectedCharacter
		})
		game.SelectedCharacter = game.Characters[i]
	}
	game.HoveredCharacter = nil
	game.HoveredItem = nil
	for faction, controllerData := range data.JobController {
		game.JobController[faction] = DecodeJobController(controllerData)
	}
	game.ParticleController = gameplay.NewParticleController()
	game.RunechatController = gameplay.NewRunechatController()
	game.SoundController = gameplay.NewSoundController()
	game.ForceController = gameplay.NewForceController()
	game.Narrator = DecodeNarrator(data.Narrator)
	gameplay.NextId = data.NextId
}
