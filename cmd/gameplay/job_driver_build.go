package gameplay

type KurinJobDriverBuildData struct {
	Prefab *KurinObject
}

func NewKurinJobDriverBuild(data KurinJobDriverBuildData) *KurinJobDriver {
	return &KurinJobDriver{
		Type: "build",
		Tile: data.Prefab.Tile,
		Toils: []*KurinJobToil{
			NewKurinJobToilGoto(data.Prefab.Tile.Position),
			NewKurinJobToilBuild(data.Prefab),
		},
		Data: data,
	}
}
