package gameplay

func NewMobTemplateCat() *MobTemplate {
	template := NewMobTemplateRaw[interface{}]("cat")
	template.Process = ProcessCat

	return template
}

func ProcessCat(mob *Mob) {
	ProcessMob(mob)
	if mob.JobTracker.Job == nil {
		job := NewJobDriver("panic", nil)
		job.Template.Initialize(job, nil)
		AssignTrackerJob(mob.JobTracker, job)
	}
	if GameInstance.Ticks%900 == 0 {
		PlaySoundVolume(&GameInstance.SoundController, "meow", 0.1)
	}
}
