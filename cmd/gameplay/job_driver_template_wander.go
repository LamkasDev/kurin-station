package gameplay

func NewJobDriverTemplateWander() *JobDriverTemplate {
	template := NewJobDriverTemplate[interface{}]("wander")
	template.ReturnsOnFail = false
	template.Initialize = func(job *JobDriver) {
		job.Toils = []*JobToil{
			NewJobToil("wander", nil),
		}
	}

	return template
}
