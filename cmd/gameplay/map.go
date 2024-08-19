package gameplay

import (
	"math/rand"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/exp/slices"
)

type Map struct {
	Seed    int64
	Size    sdlutils.Vector3
	Tiles   [][][]*Tile
	Objects []*Object
	Items   []*Item

	Random      *rand.Rand
	Pathfinding PathfindingGrid
}

func NewMap(size sdlutils.Vector3) Map {
	kmap := Map{
		Seed:  0,
		Size:  size,
		Tiles: make([][][]*Tile, size.Base.X),
		Items: []*Item{},
	}
	kmap.Random = rand.New(rand.NewSource(kmap.Seed))
	for x := range kmap.Size.Base.X {
		kmap.Tiles[x] = make([][]*Tile, size.Base.Y)
		for y := range kmap.Size.Base.Y {
			kmap.Tiles[x][y] = make([]*Tile, size.Z)
		}
	}
	kmap.Pathfinding = NewPathfindingGrid(&kmap)

	return kmap
}

func PopulateMap(kmap *Map) {
	size := sdl.Point{X: 7, Y: 7}
	start := sdl.Point{X: 32, Y: 22}

	mainRect := sdl.Rect{X: start.X, Y: start.Y, W: size.X, H: size.Y}
	frontRect := sdl.Rect{X: start.X + size.X - 1, Y: start.Y + 1, W: 5, H: 5}
	backRect := sdl.Rect{X: start.X - 9, Y: start.Y + 1, W: 6, H: 5}
	BuildRoom(kmap, frontRect, "floor", "shuttle_wall", true)
	BuildRoom(kmap, mainRect, "floor", "shuttle_wall", true)
	BuildRoom(kmap, backRect, "floor", "shuttle_wall", true)

	// shuttleWall := "shuttle_wall"
	// floor := "floor"
	catwalk := "catwalk"
	window := "window"

	BuildLine(kmap, sdlutils.Vector3{Base: sdl.Point{X: mainRect.X - 1, Y: mainRect.Y + 2}, Z: 0}, common.DirectionWest, 3, &catwalk, &window)
	BuildLine(kmap, sdlutils.Vector3{Base: sdl.Point{X: mainRect.X - 1, Y: mainRect.Y + 3}, Z: 0}, common.DirectionWest, 3, &catwalk, nil)
	BuildLine(kmap, sdlutils.Vector3{Base: sdl.Point{X: mainRect.X - 1, Y: mainRect.Y + 4}, Z: 0}, common.DirectionWest, 3, &catwalk, &window)

	ReplaceObjectRaw(kmap, GetTileAt(kmap, sdlutils.Vector3{Base: sdl.Point{X: frontRect.X, Y: frontRect.Y + 2}, Z: 0}), "airlock")
	ReplaceObjectRaw(kmap, GetTileAt(kmap, sdlutils.Vector3{Base: sdl.Point{X: mainRect.X, Y: mainRect.Y + 3}, Z: 0}), "airlock")
	ReplaceObjectRaw(kmap, GetTileAt(kmap, sdlutils.Vector3{Base: sdl.Point{X: backRect.X + backRect.W - 1, Y: backRect.Y + 2}, Z: 0}), "airlock")

	CreateObjectRaw(kmap, GetTileAt(kmap, sdlutils.Vector3{Base: sdl.Point{X: mainRect.X + 1, Y: mainRect.Y + 1}, Z: 0}), "lathe")
	CreateObjectRaw(kmap, GetTileAt(kmap, sdlutils.Vector3{Base: sdl.Point{X: mainRect.X + 3, Y: mainRect.Y + 3}, Z: 0}), "pod")

	BuildBigThruster(kmap, sdl.Point{X: mainRect.X + 2, Y: mainRect.Y - 1}, "lattice_l")
	BuildBigThruster(kmap, sdl.Point{X: mainRect.X + 2, Y: mainRect.Y + mainRect.H}, "lattice_r")

	BuildSmallThruster(kmap, sdl.Point{X: backRect.X + 2, Y: backRect.Y - 1}, "small_thruster_l")
	BuildSmallThruster(kmap, sdl.Point{X: backRect.X + 2, Y: backRect.Y + backRect.H}, "small_thruster_r")
}

func GetRandomMapPosition(kmap *Map) sdlutils.Vector3 {
	return sdlutils.Vector3{
		Base: sdl.Point{
			X: int32(rand.Float32() * float32(GameInstance.Map.Size.Base.X)),
			Y: int32(rand.Float32() * float32(GameInstance.Map.Size.Base.Y)),
		},
		Z: 0,
	}
}

func IsMapPositionOutOfBounds(kmap *Map, position sdlutils.Vector3) bool {
	if position.Base.X < 0 || position.Base.Y < 0 || position.Base.X >= kmap.Size.Base.X || position.Base.Y >= kmap.Size.Base.Y || position.Z >= kmap.Size.Z {
		return true
	}

	return false
}

func CreateTileRaw(kmap *Map, position sdlutils.Vector3, tileType string) *Tile {
	tile := NewTile(tileType, position)
	kmap.Tiles[position.Base.X][position.Base.Y][position.Z] = tile

	return tile
}

func DestroyTileRaw(kmap *Map, tile *Tile) {
	for _, object := range tile.Objects {
		DestroyObjectRaw(kmap, object)
	}

	kmap.Tiles[tile.Position.Base.X][tile.Position.Base.Y][tile.Position.Z] = nil
}

func CreateObjectRaw(kmap *Map, tile *Tile, objectType string) *Object {
	obj := NewObject(tile, objectType)
	size := GetObjectSize(obj)
	for x := tile.Position.Base.X; x < tile.Position.Base.X+size.X; x++ {
		for y := tile.Position.Base.Y; y < tile.Position.Base.Y+size.Y; y++ {
			currentTile := kmap.Tiles[x][y][tile.Position.Z]
			currentTile.Objects = append(currentTile.Objects, obj)
		}
	}
	kmap.Objects = append(kmap.Objects, obj)

	return obj
}

func DestroyObjectRaw(kmap *Map, obj *Object) {
	size := GetObjectSize(obj)
	for x := obj.Tile.Position.Base.X; x < obj.Tile.Position.Base.X+size.X; x++ {
		for y := obj.Tile.Position.Base.Y; y < obj.Tile.Position.Base.Y+size.Y; y++ {
			currentTile := kmap.Tiles[x][y][obj.Tile.Position.Z]
			i := slices.Index(currentTile.Objects, obj)
			if i >= 0 {
				currentTile.Objects = slices.Delete(currentTile.Objects, i, i+1)
			}
		}
	}
	i := slices.Index(kmap.Objects, obj)
	if i >= 0 {
		kmap.Objects = slices.Delete(kmap.Objects, i, i+1)
	}
}

func ReplaceObjectRaw(kmap *Map, tile *Tile, objectType string) *Object {
	for _, object := range tile.Objects {
		DestroyObjectRaw(kmap, object)
	}

	return CreateObjectRaw(kmap, tile, objectType)
}
