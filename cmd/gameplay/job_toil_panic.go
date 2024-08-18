package gameplay

import (
	"math/rand"

	"github.com/LamkasDev/kurin/cmd/gameplay/common"
)

func NewKurinJobToilPanic() *KurinJobToil {
	toil := NewKurinJobToilRaw[interface{}]("panic")
	toil.Process = ProcessKurinJobToilPanic

	return toil
}

func ProcessKurinJobToilPanic(driver *KurinJobDriver, toil *KurinJobToil) KurinJobToilStatus {
	if toil.Ticks > 60 {
		MoveKurinCharacterDirection(driver.Character, common.KurinDirection(rand.Intn(4)))
		return KurinJobToilStatusComplete
	}

	return KurinJobToilStatusWorking
}
