package gameplay

type JobDriverBuildData struct {
	ObjectType string
}

func NewJobDriverTemplateBuild() *JobDriverTemplate {
	template := NewJobDriverTemplate[*JobDriverBuildData]("build")
	template.Initialize = func(job *JobDriver) {
		data := job.Data.(*JobDriverBuildData)
		job.Toils = []*JobToil{}
		objectTemplate := ObjectContainer[data.ObjectType]
		for _, requirement := range objectTemplate.Requirements {
			pickupToil := NewJobToil(
				"pickup",
				&JobToilPickupData{
					ItemType:  requirement.Type,
					ItemCount: requirement.Count,
				},
			)
			job.Toils = append(job.Toils, pickupToil)
		}
		job.Toils = append(job.Toils, NewJobToil("goto", &JobToilGotoData{Target: job.Tile.Position}))
		job.Toils = append(job.Toils, NewJobToil("build", &JobToilBuildData{ObjectType: data.ObjectType}))
	}

	return template
}
