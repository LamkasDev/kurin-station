package gameplay

func NewThinktreeNodeTemplatePanic() *ThinktreeNodeTemplate {
	template := NewThinktreeNodeTemplate[interface{}]("panic")
	template.Process = func(mob *Mob, node *ThinktreeNode) bool {
		if IsMobAggro(mob) {
			if mob.JobTracker.Job != nil && mob.JobTracker.Job.Type != "panic" {
				UnassignTrackerJob(mob.JobTracker)
			}
			if mob.JobTracker.Job == nil {
				job := NewJobDriver("panic", nil)
				job.Template.Initialize(job)
				AssignTrackerJob(mob.JobTracker, job)
			}

			return true
		}

		return false
	}

	return template
}
