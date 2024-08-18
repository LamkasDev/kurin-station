package serialization

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
)

type KurinJobControllerData struct {
	Jobs []KurinJobDriverData
}

func EncodeKurinJobController(controller *gameplay.KurinJobController) KurinJobControllerData {
	data := KurinJobControllerData{
		Jobs: []KurinJobDriverData{},
	}
	for _, job := range controller.Jobs {
		data.Jobs = append(data.Jobs, EncodeKurinJobDriver(job))
	}

	return data
}

func DecodeKurinJobController(data KurinJobControllerData) *gameplay.KurinJobController {
	controller := gameplay.NewKurinJobController()
	for _, jobData := range data.Jobs {
		gameplay.PushKurinJobToController(controller, DecodeKurinJobDriver(jobData))
	}

	return controller
}
