package serialization

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
)

type JobTrackerData struct {
	Job *JobDriverData
}

func EncodeJobTracker(tracker *gameplay.JobTracker) JobTrackerData {
	data := JobTrackerData{}
	if tracker.Job != nil {
		jobData := EncodeJobDriver(tracker.Job)
		data.Job = &jobData
	}

	return data
}

func DecodeJobTracker(kmap *gameplay.Map, data JobTrackerData, character *gameplay.Mob) {
	character.JobTracker = gameplay.NewJobTracker(character)
	if data.Job != nil {
		gameplay.AssignTrackerJob(character.JobTracker, DecodeJobDriver(kmap, *data.Job))
	}
}
