package gameplay

import "math/rand"

func NewKurinJobToilPanic() *KurinJobToil {
	return &KurinJobToil{
		Start:   func(driver *KurinJobDriver, toil *KurinJobToil) {},
		Process: ProcessKurinJobToilBuild,
	}
}

func ProcessKurinJobToilPanic(driver *KurinJobDriver, toil *KurinJobToil) bool {
	if toil.Ticks > 60 {
		MoveKurinCharacterDirection(driver.Character, KurinDirection(rand.Intn(4)))
		return true
	}

	return false
}
