package gameplay

import (
	"math/rand"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinMap struct {
	Seed  int64
	Size  sdlutils.Vector3
	Tiles [][][]*KurinTile
	Items       []*KurinItem

	Random      *rand.Rand
	Pathfinding KurinPathfindingGrid
}

type KurinMapMarshal struct {
	Seed int64
	Size sdlutils.Vector3
}

func NewKurinMap(size sdlutils.Vector3) KurinMap {
	kmap := KurinMap{
		Seed:  0,
		Size:  size,
		Tiles: make([][][]*KurinTile, size.Base.X),
		Items:              []*KurinItem{},
	}
	kmap.Random = rand.New(rand.NewSource(kmap.Seed))
	for x := int32(0); x < kmap.Size.Base.X; x++ {
		kmap.Tiles[x] = make([][]*KurinTile, size.Base.Y)
		for y := int32(0); y < kmap.Size.Base.Y; y++ {
			kmap.Tiles[x][y] = make([]*KurinTile, size.Z)
			for z := uint8(0); z < kmap.Size.Z; z++ {
				tile := NewKurinTile("floor", sdlutils.Vector3{Base: sdl.Point{X: x, Y: y}, Z: z})
				if kmap.Random.Float32() < 0.1 {
					CreateKurinObject(tile, "grille")
				}
				kmap.Tiles[x][y][z] = tile
			}
		}
	}
	for i := 0; i < 3; i++ {
		item := NewKurinItemRandom("survivalknife", &kmap)
		kmap.Items = append(kmap.Items, item)
	}
	for i := 0; i < 3; i++ {
		item := NewKurinItemRandom("welder", &kmap)
		kmap.Items = append(kmap.Items, item)
	}
	kmap.Pathfinding = NewKurinPathfindingGrid(&kmap)

	return kmap
}

func MarshalKurinMap(kmap *KurinMap) (KurinMapMarshal, *error) {
	return KurinMapMarshal{
		Seed: kmap.Seed,
	}, nil
}

func GetTileAt(kmap *KurinMap, position sdlutils.Vector3) *KurinTile {
	if position.Base.X < 0 || position.Base.Y < 0 || position.Base.X >= kmap.Size.Base.X || position.Base.Y >= kmap.Size.Base.Y || position.Z >= kmap.Size.Z {
		return nil
	}

	return kmap.Tiles[position.Base.X][position.Base.Y][position.Z]
}

func CanEnterPosition(kmap *KurinMap, position sdlutils.Vector3) bool {
	tile := GetTileAt(kmap, position)
	return tile != nil && CanEnterKurinTile(tile)
}
