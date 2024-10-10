package gameplay

func NewJobDriverTemplateDestroyFloor() *JobDriverTemplate {
	template := NewJobDriverTemplate[interface{}]("destroy_floor")
	template.Initialize = func(job *JobDriver) {
		job.Toils = []*JobToil{}
		job.Toils = append(job.Toils, NewJobToil("goto", &JobToilGotoData{Target: job.Tile.Position}))
		job.Toils = append(job.Toils, NewJobToil("destroy_floor", nil))
	}

	return template
}
