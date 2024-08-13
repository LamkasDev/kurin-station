package gameplay

import (
	"math/rand"
	"slices"
)

type KurinJobTracker struct {
	Job *KurinJobDriver
}

func NewKurinJobTracker() KurinJobTracker {
	return KurinJobTracker{
		Job: nil,
	}
}

func ProcessKurinJobTracker(character *KurinCharacter) bool {
	if character.JobTracker.Job == nil {
		character.JobTracker.Job = PopKurinJobFromController(&KurinGameInstance.JobController)
		if character.JobTracker.Job == nil {
			return false
		}
		character.JobTracker.Job.Character = character
		character.JobTracker.Job.Toils[0].Start(character.JobTracker.Job, character.JobTracker.Job.Toils[0])
		if rand.Intn(10) < 4 {
			switch character.JobTracker.Job.Type {
			case "build":
				CreateKurinRunechatMessage(&KurinGameInstance.RunechatController, NewKurinRunechatCharacter(character, "Work, work."))
			}
		}
	}
	toil := character.JobTracker.Job.Toils[0]
	if toil.Process(character.JobTracker.Job, toil) {
		character.JobTracker.Job.Toils = slices.Delete(character.JobTracker.Job.Toils, 0, 1)
		if len(character.JobTracker.Job.Toils) == 0 {
			if character.JobTracker.Job.Tile.Job == character.JobTracker.Job {
				character.JobTracker.Job.Tile.Job = nil
			}
			character.JobTracker.Job = nil
			return true
		}
		character.JobTracker.Job.Toils[0].Start(character.JobTracker.Job, character.JobTracker.Job.Toils[0])
	}
	toil.Ticks++

	return true
}
