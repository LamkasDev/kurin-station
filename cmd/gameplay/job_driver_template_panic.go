package gameplay

func NewJobDriverTemplatePanic() *JobDriverTemplate {
	template := NewJobDriverTemplate[interface{}]("panic")
	template.Initialize = func(job *JobDriver, data interface{}) {
		job.Toils = []*JobToil{
			NewJobToil("panic", nil),
		}
	}

	return template
}
