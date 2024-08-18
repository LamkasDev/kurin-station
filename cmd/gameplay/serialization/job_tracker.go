package serialization

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
)

type KurinJobTrackerData struct {
	Job *KurinJobDriverData
}

func EncodeKurinJobTracker(tracker *gameplay.KurinJobTracker) KurinJobTrackerData {
	data := KurinJobTrackerData{}
	if tracker.Job != nil {
		jobData := EncodeKurinJobDriver(tracker.Job)
		data.Job = &jobData
	}

	return data
}

func DecodeKurinJobTracker(character *gameplay.KurinCharacter, data KurinJobTrackerData) *gameplay.KurinJobTracker {
	tracker := gameplay.NewKurinJobTracker(character)
	if data.Job != nil {
		gameplay.AssignKurinTrackerJob(tracker, DecodeKurinJobDriver(*data.Job))
	}

	return tracker
}
