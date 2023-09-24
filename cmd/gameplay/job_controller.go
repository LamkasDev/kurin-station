package gameplay

import (
	"github.com/phf/go-queue/queue"
)

type KurinJobController struct {
	Jobs *queue.Queue
}

func NewKurinJobController() KurinJobController {
	return KurinJobController{
		Jobs: queue.New(),
	}
}

func PushKurinJobToController(controller *KurinJobController, job *KurinJobDriver) bool {
	if job.Tile.Job != nil {
		return false
	}

	controller.Jobs.PushBack(job)
	return true
}

func PopKurinJobFromController(controller *KurinJobController) *KurinJobDriver {
	job := controller.Jobs.PopFront()
	if job == nil {
		return nil
	}

	return job.(*KurinJobDriver)
}
