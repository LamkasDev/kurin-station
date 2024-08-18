package gameplay

type KurinJobDriverBuildData struct {
	Prefab string
}

func NewKurinJobDriverBuild() *KurinJobDriver {
	job := NewKurinJobDriverRaw[*KurinJobDriverBuildData]("build")
	job.Initialize = func(job *KurinJobDriver, data interface{}) {
		buildData := data.(*KurinJobDriverBuildData)
		job.Toils = []*KurinJobToil{
			NewKurinJobToilPickup("rod", 2),
			NewKurinJobToilGoto(job.Tile.Position),
			NewKurinJobToilBuild(buildData.Prefab),
		}
		job.Data = data
	}

	return job
}
