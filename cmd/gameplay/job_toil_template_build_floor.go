package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

type JobToilBuildFloorData struct {
	Position sdlutils.Vector3
	TileType uint8
}

func NewJobToilTemplateBuildFloor() *JobToilTemplate {
	template := NewJobToilTemplate[*JobToilBuildFloorData]("build_floor")
	template.Process = ProcessJobToilBuildFloor

	return template
}

func ProcessJobToilBuildFloor(driver *JobDriver, toil *JobToil) JobToilStatus {
	data := toil.Data.(*JobToilBuildFloorData)
	if toil.Ticks >= 180 {
		PlaySoundVolume(&GameInstance.SoundController, "welder2", 0.35)
		CreateTile(data.Position, data.TileType)
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
