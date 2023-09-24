package gameplay

import (
	"fmt"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/exp/slices"
)

var KurinTileSize = sdl.Point{
	X: 32,
	Y: 32,
}

type KurinTile struct {
	Type      string
	Position  sdlutils.Vector3
	Direction uint8

	Job     *KurinJobDriver
	Objects []*KurinObject
}

func NewKurinTile(ttype string, position sdlutils.Vector3) *KurinTile {
	return &KurinTile{
		Type:     ttype,
		Position: position,
		Job:      nil,
		Objects:  []*KurinObject{},
	}
}

func CreateKurinObject(tile *KurinTile, otype string) {
	tile.Objects = append(tile.Objects, NewKurinObject(otype))
}

func DestroyKurinObject(tile *KurinTile, obj *KurinObject) {
	i := slices.Index(tile.Objects, obj)
	if i >= 0 {
		tile.Objects = slices.Delete(tile.Objects, i, i+1)
	}
}

func CanEnterKurinTile(tile *KurinTile) bool {
	return len(tile.Objects) == 0
}

func GetKurinTileDescription(tile *KurinTile) string {
	return fmt.Sprintf("%s [%d_%d]", tile.Type, tile.Position.Base.X, tile.Position.Base.Y)
}
