package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

type JobToilBuildData struct {
	Position   sdlutils.Vector3
	ObjectType string
}

func NewJobToilTemplateBuild() *JobToilTemplate {
	template := NewJobToilTemplate[*JobToilBuildData]("build")
	template.Start = ProcessJobToilBuild
	template.Process = ProcessJobToilBuild

	return template
}

func ProcessJobToilBuild(driver *JobDriver, toil *JobToil) JobToilStatus {
	data := toil.Data.(*JobToilBuildData)
	if toil.Ticks >= 180 {
		PlaySoundVolume(&GameInstance.SoundController, "welder2", 0.35)
		objectTemplate := ObjectContainer[data.ObjectType]
		CreateObject(driver.Tile, data.ObjectType)
		for _, requirement := range objectTemplate.Requirements {
			if item := FindItemInInventory(driver.Mob.Data.(*MobCharacterData).Inventory, requirement.Type); item != nil {
				RemoveItemFromCharacterRaw(item, driver.Mob)
			}
		}
		return JobToilStatusComplete
	}
	if toil.Ticks%10 == 0 {
		particle := NewParticleCross(sdlutils.Vector3ToFVector3Center(data.Position), 0.35, sdl.Color{R: 210, G: 210, B: 210})
		CreateParticle(&GameInstance.ParticleController, particle)
	}
	if toil.Ticks%90 == 0 {
		PlaySoundVolume(&GameInstance.SoundController, "welder", 0.5)
	}

	return JobToilStatusWorking
}
