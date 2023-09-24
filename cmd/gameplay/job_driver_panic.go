package gameplay

import "math/rand"

type KurinJobDriverPanicData struct{}

func NewKurinJobDriverPanic() *KurinJobDriver {
	return &KurinJobDriver{
		Data:    KurinJobDriverPanicData{},
		Process: ProcessKurinJobDriverPanic,
	}
}

func ProcessKurinJobDriverPanic(driver *KurinJobDriver, game *KurinGame, character *KurinCharacter) bool {
	if driver.Ticks > 60 {
		MoveKurinCharacterDirection(character, KurinDirection(rand.Intn(4)))
		return true
	}

	return false
}
