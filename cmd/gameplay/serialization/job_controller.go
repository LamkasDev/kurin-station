package serialization

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
)

type JobControllerData struct {
	Jobs []JobDriverData
}

func EncodeJobController(controller *gameplay.JobController) JobControllerData {
	data := JobControllerData{
		Jobs: []JobDriverData{},
	}
	for _, job := range controller.Jobs {
		data.Jobs = append(data.Jobs, EncodeJobDriver(job))
	}

	return data
}

func DecodeJobController(data JobControllerData) *gameplay.JobController {
	controller := gameplay.NewJobController()
	for _, jobData := range data.Jobs {
		gameplay.PushJobToController(controller, DecodeJobDriver(jobData))
	}

	return controller
}
