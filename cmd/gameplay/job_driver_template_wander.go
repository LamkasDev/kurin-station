package gameplay

func NewJobDriverTemplateWander() *JobDriverTemplate {
	template := NewJobDriverTemplate[interface{}]("wander")
	template.Initialize = func(job *JobDriver) {
		job.Toils = []*JobToil{
			NewJobToil("wander", nil),
		}
	}

	return template
}
