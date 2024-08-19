package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
	"github.com/veandco/go-sdl2/sdl"
)

func BuildLine(kmap *Map, start sdlutils.Vector3, direction common.Direction, length uint8, floor *string, wall *string) {
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

func BuildRoom(kmap *Map, rect sdl.Rect, floor string, wall string, blanks bool) {
	for x := rect.X; x < rect.X+rect.W; x++ {
		for y := rect.Y; y < rect.Y+rect.H; y++ {
			pos := sdlutils.Vector3{Base: sdl.Point{X: x, Y: y}, Z: 0}
			tile := GetTileAt(kmap, pos)
			if tile != nil {
				tile.Type = floor
			} else {
				CreateTileRaw(kmap, pos, floor)
			}
		}
	}
	if blanks {
		CreateTileRaw(kmap, sdlutils.Vector3{Base: sdl.Point{X: rect.X, Y: rect.Y}}, "blank")
		CreateTileRaw(kmap, sdlutils.Vector3{Base: sdl.Point{X: rect.X + rect.W - 1, Y: rect.Y}}, "blank")
		CreateTileRaw(kmap, sdlutils.Vector3{Base: sdl.Point{X: rect.X + rect.W - 1, Y: rect.Y + rect.H - 1}}, "blank")
		CreateTileRaw(kmap, sdlutils.Vector3{Base: sdl.Point{X: rect.X, Y: rect.Y + rect.H - 1}}, "blank")
	}
	for x := rect.X + 1; x < rect.X+rect.W; x++ {
		tile := GetTileAt(kmap, sdlutils.Vector3{Base: sdl.Point{X: x, Y: rect.Y}, Z: 0})
		ReplaceObjectRaw(kmap, tile, wall)
	}
	for y := rect.Y + 1; y < rect.Y+rect.H; y++ {
		tile := GetTileAt(kmap, sdlutils.Vector3{Base: sdl.Point{X: rect.X + rect.W - 1, Y: y}, Z: 0})
		ReplaceObjectRaw(kmap, tile, wall)
	}
	for x := rect.X; x < rect.X+rect.W-1; x++ {
		tile := GetTileAt(kmap, sdlutils.Vector3{Base: sdl.Point{X: x, Y: rect.Y + rect.H - 1}, Z: 0})
		ReplaceObjectRaw(kmap, tile, wall)
	}
	for y := rect.Y; y < rect.Y+rect.H-1; y++ {
		tile := GetTileAt(kmap, sdlutils.Vector3{Base: sdl.Point{X: rect.X, Y: y}, Z: 0})
		ReplaceObjectRaw(kmap, tile, wall)
	}
}

func BuildSmallThruster(kmap *Map, position sdl.Point, thrusterType string) {
	position1 := sdlutils.Vector3{Base: sdl.Point{X: position.X, Y: position.Y}, Z: 0}
	CreateTileRaw(kmap, position1, "blank")
	thruster := CreateObjectRaw(kmap, GetTileAt(kmap, position1), thrusterType)
	thruster.Direction = common.DirectionWest
}

func BuildBigThruster(kmap *Map, position sdl.Point, lattice string) {
	position1 := sdlutils.Vector3{Base: sdl.Point{X: position.X, Y: position.Y}, Z: 0}
	position2 := sdlutils.Vector3{Base: sdl.Point{X: position.X + 1, Y: position.Y}, Z: 0}
	position3 := sdlutils.Vector3{Base: sdl.Point{X: position.X + 2, Y: position.Y}, Z: 0}
	CreateTileRaw(kmap, position1, "blank")
	CreateTileRaw(kmap, position2, "blank")
	CreateTileRaw(kmap, position3, "blank")
	CreateObjectRaw(kmap, GetTileAt(kmap, position1), lattice)
	CreateObjectRaw(kmap, GetTileAt(kmap, position3), lattice)
	thruster := CreateObjectRaw(kmap, GetTileAt(kmap, position1), "big_thruster")
	thruster.Direction = common.DirectionWest
}
