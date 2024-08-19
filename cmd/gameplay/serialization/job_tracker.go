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

func DecodeJobTracker(character *gameplay.Character, data JobTrackerData) *gameplay.JobTracker {
	tracker := gameplay.NewJobTracker(character)
	if data.Job != nil {
		gameplay.AssignTrackerJob(tracker, DecodeJobDriver(*data.Job))
	}

	return tracker
}
