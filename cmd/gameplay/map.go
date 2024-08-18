package gameplay

import (
	"math/rand"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/exp/slices"
)

type KurinMap struct {
	Seed    int64
	Size    sdlutils.Vector3
	Tiles   [][][]*KurinTile
	Objects []*KurinObject
	Items   []*KurinItem

	Random      *rand.Rand
	Pathfinding KurinPathfindingGrid
}

func NewKurinMap(size sdlutils.Vector3) KurinMap {
	kmap := KurinMap{
		Seed:  0,
		Size:  size,
		Tiles: make([][][]*KurinTile, size.Base.X),
		Items: []*KurinItem{},
	}
	kmap.Random = rand.New(rand.NewSource(kmap.Seed))
	for x := range kmap.Size.Base.X {
		kmap.Tiles[x] = make([][]*KurinTile, size.Base.Y)
		for y := range kmap.Size.Base.Y {
			kmap.Tiles[x][y] = make([]*KurinTile, size.Z)
		}
	}
	kmap.Pathfinding = NewKurinPathfindingGrid(&kmap)

	return kmap
}

func PopulateKurinMap(kmap *KurinMap) {
	size := int32(8)
	start := int32((kmap.Size.Base.X / 2) - (size / 2))

	for x := start; x < start+size; x++ {
		for y := start; y < start+size; y++ {
			CreateKurinTileRaw(kmap, sdlutils.Vector3{Base: sdl.Point{X: x, Y: y}, Z: 0}, "floor")
		}
	}
	CreateKurinTileRaw(kmap, sdlutils.Vector3{Base: sdl.Point{X: start, Y: start}}, "blank")
	CreateKurinTileRaw(kmap, sdlutils.Vector3{Base: sdl.Point{X: start + size - 1, Y: start}}, "blank")
	CreateKurinTileRaw(kmap, sdlutils.Vector3{Base: sdl.Point{X: start + size - 1, Y: start + size - 1}}, "blank")
	CreateKurinTileRaw(kmap, sdlutils.Vector3{Base: sdl.Point{X: start, Y: start + size - 1}}, "blank")

	CreateKurinObjectRaw(kmap, GetKurinTileAt(kmap, sdlutils.Vector3{Base: sdl.Point{X: start + 1, Y: start + 1}, Z: 0}), "lathe")

	CreateKurinObjectRaw(kmap, GetKurinTileAt(kmap, sdlutils.Vector3{Base: sdl.Point{X: start + 3, Y: start + 3}, Z: 0}), "pod")
	for x := int32(1); x < size; x++ {
		CreateKurinObjectRaw(kmap, GetKurinTileAt(kmap, sdlutils.Vector3{Base: sdl.Point{X: start + x, Y: start}, Z: 0}), "shuttle_wall")
	}
	for y := int32(1); y < size; y++ {
		CreateKurinObjectRaw(kmap, GetKurinTileAt(kmap, sdlutils.Vector3{Base: sdl.Point{X: start + size - 1, Y: start + y}, Z: 0}), "shuttle_wall")
	}
	for x := range size - 1 {
		CreateKurinObjectRaw(kmap, GetKurinTileAt(kmap, sdlutils.Vector3{Base: sdl.Point{X: start + x, Y: start + size - 1}, Z: 0}), "shuttle_wall")
	}
	for y := range size - 1 {
		CreateKurinObjectRaw(kmap, GetKurinTileAt(kmap, sdlutils.Vector3{Base: sdl.Point{X: start, Y: start + y}, Z: 0}), "shuttle_wall")
	}

	thrusterLeftPos := sdlutils.Vector3{Base: sdl.Point{X: start + 2, Y: start - 1}, Z: 0}
	thrusterLeftPos2 := sdlutils.Vector3{Base: sdl.Point{X: start + 3, Y: start - 1}, Z: 0}
	thrusterLeftPos3 := sdlutils.Vector3{Base: sdl.Point{X: start + 4, Y: start - 1}, Z: 0}
	CreateKurinTileRaw(kmap, thrusterLeftPos, "blank")
	CreateKurinTileRaw(kmap, thrusterLeftPos2, "blank")
	CreateKurinTileRaw(kmap, thrusterLeftPos3, "blank")
	CreateKurinObjectRaw(kmap, GetKurinTileAt(kmap, thrusterLeftPos), "lattice_l")
	CreateKurinObjectRaw(kmap, GetKurinTileAt(kmap, thrusterLeftPos3), "lattice_l")
	thrusterLeft := CreateKurinObjectRaw(kmap, GetKurinTileAt(kmap, thrusterLeftPos), "big_thruster")
	thrusterLeft.Direction = common.KurinDirectionWest

	thrusterRightPos := sdlutils.Vector3{Base: sdl.Point{X: start + 2, Y: start + size}, Z: 0}
	thrusterRightPos2 := sdlutils.Vector3{Base: sdl.Point{X: start + 3, Y: start + size}, Z: 0}
	thrusterRightPos3 := sdlutils.Vector3{Base: sdl.Point{X: start + 4, Y: start + size}, Z: 0}
	CreateKurinTileRaw(kmap, thrusterRightPos, "blank")
	CreateKurinTileRaw(kmap, thrusterRightPos2, "blank")
	CreateKurinTileRaw(kmap, thrusterRightPos3, "blank")
	CreateKurinObjectRaw(kmap, GetKurinTileAt(kmap, thrusterRightPos), "lattice_r")
	CreateKurinObjectRaw(kmap, GetKurinTileAt(kmap, thrusterRightPos3), "lattice_r")
	thrusterRight := CreateKurinObjectRaw(kmap, GetKurinTileAt(kmap, thrusterRightPos), "big_thruster")
	thrusterRight.Direction = common.KurinDirectionWest
}

func GetRandomMapPosition(kmap *KurinMap) sdlutils.Vector3 {
	return sdlutils.Vector3{
		Base: sdl.Point{
			X: int32(rand.Float32() * float32(GameInstance.Map.Size.Base.X)),
			Y: int32(rand.Float32() * float32(GameInstance.Map.Size.Base.Y)),
		},
		Z: 0,
	}
}

func IsMapPositionOutOfBounds(kmap *KurinMap, position sdlutils.Vector3) bool {
	if position.Base.X < 0 || position.Base.Y < 0 || position.Base.X >= kmap.Size.Base.X || position.Base.Y >= kmap.Size.Base.Y || position.Z >= kmap.Size.Z {
		return true
	}

	return false
}

func CanEnterMapPosition(kmap *KurinMap, position sdlutils.Vector3) bool {
	tile := GetKurinTileAt(kmap, position)
	return tile != nil && CanEnterKurinTile(tile)
}

func CreateKurinTileRaw(kmap *KurinMap, position sdlutils.Vector3, tileType string) *KurinTile {
	tile := NewKurinTile(tileType, position)
	kmap.Tiles[position.Base.X][position.Base.Y][position.Z] = tile

	return tile
}

func CreateKurinObjectRaw(kmap *KurinMap, tile *KurinTile, objectType string) *KurinObject {
	obj := NewKurinObject(tile, objectType)
	size := GetKurinObjectSize(obj)
	for x := tile.Position.Base.X; x < tile.Position.Base.X+size.X; x++ {
		for y := tile.Position.Base.Y; y < tile.Position.Base.Y+size.Y; y++ {
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
