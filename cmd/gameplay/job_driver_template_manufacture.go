package gameplay

func NewJobDriverTemplateManufacture() *JobDriverTemplate {
	template := NewJobDriverTemplate[interface{}]("manufacture")
	template.Initialize = func(job *JobDriver, data interface{}) {
		job.Toils = []*JobToil{
			NewJobToil("goto", &JobToilGotoData{Target: job.Tile.Position}),
			NewJobToil("manufacture", nil),
		}
	}

	return template
}
