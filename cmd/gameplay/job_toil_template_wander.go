package gameplay

import (
	"math/rand"

	"github.com/LamkasDev/kurin/cmd/gameplay/common"
)

func NewJobToilTemplateWander() *JobToilTemplate {
	template := NewJobToilTemplate[interface{}]("wander")
	template.Process = ProcessJobToilWander

	return template
}

func ProcessJobToilWander(driver *JobDriver, toil *JobToil) JobToilStatus {
	if toil.Ticks >= MobMovementTicks*5 {
		MoveMobDirection(driver.Mob, common.Direction(rand.Intn(4)))
		return JobToilStatusComplete
	}

	return JobToilStatusWorking
}
