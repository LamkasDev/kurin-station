package gameplay

type KurinJobDriver struct {
	Ticks   int32
	Assign  KurinJobDriverAssign
	Process KurinJobDriverProcess

	Tile *KurinTile
	Data interface{}
}

type KurinJobDriverAssign func(driver *KurinJobDriver, game *KurinGame, character *KurinCharacter)

type KurinJobDriverProcess func(driver *KurinJobDriver, game *KurinGame, character *KurinCharacter) bool
