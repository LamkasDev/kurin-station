package gameplay

type KurinJobDriver struct {
	Type      string
	Tile      *KurinTile
	Character *KurinCharacter
	Data      interface{}
	Toils     []*KurinJobToil
}
