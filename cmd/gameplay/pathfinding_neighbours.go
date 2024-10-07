package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
)

func GetPathfindingNodeNeighbours(grid *PathfindingGrid, node *PathfindingNode, walk bool) []*PathfindingNode {
	neighbours := []*PathfindingNode{}
	north := GetNodeAt(grid, sdlutils.Vector3{Base: sdl.Point{X: node.Position.Base.X, Y: node.Position.Base.Y - 1}, Z: node.Position.Z})
	if north != nil {
		neighbours = append(neighbours, north)
	}
	east := GetNodeAt(grid, sdlutils.Vector3{Base: sdl.Point{X: node.Position.Base.X + 1, Y: node.Position.Base.Y}, Z: node.Position.Z})
	if east != nil {
		neighbours = append(neighbours, east)
	}
	south := GetNodeAt(grid, sdlutils.Vector3{Base: sdl.Point{X: node.Position.Base.X, Y: node.Position.Base.Y + 1}, Z: node.Position.Z})
	if south != nil {
		neighbours = append(neighbours, south)
	}
	west := GetNodeAt(grid, sdlutils.Vector3{Base: sdl.Point{X: node.Position.Base.X - 1, Y: node.Position.Base.Y}, Z: node.Position.Z})
	if west != nil {
		neighbours = append(neighbours, west)
	}

	if (north != nil && east != nil) || !walk {
		if northEast := GetNodeAt(grid, sdlutils.Vector3{Base: sdl.Point{X: node.Position.Base.X + 1, Y: node.Position.Base.Y - 1}, Z: node.Position.Z}); northEast != nil {
			neighbours = append(neighbours, northEast)
		}
	}
	if (east != nil && south != nil) || !walk {
		if eastSouth := GetNodeAt(grid, sdlutils.Vector3{Base: sdl.Point{X: node.Position.Base.X + 1, Y: node.Position.Base.Y + 1}, Z: node.Position.Z}); eastSouth != nil {
			neighbours = append(neighbours, eastSouth)
		}
	}
	if (south != nil && west != nil) || !walk {
		if southWest := GetNodeAt(grid, sdlutils.Vector3{Base: sdl.Point{X: node.Position.Base.X - 1, Y: node.Position.Base.Y + 1}, Z: node.Position.Z}); southWest != nil {
			neighbours = append(neighbours, southWest)
		}
	}
	if (west != nil && north != nil) || !walk {
		if westNorth := GetNodeAt(grid, sdlutils.Vector3{Base: sdl.Point{X: node.Position.Base.X - 1, Y: node.Position.Base.Y - 1}, Z: node.Position.Z}); westNorth != nil {
			neighbours = append(neighbours, westNorth)
		}
	}

	return neighbours
}
