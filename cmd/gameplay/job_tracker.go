package gameplay

type JobTracker struct {
	Character *Mob
	Job       *JobDriver
}

func NewJobTracker(character *Mob) *JobTracker {
	return &JobTracker{
		Character: character,
		Job:       nil,
	}
}

func ProcessJobTracker(tracker *JobTracker) bool {
	if tracker.Job == nil {
		job := PopJobFromController(GameInstance.JobController[tracker.Character.Faction])
		if job == nil {
			return false
		}
		AssignTrackerJob(tracker, job)
	}

	toil := tracker.Job.Toils[tracker.Job.ToilIndex]
	if !toil.Started {
		if !StartTrackerJobToil(tracker) {
			return true
		}
	}
	status := toil.Template.Process(tracker.Job, toil)
	switch status {
	case JobToilStatusComplete:
		EndTrackerJobToil(tracker)
		AdvanceTrackerJob(tracker)
		return true
	case JobToilStatusFailed:
		EndTrackerJobToil(tracker)
		TimeoutTrackerJob(tracker)
		return true
	}
	toil.Ticks++

	return true
}

func StartTrackerJobToil(tracker *JobTracker) bool {
	toil := tracker.Job.Toils[tracker.Job.ToilIndex]
	status := toil.Template.Start(tracker.Job, toil)
	toil.Started = true
	switch status {
	case JobToilStatusFailed:
		EndTrackerJobToil(tracker)
		TimeoutTrackerJob(tracker)
		return false
	case JobToilStatusComplete:
		EndTrackerJobToil(tracker)
		if !AdvanceTrackerJob(tracker) {
			return true
		}
		return StartTrackerJobToil(tracker)
	}

	return true
}

func EndTrackerJobToil(tracker *JobTracker) {
	toil := tracker.Job.Toils[tracker.Job.ToilIndex]
	toil.Template.End(tracker.Job, toil)
}

func AssignTrackerJob(tracker *JobTracker, job *JobDriver) {
	job.Mob = tracker.Character
	tracker.Job = job
}

func UnassignTrackerJob(tracker *JobTracker) {
	if tracker.Job.Tile != nil && tracker.Job.Tile.Job == tracker.Job {
		tracker.Job.Tile.Job = nil
	}
	tracker.Job.Mob = nil
	for _, toil := range tracker.Job.Toils {
		toil.Started = false
	}
	tracker.Job = nil
}

func AdvanceTrackerJob(tracker *JobTracker) bool {
	tracker.Job.ToilIndex++
	if tracker.Job.ToilIndex >= uint8(len(tracker.Job.Toils)) {
		UnassignTrackerJob(tracker)
		return false
	}

	return true
}

func TimeoutTrackerJob(tracker *JobTracker) {
	job := tracker.Job
	job.TimeoutTicks = GameInstance.Ticks + 300
	UnassignTrackerJob(tracker)
	PushJobToController(GameInstance.JobController[tracker.Character.Faction], job)
}
