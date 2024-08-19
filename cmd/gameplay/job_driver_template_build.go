package gameplay

type JobDriverBuildData struct {
	ObjectType string
}

func NewJobDriverTemplateBuild() *JobDriverTemplate {
	template := NewJobDriverTemplate[*JobDriverBuildData]("build")
	template.Initialize = func(job *JobDriver, data interface{}) {
		buildData := data.(*JobDriverBuildData)
		job.Toils = []*JobToil{}
		objectTemplate := ObjectContainer[buildData.ObjectType]
		for _, requirement := range objectTemplate.Requirements {
			job.Toils = append(job.Toils, NewJobToil("pickup", &JobToilPickupData{ItemType: requirement.Type, ItemCount: requirement.Count}))
		}
		job.Toils = append(job.Toils, NewJobToil("goto", &JobToilGotoData{Target: job.Tile.Position}))
		job.Toils = append(job.Toils, NewJobToil("build", &JobToilBuildData{ObjectType: buildData.ObjectType}))
		job.Data = data
	}

	return template
}
