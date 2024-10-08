package gameplay

func NewJobDriverTemplateDestroy() *JobDriverTemplate {
	template := NewJobDriverTemplate[interface{}]("destroy")
	template.Initialize = func(job *JobDriver) {
		job.Toils = []*JobToil{}
		job.Toils = append(job.Toils, NewJobToil("goto", &JobToilGotoData{Target: job.Tile.Position}))
		job.Toils = append(job.Toils, NewJobToil("destroy", nil))
	}

	return template
}
