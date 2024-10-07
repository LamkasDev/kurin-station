package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

type PathfindingGrid struct {
	Size  sdlutils.Vector3
	Nodes [][][]*PathfindingNode
}

func NewPathfindingGrid(kmap *Map) PathfindingGrid {
	grid := PathfindingGrid{
		Size:  kmap.Size,
		Nodes: make([][][]*PathfindingNode, kmap.Size.Base.X),
	}
	for x := range kmap.Size.Base.X {
		grid.Nodes[x] = make([][]*PathfindingNode, kmap.Size.Base.Y)
		for y := range kmap.Size.Base.Y {
			grid.Nodes[x][y] = make([]*PathfindingNode, kmap.Size.Z)
			for z := range kmap.Size.Z {
				grid.Nodes[x][y][z] = NewPathfindingNode(sdlutils.Vector3{Base: sdl.Point{X: x, Y: y}, Z: z})
			}
		}
	}

	return grid
}

func GetNodeAt(grid *PathfindingGrid, position sdlutils.Vector3) *PathfindingNode {
	if CanEnterMapPosition(&GameInstance.Map, position) == EnteranceStatusNo {
		return nil
	}

	return grid.Nodes[position.Base.X][position.Base.Y][position.Z]
}
