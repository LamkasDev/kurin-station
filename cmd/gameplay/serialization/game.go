package serialization

import (
	"slices"

	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/kelindar/binary"
)

type KurinGameData struct {
	NextId uint32

	Map KurinMapData
	Ticks uint64
	Credits uint32
	Characters []KurinCharacterData
	SelectedCharacter uint32
}

func EncodeKurinGame(game *gameplay.KurinGame) []byte {
	data := &KurinGameData{
		NextId: gameplay.NextId,
		Map: EncodeKurinMap(&game.Map),
		Ticks: game.Ticks,
		Credits: game.Credits,
		Characters: []KurinCharacterData{},
	}
	for _, character := range game.Characters {
		data.Characters = append(data.Characters, EncodeKurinCharacter(character))
	}
	if game.SelectedCharacter != nil {
		data.SelectedCharacter = game.SelectedCharacter.Id
	}

	buffer, err := binary.Marshal(data)
	if err != nil {
		panic(err)
	}

	return buffer
}

func DecodeKurinGame(buffer []byte, game *gameplay.KurinGame) {
	var data KurinGameData
	if err := binary.Unmarshal(buffer, &data); err != nil {
		panic(err)
	}

	gameplay.NextId = 0
	game.Map = DecodeKurinMap(data.Map)
	game.Ticks = data.Ticks
	game.Credits = data.Credits
	game.Characters = []*gameplay.KurinCharacter{}
	for _, characterData := range data.Characters {
		game.Characters = append(game.Characters, DecodeKurinCharacter(characterData))
	}
	if data.SelectedCharacter != 0 {
		i := slices.IndexFunc(game.Characters, func(character *gameplay.KurinCharacter) bool {
			return character.Id == data.SelectedCharacter
		})
		game.SelectedCharacter = game.Characters[i]
	}
	game.HoveredCharacter = nil
	game.HoveredItem = nil
	game.JobController = gameplay.NewKurinJobController()
	game.ParticleController = gameplay.NewKurinParticleController()
	game.RunechatController = gameplay.NewKurinRunechatController()
	game.SoundController = gameplay.NewKurinSoundController()
	game.ForceController = gameplay.NewKurinForceController()
	game.Narrator = gameplay.NewKurinNarrator()
	gameplay.NextId = data.NextId
}
