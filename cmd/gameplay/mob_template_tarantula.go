package gameplay

func NewMobTemplateTarantula() *MobTemplate {
	template := NewMobTemplateRaw[interface{}]("tarantula")
	template.Process = ProcessTarantula

	return template
}

func ProcessTarantula(mob *Mob) {
	ProcessMob(mob)
	if mob.JobTracker.Job == nil {
		job := NewJobDriver("wander", nil)
		job.Template.Initialize(job)
		AssignTrackerJob(mob.JobTracker, job)
	}
}
