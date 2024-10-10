package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
)

type JobDriverBuildFloorData struct {
	Position sdlutils.Vector3
	TileType uint8
}

func NewJobDriverTemplateBuildFloor() *JobDriverTemplate {
	template := NewJobDriverTemplate[*JobDriverBuildFloorData]("build_floor")
	template.Initialize = func(job *JobDriver) {
		data := job.Data.(*JobDriverBuildFloorData)
		job.Toils = []*JobToil{}
		job.Toils = append(job.Toils, NewJobToil("goto", &JobToilGotoData{Target: data.Position}))
		buildToil := NewJobToil(
			"build_floor",
			&JobToilBuildFloorData{
				Position: data.Position,
				TileType: data.TileType,
			},
		)
		job.Toils = append(job.Toils, buildToil)
	}

	return template
}
