package gameplay

type KurinJobDriver struct {
	Tile  *KurinTile
	Ticks int32

	Assign  KurinJobDriverAssign
	Process KurinJobDriverProcess
	Data    interface{}
}

type KurinJobDriverAssign func(driver *KurinJobDriver, game *KurinGame, character *KurinCharacter)

type KurinJobDriverProcess func(driver *KurinJobDriver, game *KurinGame, character *KurinCharacter) bool
