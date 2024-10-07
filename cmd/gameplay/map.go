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
	BaseZ   uint8
	Tiles   [][][]*Tile
	Objects []*Object
	Items   []*Item

	Random      *rand.Rand
	Pathfinding PathfindingGrid
}

func NewMap(size sdlutils.Vector3, baseZ uint8) Map {
	kmap := Map{
		Seed:  0,
		Size:  size,
		BaseZ: baseZ,
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
	PlaceFloors(kmap, sdlutils.Rect3{Base: sdl.Rect{X: 0, Y: 0, W: kmap.Size.Base.X, H: kmap.Size.Base.Y}, Z: kmap.BaseZ - 1}, "asteroid")

	size := sdl.Point{X: 7, Y: 7}
	start := sdl.Point{X: 22, Y: 22}

	mainRect := sdlutils.Rect3{Base: sdl.Rect{X: start.X, Y: start.Y, W: size.X, H: size.Y}, Z: kmap.BaseZ}
	frontRect := sdlutils.Rect3{Base: sdl.Rect{X: start.X + size.X - 1, Y: start.Y + 1, W: 5, H: 5}, Z: kmap.BaseZ}
	backRect := sdlutils.Rect3{Base: sdl.Rect{X: start.X - 9, Y: start.Y + 1, W: 6, H: 5}, Z: kmap.BaseZ}
	// bottomRect := sdlutils.Rect3{Base: backRect.Base, Z: kmap.BaseZ - 1}
	BuildRoom(kmap, frontRect, "floor", "shuttle_wall", true)
	BuildRoom(kmap, mainRect, "floor", "shuttle_wall", true)
	BuildRoom(kmap, backRect, "floor", "shuttle_wall", true)
	// BuildRoom(kmap, bottomRect, "floor", "shuttle_wall", true)

	// shuttleWall := "shuttle_wall"
	// floor := "floor"
	catwalk := "catwalk"
	window := "window"

	BuildLine(kmap, sdlutils.Vector3{
		Base: sdl.Point{X: mainRect.Base.X - 1, Y: mainRect.Base.Y + 2},
		Z:    kmap.BaseZ,
	}, common.DirectionWest, 3, &catwalk, &window)
	BuildLine(kmap, sdlutils.Vector3{
		Base: sdl.Point{X: mainRect.Base.X - 1, Y: mainRect.Base.Y + 3},
		Z:    kmap.BaseZ,
	}, common.DirectionWest, 3, &catwalk, nil)
	BuildLine(kmap, sdlutils.Vector3{
		Base: sdl.Point{X: mainRect.Base.X - 1, Y: mainRect.Base.Y + 4},
		Z:    kmap.BaseZ,
	}, common.DirectionWest, 3, &catwalk, &window)

	ReplaceObjectRaw(kmap, GetTileAt(kmap, sdlutils.Vector3{
		Base: sdl.Point{X: frontRect.Base.X, Y: frontRect.Base.Y + 2},
		Z:    kmap.BaseZ,
	}), "airlock")
	ReplaceObjectRaw(kmap, GetTileAt(kmap, sdlutils.Vector3{
		Base: sdl.Point{X: mainRect.Base.X, Y: mainRect.Base.Y + 3},
		Z:    kmap.BaseZ,
	}), "airlock")
	ReplaceObjectRaw(kmap, GetTileAt(kmap, sdlutils.Vector3{
		Base: sdl.Point{X: backRect.Base.X + backRect.Base.W - 1, Y: backRect.Base.Y + 2},
		Z:    kmap.BaseZ,
	}), "airlock")
	teleporter := CreateObjectRaw(kmap, GetTileAt(kmap, sdlutils.Vector3{
		Base: sdl.Point{X: backRect.Base.X + 1, Y: backRect.Base.Y + 2},
		Z:    kmap.BaseZ,
	}), "teleporter")
	teleporter2 := CreateObjectRaw(kmap, GetTileAt(kmap, sdlutils.Vector3{
		Base: sdl.Point{X: backRect.Base.X + 1, Y: backRect.Base.Y + 2},
		Z:    kmap.BaseZ - 1,
	}), "teleporter")
	teleporter.Data.(*ObjectTeleporterData).Target = teleporter2.Tile.Position
	teleporter2.Data.(*ObjectTeleporterData).Target = teleporter.Tile.Position

	CreateObjectRaw(kmap, GetTileAt(kmap, sdlutils.Vector3{
		Base: sdl.Point{X: mainRect.Base.X + 1, Y: mainRect.Base.Y + 1},
		Z:    kmap.BaseZ,
	}), "lathe")
	CreateObjectRaw(kmap, GetTileAt(kmap, sdlutils.Vector3{
		Base: sdl.Point{X: mainRect.Base.X + 3, Y: mainRect.Base.Y + 1},
		Z:    kmap.BaseZ,
	}), "console")
	CreateObjectRaw(kmap, GetTileAt(kmap, sdlutils.Vector3{
		Base: sdl.Point{X: mainRect.Base.X + 4, Y: mainRect.Base.Y + 1},
		Z:    kmap.BaseZ,
	}), "telepad")

	BuildBigThruster(kmap, sdlutils.Vector3{Base: sdl.Point{X: mainRect.Base.X + 2, Y: mainRect.Base.Y - 1}, Z: kmap.BaseZ}, "lattice_l")
	BuildBigThruster(kmap, sdlutils.Vector3{Base: sdl.Point{X: mainRect.Base.X + 2, Y: mainRect.Base.Y + mainRect.Base.H}, Z: kmap.BaseZ}, "lattice_r")

	BuildSmallThruster(kmap, sdlutils.Vector3{Base: sdl.Point{X: backRect.Base.X + 2, Y: backRect.Base.Y - 1}, Z: kmap.BaseZ}, "small_thruster_l")
	BuildSmallThruster(kmap, sdlutils.Vector3{Base: sdl.Point{X: backRect.Base.X + 2, Y: backRect.Base.Y + backRect.Base.H}, Z: kmap.BaseZ}, "small_thruster_r")
}

func GetRandomMapPosition(kmap *Map, z uint8) sdlutils.Vector3 {
	return sdlutils.Vector3{
		Base: sdl.Point{
			X: int32(rand.Float32() * float32(GameInstance.Map.Size.Base.X)),
			Y: int32(rand.Float32() * float32(GameInstance.Map.Size.Base.Y)),
		},
		Z: kmap.BaseZ,
	}
}

func GetRandomMapZ(kmap *Map) uint8 {
	return uint8(rand.Float32() * float32(GameInstance.Map.Size.Z))
}

func IsMapPositionOutOfBounds(kmap *Map, position sdlutils.Vector3) bool {
	if position.Base.X < 0 ||
		position.Base.Y < 0 ||
		position.Base.X >= kmap.Size.Base.X ||
		position.Base.Y >= kmap.Size.Base.Y ||
		position.Z >= kmap.Size.Z {
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
