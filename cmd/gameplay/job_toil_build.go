package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinJobToilBuildData struct {
	Prefab *KurinObject
}

func NewKurinJobToilBuild(prefab *KurinObject) *KurinJobToil {
	return &KurinJobToil{
		Data: KurinJobToilBuildData{
			Prefab: prefab,
		},
		Start:   func(driver *KurinJobDriver, toil *KurinJobToil) {},
		Process: ProcessKurinJobToilBuild,
	}
}

func ProcessKurinJobToilBuild(driver *KurinJobDriver, toil *KurinJobToil) bool {
	data := toil.Data.(KurinJobToilBuildData)
	if toil.Ticks >= 180 {
		PlaySoundVolume(&KurinGameInstance.SoundController, "welder2", 0.35)
		CreateKurinObject(driver.Tile, data.Prefab.Type)
		return true
	}
	if toil.Ticks%10 == 0 {
		CreateKurinParticle(&KurinGameInstance.ParticleController, NewKurinParticleCross(sdlutils.Vector3ToFVector3Center(driver.Tile.Position), 0.35, sdl.Color{R: 210, G: 210, B: 210}))
	}
	if toil.Ticks%90 == 0 {
		PlaySoundVolume(&KurinGameInstance.SoundController, "welder", 0.5)
	}

	return false
}
