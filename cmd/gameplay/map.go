package gameplay

import (
	"math/rand"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/exp/slices"
)

type KurinMap struct {
	Seed  int64
	Size  sdlutils.Vector3
	Tiles [][][]*KurinTile
	Objects []*KurinObject
	Items       []*KurinItem

	Random      *rand.Rand
	Pathfinding KurinPathfindingGrid
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
				kmap.Tiles[x][y][z] = NewKurinTile("floor", sdlutils.Vector3{Base: sdl.Point{X: x, Y: y}, Z: z})
			}
		}
	}
	kmap.Pathfinding = NewKurinPathfindingGrid(&kmap)

	return kmap
}

func PopulateKurinMap(kmap *KurinMap) {
	CreateKurinObjectRaw(kmap, GetTileAt(kmap, sdlutils.Vector3{Base: sdlutils.DividePoint(kmap.Size.Base, 2), Z: 0}), "pod")
	for x := int32(0); x < kmap.Size.Base.X; x++ {
		for y := int32(0); y < kmap.Size.Base.Y; y++ {
			for z := uint8(0); z < kmap.Size.Z; z++ {
				tile := kmap.Tiles[x][y][z]
				if len(tile.Objects) > 0 {
					continue
				}
				if kmap.Random.Float32() < 0.1 {
					CreateKurinObjectRaw(kmap, tile, "grille")
				}
			}
		}
	}
	for i := 0; i < 3; i++ {
		item := NewKurinItemRandom("survivalknife", kmap)
		kmap.Items = append(kmap.Items, item)
	}
	for i := 0; i < 3; i++ {
		item := NewKurinItemRandom("welder", kmap)
		kmap.Items = append(kmap.Items, item)
	}
	for i := 0; i < 10; i++ {
		item := NewKurinItemRandom("credit", kmap)
		kmap.Items = append(kmap.Items, item)
	}
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

func CreateKurinObjectRaw(kmap *KurinMap, tile *KurinTile, objectType string) *KurinObject {
	obj := NewKurinObject(tile, objectType)
	size := GetKurinObjectSize(obj)
	for x := tile.Position.Base.X; x < tile.Position.Base.X + size.X; x++ {
		for y := tile.Position.Base.Y; y < tile.Position.Base.Y + size.Y; y++ {
			currentTile := kmap.Tiles[x][y][tile.Position.Z]
			currentTile.Objects = append(currentTile.Objects, obj)
		}
	}
	kmap.Objects = append(kmap.Objects, obj)
	
	return obj
}

func DestroyKurinObjectRaw(kmap *KurinMap, obj *KurinObject) {
	i := slices.Index(obj.Tile.Objects, obj)
	if i >= 0 {
		obj.Tile.Objects = slices.Delete(obj.Tile.Objects, i, i+1)
	}
	i = slices.Index(kmap.Objects, obj)
	if i >= 0 {
		kmap.Objects = slices.Delete(kmap.Objects, i, i+1)
	}
}
