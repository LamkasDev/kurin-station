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
	Mobs              []MobData
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
		Mobs:          []MobData{},
		Narrator:      EncodeNarrator(game.Narrator),
		JobController: map[gameplay.Faction]JobControllerData{},
	}
	for _, mob := range game.Mobs {
		data.Mobs = append(data.Mobs, EncodeMob(mob))
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
	game.Mobs = []*gameplay.Mob{}
	for _, mobData := range data.Mobs {
		game.Mobs = append(game.Mobs, DecodeMob(mobData))
	}
	if data.SelectedCharacter != 0 {
		i := slices.IndexFunc(game.Mobs, func(mob *gameplay.Mob) bool {
			return mob.Id == data.SelectedCharacter
		})
		game.SelectedCharacter = game.Mobs[i]
	}
	game.HoveredMob = nil
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
