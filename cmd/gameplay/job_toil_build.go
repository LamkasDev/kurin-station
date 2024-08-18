package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinJobToilBuildData struct {
	Prefab string
}

func NewKurinJobToilBuild(prefab string) *KurinJobToil {
	toil := NewKurinJobToilRaw[KurinJobToilBuildData]("build")
	toil.Start = ProcessKurinJobToilBuild
	toil.Process = ProcessKurinJobToilBuild
	toil.Data = &KurinJobToilBuildData{
		Prefab: prefab,
	}

	return toil
}

func ProcessKurinJobToilBuild(driver *KurinJobDriver, toil *KurinJobToil) KurinJobToilStatus {
	data := toil.Data.(*KurinJobToilBuildData)
	if toil.Ticks >= 180 {
		PlaySoundVolume(&GameInstance.SoundController, "welder2", 0.35)
		CreateKurinObject(driver.Tile, data.Prefab)
		if item := FindItemInInventory(&driver.Character.Inventory, "rod"); item != nil {
			RemoveKurinItemFromCharacterRaw(item, driver.Character)
		}
		return KurinJobToilStatusComplete
	}
	if toil.Ticks%10 == 0 {
		CreateKurinParticle(&GameInstance.ParticleController, NewKurinParticleCross(sdlutils.Vector3ToFVector3Center(driver.Tile.Position), 0.35, sdl.Color{R: 210, G: 210, B: 210}))
	}
	if toil.Ticks%90 == 0 {
		PlaySoundVolume(&GameInstance.SoundController, "welder", 0.5)
	}

	return KurinJobToilStatusWorking
}
