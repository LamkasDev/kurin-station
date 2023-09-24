package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinGame struct {
	Map KurinMap

	Characters        []*KurinCharacter
	SelectedCharacter *KurinCharacter
	HoveredCharacter  *KurinCharacter

	Items       []*KurinItem
	HoveredItem *KurinItem

	JobController      KurinJobController
	ParticleController KurinParticleController
	RunechatController KurinRunechatController
	SoundController    KurinSoundController
}

type KurinGameMarshal struct {
	Map KurinMapMarshal
}

func NewKurinGame() KurinGame {
	kmap := NewKurinMap(sdlutils.Vector3{Base: sdl.Point{X: 25, Y: 25}, Z: 1})
	game := KurinGame{
		Map:                kmap,
		Characters:         []*KurinCharacter{},
		Items:              []*KurinItem{},
		JobController:      NewKurinJobController(),
		ParticleController: NewKurinParticleController(),
		SoundController:    NewKurinSoundController(),
	}
	for i := 0; i < 2; i++ {
		character := NewKurinCharacterRandom(&kmap)
		game.Characters = append(game.Characters, character)
		game.SelectedCharacter = character
	}
	for i := 0; i < 10; i++ {
		item := NewKurinItemRandom("survivalknife", &kmap)
		game.Items = append(game.Items, item)
	}

	return game
}

func MarshalKurinGame(game *KurinGame) (KurinGameMarshal, *error) {
	mgame := KurinGameMarshal{}
	var err *error
	if mgame.Map, err = MarshalKurinMap(&game.Map); err != nil {
		return mgame, err
	}

	return mgame, nil
}
