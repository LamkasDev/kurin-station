package gameplay

func NewJobDriverTemplatePanic() *JobDriverTemplate {
	template := NewJobDriverTemplate[interface{}]("panic")
	template.ReturnsOnFail = false
	template.Initialize = func(job *JobDriver) {
		job.Toils = []*JobToil{
			NewJobToil("panic", nil),
		}
	}

	return template
}
