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

func DecodeJobTracker(kmap *gameplay.Map, data JobTrackerData, character *gameplay.Mob) *gameplay.JobTracker {
	tracker := gameplay.NewJobTracker(character)
	if data.Job != nil {
		gameplay.AssignTrackerJob(tracker, DecodeJobDriver(kmap, *data.Job))
	}

	return tracker
}
