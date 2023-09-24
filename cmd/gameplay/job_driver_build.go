package gameplay

type KurinJobDriverBuildData struct {
	Prefab *KurinObject
	Path   *KurinPath
}

func NewKurinJobDriverBuild(tile *KurinTile, prefab *KurinObject) *KurinJobDriver {
	return &KurinJobDriver{
		Tile: tile,
		Data: KurinJobDriverBuildData{
			Prefab: prefab,
		},
		Assign:  AssignKurinJobDriverBuild,
		Process: ProcessKurinJobDriverBuild,
	}
}

func AssignKurinJobDriverBuild(driver *KurinJobDriver, game *KurinGame, character *KurinCharacter) {
	data := driver.Data.(KurinJobDriverBuildData)
	data.Path = FindKurinPathAdjacent(&game.Map.Pathfinding, character.Position, driver.Tile.Position)
	driver.Data = data
}

func ProcessKurinJobDriverBuild(driver *KurinJobDriver, game *KurinGame, character *KurinCharacter) bool {
	data := driver.Data.(KurinJobDriverBuildData)
	if data.Path != nil {
		if FollowKurinPath(character, data.Path) {
			data.Path = nil
		}
	} else {
		CreateKurinObject(driver.Tile, data.Prefab.Type)
		return true
	}
	driver.Data = data

	return false
}
