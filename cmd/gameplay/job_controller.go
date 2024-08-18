package gameplay

import "slices"

type KurinJobController struct {
	Jobs []*KurinJobDriver
}

func NewKurinJobController() *KurinJobController {
	return &KurinJobController{
		Jobs: []*KurinJobDriver{},
	}
}

func PushKurinJobToController(controller *KurinJobController, job *KurinJobDriver) bool {
	if job.Tile != nil {
		if job.Tile.Job != nil {
			return false
		}
		job.Tile.Job = job
	}

	controller.Jobs = append(controller.Jobs, job)
	return true
}

func PopKurinJobFromController(controller *KurinJobController) *KurinJobDriver {
	for i, job := range controller.Jobs {
		if GameInstance.Ticks < job.TimeoutTicks {
			continue
		}

		controller.Jobs = slices.Delete(controller.Jobs, i, i+1)
		return job
	}

	return nil
}
