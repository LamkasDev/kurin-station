package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
	"github.com/veandco/go-sdl2/sdl"
)

func BuildLine(kmap *Map, start sdlutils.Vector3, direction common.Direction, length uint8, floor *uint8, wall *string) {
	pos := start
	for range length {
		if floor != nil {
			CreateTileRaw(kmap, pos, *floor)
		}
		if wall != nil {
			ReplaceObjectRaw(kmap, GetTileAt(kmap, pos), *wall)
		}
		pos.Base = common.GetPositionInDirection(pos.Base, direction)
	}
}

func PlaceFloors(kmap *Map, rect sdlutils.Rect3, floor uint8) {
	for x := rect.Base.X; x < rect.Base.X+rect.Base.W; x++ {
		for y := rect.Base.Y; y < rect.Base.Y+rect.Base.H; y++ {
			pos := sdlutils.Vector3{Base: sdl.Point{X: x, Y: y}, Z: rect.Z}
			tile := GetTileAt(kmap, pos)
			if tile != nil {
				tile.Type = floor
			} else {
				CreateTileRaw(kmap, pos, floor)
			}
		}
	}
}

func BuildRoom(kmap *Map, rect sdlutils.Rect3, floor uint8, wall string, blanks bool) {
	PlaceFloors(kmap, rect, floor)
	if blanks {
		CreateTileRaw(kmap, sdlutils.Vector3{Base: sdl.Point{X: rect.Base.X, Y: rect.Base.Y}, Z: rect.Z}, TileIdBlank)
		CreateTileRaw(kmap, sdlutils.Vector3{Base: sdl.Point{X: rect.Base.X + rect.Base.W - 1, Y: rect.Base.Y}, Z: rect.Z}, TileIdBlank)
		CreateTileRaw(kmap, sdlutils.Vector3{Base: sdl.Point{X: rect.Base.X + rect.Base.W - 1, Y: rect.Base.Y + rect.Base.H - 1}, Z: rect.Z}, TileIdBlank)
		CreateTileRaw(kmap, sdlutils.Vector3{Base: sdl.Point{X: rect.Base.X, Y: rect.Base.Y + rect.Base.H - 1}, Z: rect.Z}, TileIdBlank)
	}
	for x := rect.Base.X + 1; x < rect.Base.X+rect.Base.W; x++ {
		tile := GetTileAt(kmap, sdlutils.Vector3{Base: sdl.Point{X: x, Y: rect.Base.Y}, Z: rect.Z})
		ReplaceObjectRaw(kmap, tile, wall)
	}
	for y := rect.Base.Y + 1; y < rect.Base.Y+rect.Base.H; y++ {
		tile := GetTileAt(kmap, sdlutils.Vector3{Base: sdl.Point{X: rect.Base.X + rect.Base.W - 1, Y: y}, Z: rect.Z})
		ReplaceObjectRaw(kmap, tile, wall)
	}
	for x := rect.Base.X; x < rect.Base.X+rect.Base.W-1; x++ {
		tile := GetTileAt(kmap, sdlutils.Vector3{Base: sdl.Point{X: x, Y: rect.Base.Y + rect.Base.H - 1}, Z: rect.Z})
		ReplaceObjectRaw(kmap, tile, wall)
	}
	for y := rect.Base.Y; y < rect.Base.Y+rect.Base.H-1; y++ {
		tile := GetTileAt(kmap, sdlutils.Vector3{Base: sdl.Point{X: rect.Base.X, Y: y}, Z: rect.Z})
		ReplaceObjectRaw(kmap, tile, wall)
	}
}

func BuildSmallThruster(kmap *Map, position sdlutils.Vector3, thrusterType string) {
	CreateTileRaw(kmap, position, TileIdBlank)
	thruster := CreateObjectRaw(kmap, GetTileAt(kmap, position), thrusterType)
	thruster.Direction = common.DirectionWest
}

func BuildBigThruster(kmap *Map, position1 sdlutils.Vector3, lattice string) {
	position2 := sdlutils.Vector3{Base: sdl.Point{X: position1.Base.X + 1, Y: position1.Base.Y}, Z: position1.Z}
	position3 := sdlutils.Vector3{Base: sdl.Point{X: position1.Base.X + 2, Y: position1.Base.Y}, Z: position1.Z}
	CreateTileRaw(kmap, position1, TileIdBlank)
	CreateTileRaw(kmap, position2, TileIdBlank)
	CreateTileRaw(kmap, position3, TileIdBlank)
	CreateObjectRaw(kmap, GetTileAt(kmap, position1), lattice)
	CreateObjectRaw(kmap, GetTileAt(kmap, position3), lattice)
	thruster := CreateObjectRaw(kmap, GetTileAt(kmap, position1), "big_thruster")
	thruster.Direction = common.DirectionWest
}
