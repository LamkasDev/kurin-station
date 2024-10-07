package gameplay

import (
	"math/rand"

	"github.com/LamkasDev/kurin/cmd/gameplay/common"
)

func NewJobToilTemplatePanic() *JobToilTemplate {
	template := NewJobToilTemplate[interface{}]("panic")
	template.Process = ProcessJobToilPanic

	return template
}

func ProcessJobToilPanic(driver *JobDriver, toil *JobToil) JobToilStatus {
	if toil.Ticks >= MobMovementTicks {
		MoveMobDirection(driver.Mob, common.Direction(rand.Intn(4)))
		return JobToilStatusComplete
	}

	return JobToilStatusWorking
}
