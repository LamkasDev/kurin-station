package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

type PathfindingGrid struct {
	Size  sdl.Point
	Nodes [][]*PathfindingNode
}

func NewPathfindingGrid(kmap *Map) PathfindingGrid {
	grid := PathfindingGrid{
		Size:  kmap.Size.Base,
		Nodes: make([][]*PathfindingNode, kmap.Size.Base.X),
	}
	for x := range kmap.Size.Base.X {
		grid.Nodes[x] = make([]*PathfindingNode, kmap.Size.Base.Y)
		for y := range kmap.Size.Base.Y {
			grid.Nodes[x][y] = NewPathfindingNode(sdlutils.Vector3{Base: sdl.Point{X: x, Y: y}, Z: 0})
		}
	}

	return grid
}

func GetNodeAt(grid *PathfindingGrid, position sdl.Point) *PathfindingNode {
	if CanEnterMapPosition(&GameInstance.Map, sdlutils.Vector3{Base: position, Z: 0}) == EnteranceStatusNo {
		return nil
	}

	return grid.Nodes[position.X][position.Y]
}
