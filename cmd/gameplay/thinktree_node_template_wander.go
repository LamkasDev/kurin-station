package gameplay

func NewThinktreeNodeTemplateCreate() *ThinktreeNodeTemplate {
	template := NewThinktreeNodeTemplate[interface{}]("wander")
	template.Process = func(mob *Mob, node *ThinktreeNode) bool {
		if mob.JobTracker.Job == nil {
			job := NewJobDriver("wander", nil)
			job.Template.Initialize(job)
			AssignTrackerJob(mob.JobTracker, job)
		}

		return true
	}

	return template
}
