package gameplay

func NewMobTemplateCat() *MobTemplate {
	template := NewMobTemplateRaw[interface{}]("cat")
	template.Process = ProcessCat

	return template
}

func ProcessCat(mob *Mob) {
	ProcessMob(mob)
	if mob.JobTracker.Job == nil {
		job := NewJobDriver("wander", nil)
		job.Template.Initialize(job)
		AssignTrackerJob(mob.JobTracker, job)
	}
	if GameInstance.Ticks%900 == 0 {
		PlaySoundVolume(&GameInstance.SoundController, "meow", 0.1)
	}
}
