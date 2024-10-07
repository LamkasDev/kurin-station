package gameplay

import (
	"slices"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
)

type JobController struct {
	Jobs []*JobDriver
}

func NewJobController() *JobController {
	return &JobController{
		Jobs: []*JobDriver{},
	}
}

func PushJobToController(controller *JobController, job *JobDriver) bool {
	if job.Tile != nil {
		if job.Tile.Job != nil {
			return false
		}
		job.Tile.Job = job
	}

	controller.Jobs = append(controller.Jobs, job)
	return true
}

func PopJobFromController(controller *JobController) *JobDriver {
	for i, job := range controller.Jobs {
		if GameInstance.Ticks < job.TimeoutTicks {
			continue
		}

		controller.Jobs = slices.Delete(controller.Jobs, i, i+1)
		return job
	}

	return nil
}

func DoesBuildFloorJobExistAtPosition(position sdlutils.Vector3) bool {
	Matches := func(job *JobDriver) bool {
		switch data := job.Data.(type) {
		case *JobDriverBuildFloorData:
			if sdlutils.CompareVector3(data.Position, position) {
				return true
			}
		}

		return false
	}
	for _, mob := range GameInstance.Mobs {
		if mob.JobTracker.Job != nil && Matches(mob.JobTracker.Job) {
			return true
		}
	}
	for _, job := range GameInstance.JobController[FactionPlayer].Jobs {
		if Matches(job) {
			return true
		}
	}

	return false
}
