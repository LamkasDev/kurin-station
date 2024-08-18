package gameplay

type KurinJobTracker struct {
	Character *KurinCharacter
	Job       *KurinJobDriver
}

func NewKurinJobTracker(character *KurinCharacter) *KurinJobTracker {
	return &KurinJobTracker{
		Character: character,
		Job:       nil,
	}
}

func ProcessKurinJobTracker(tracker *KurinJobTracker) bool {
	if tracker.Job == nil {
		job := PopKurinJobFromController(GameInstance.JobController)
		if job == nil {
			return false
		}
		AssignKurinTrackerJob(tracker, job)
	}

	toil := tracker.Job.Toils[tracker.Job.ToilIndex]
	if !toil.Started {
		if !StartKurinTrackerJobToil(tracker) {
			return true
		}
	}
	switch toil.Process(tracker.Job, toil) {
	case KurinJobToilStatusComplete:
		EndKurinTrackerJobToil(tracker)
		AdvanceKurinTrackerJob(tracker)
		return true
	case KurinJobToilStatusFailed:
		EndKurinTrackerJobToil(tracker)
		TimeoutKurinTrackerJob(tracker)
		return true
	}
	toil.Ticks++

	return true
}

func StartKurinTrackerJobToil(tracker *KurinJobTracker) bool {
	toil := tracker.Job.Toils[tracker.Job.ToilIndex]
	status := toil.Start(tracker.Job, toil)
	toil.Started = true
	switch status {
	case KurinJobToilStatusFailed:
		EndKurinTrackerJobToil(tracker)
		TimeoutKurinTrackerJob(tracker)
		return false
	case KurinJobToilStatusComplete:
		EndKurinTrackerJobToil(tracker)
		if !AdvanceKurinTrackerJob(tracker) {
			return true
		}
		return StartKurinTrackerJobToil(tracker)
	}

	return true
}

func EndKurinTrackerJobToil(tracker *KurinJobTracker) {
	toil := tracker.Job.Toils[tracker.Job.ToilIndex]
	toil.End(tracker.Job, toil)
}

func AssignKurinTrackerJob(tracker *KurinJobTracker, job *KurinJobDriver) {
	job.Character = tracker.Character
	tracker.Job = job
}

func UnassignKurinTrackerJob(tracker *KurinJobTracker) {
	if tracker.Job.Tile.Job == tracker.Job {
		tracker.Job.Tile.Job = nil
	}
	tracker.Job.Character = nil
	for _, toil := range tracker.Job.Toils {
		toil.Started = false
	}
	tracker.Job = nil
}

func AdvanceKurinTrackerJob(tracker *KurinJobTracker) bool {
	tracker.Job.ToilIndex++
	if tracker.Job.ToilIndex >= uint8(len(tracker.Job.Toils)) {
		UnassignKurinTrackerJob(tracker)
		return false
	}

	return true
}

func TimeoutKurinTrackerJob(tracker *KurinJobTracker) {
	job := tracker.Job
	job.TimeoutTicks = GameInstance.Ticks + 300
	UnassignKurinTrackerJob(tracker)
	PushKurinJobToController(GameInstance.JobController, job)
}
