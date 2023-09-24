package gameplay

import "github.com/veandco/go-sdl2/sdl"

func GetKurinPathfindingNodeNeighbours(grid *KurinPathfindingGrid, node *KurinPathfindingNode, walk bool) []*KurinPathfindingNode {
	neighbours := []*KurinPathfindingNode{}
	north := GetNodeAt(grid, sdl.Point{X: node.Tile.Position.Base.X, Y: node.Tile.Position.Base.Y - 1})
	if north != nil {
		neighbours = append(neighbours, north)
	}
	east := GetNodeAt(grid, sdl.Point{X: node.Tile.Position.Base.X + 1, Y: node.Tile.Position.Base.Y})
	if east != nil {
		neighbours = append(neighbours, east)
	}
	south := GetNodeAt(grid, sdl.Point{X: node.Tile.Position.Base.X, Y: node.Tile.Position.Base.Y + 1})
	if south != nil {
		neighbours = append(neighbours, south)
	}
	west := GetNodeAt(grid, sdl.Point{X: node.Tile.Position.Base.X - 1, Y: node.Tile.Position.Base.Y})
	if west != nil {
		neighbours = append(neighbours, west)
	}

	if (north != nil && east != nil) || !walk {
		if northEast := GetNodeAt(grid, sdl.Point{X: node.Tile.Position.Base.X + 1, Y: node.Tile.Position.Base.Y - 1}); northEast != nil {
			neighbours = append(neighbours, northEast)
		}
	}
	if (east != nil && south != nil) || !walk {
		if eastSouth := GetNodeAt(grid, sdl.Point{X: node.Tile.Position.Base.X + 1, Y: node.Tile.Position.Base.Y + 1}); eastSouth != nil {
			neighbours = append(neighbours, eastSouth)
		}
	}
	if (south != nil && west != nil) || !walk {
		if southWest := GetNodeAt(grid, sdl.Point{X: node.Tile.Position.Base.X - 1, Y: node.Tile.Position.Base.Y + 1}); southWest != nil {
			neighbours = append(neighbours, southWest)
		}
	}
	if (west != nil && north != nil) || !walk {
		if westNorth := GetNodeAt(grid, sdl.Point{X: node.Tile.Position.Base.X - 1, Y: node.Tile.Position.Base.Y - 1}); westNorth != nil {
			neighbours = append(neighbours, westNorth)
		}
	}

	return neighbours
}
