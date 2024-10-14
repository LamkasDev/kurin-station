package gameplay

type JobDriverAttackData struct {
	Target *Mob
}

func NewJobDriverTemplateAttack() *JobDriverTemplate {
	template := NewJobDriverTemplate[*JobDriverAttackData]("attack")
	template.ReturnsOnFail = false
	template.Initialize = func(job *JobDriver) {
		data := job.Data.(*JobDriverAttackData)
		job.Toils = []*JobToil{
			NewJobToil("goto", &JobToilGotoData{Target: data.Target.Position}),
			NewJobToil("attack", &JobToilAttackData{Target: data.Target}),
		}
	}

	return template
}
