package gameplay

import "math/rand"

type KurinJobTracker struct {
	Job *KurinJobDriver
}

func NewKurinJobTracker() KurinJobTracker {
	return KurinJobTracker{
		Job: nil,
	}
}

func ProcessKurinJobTracker(game *KurinGame, character *KurinCharacter) bool {
	if character.JobTracker.Job == nil {
		character.JobTracker.Job = PopKurinJobFromController(&game.JobController)
		if character.JobTracker.Job == nil {
			return false
		}
		if character.JobTracker.Job.Assign != nil {
			character.JobTracker.Job.Assign(character.JobTracker.Job, game, character)
		}
		if rand.Intn(10) < 4 {
			switch character.JobTracker.Job.Data.(type) {
			case KurinJobDriverBuildData:
				CreateKurinRunechatMessage(&game.RunechatController, NewKurinRunechatCharacter(character, "Work, work."))
			}
		}
	}
	if character.JobTracker.Job != nil {
		if character.JobTracker.Job.Process(character.JobTracker.Job, game, character) {
			if character.JobTracker.Job.Tile.Job == character.JobTracker.Job {
				character.JobTracker.Job.Tile.Job = nil
			}
			character.JobTracker.Job = nil
			return true
		}
		character.JobTracker.Job.Ticks++
	}

	return true
}
