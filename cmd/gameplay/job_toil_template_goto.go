package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
)

type JobToilGotoData struct {
	Target sdlutils.Vector3
	Path   *Path
}

func NewJobToilTemplateGoto() *JobToilTemplate {
	template := NewJobToilTemplate[JobToilGotoData]("goto")
	template.Start = StartJobToilGoto
	template.Process = ProcessJobToilGoto

	return template
}

func StartJobToilGoto(driver *JobDriver, toil *JobToil) JobToilStatus {
	data := toil.Data.(*JobToilGotoData)
	data.Path = FindPathAdjacent(&GameInstance.Map.Pathfinding, driver.Mob.Position, data.Target)
	if data.Path == nil {
		return JobToilStatusFailed
	}

	return JobToilStatusWorking
}

func ProcessJobToilGoto(driver *JobDriver, toil *JobToil) JobToilStatus {
	data := toil.Data.(*JobToilGotoData)
	if FollowPath(driver.Mob, data.Path) {
		return JobToilStatusComplete
	}

	return JobToilStatusWorking
}
