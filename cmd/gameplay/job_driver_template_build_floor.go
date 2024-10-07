package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
)

type JobDriverBuildFloorData struct {
	Position sdlutils.Vector3
	TileType string
}

func NewJobDriverTemplateBuildFloor() *JobDriverTemplate {
	template := NewJobDriverTemplate[*JobDriverBuildFloorData]("build_floor")
	template.Initialize = func(job *JobDriver, data interface{}) {
		buildData := data.(*JobDriverBuildFloorData)
		job.Toils = []*JobToil{}
		job.Toils = append(job.Toils, NewJobToil("goto", &JobToilGotoData{Target: buildData.Position}))
		buildToil := NewJobToil(
			"build_floor",
			&JobToilBuildFloorData{
				Position: buildData.Position,
				TileType: buildData.TileType,
			},
		)
		job.Toils = append(job.Toils, buildToil)
		job.Data = data
	}

	return template
}
