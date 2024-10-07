package gameplay

import (
	"container/heap"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"golang.org/x/exp/slices"
	"robpike.io/filter"
)

type Path struct {
	Cost  float32
	Nodes []*PathfindingNode
	Index int
}

func JoinPaths(a *Path, b *Path) *Path {
	if a == nil || b == nil {
		return nil
	}

	return &Path{
		Cost:  a.Cost + b.Cost,
		Nodes: append(a.Nodes, b.Nodes...),
		Index: 0,
	}
}

func FindPathAdjacent(grid *PathfindingGrid, from sdlutils.Vector3, to sdlutils.Vector3) *Path {
	fromNode := grid.Nodes[from.Base.X][from.Base.Y][from.Z]
	toNode := grid.Nodes[to.Base.X][to.Base.Y][to.Z]
	neighbours := GetPathfindingNodeNeighbours(grid, toNode, false)
	slices.SortFunc(neighbours, func(a, b *PathfindingNode) int {
		return int(GetPathfindingNodeEstimatedCost(fromNode, a) - GetPathfindingNodeEstimatedCost(fromNode, b))
	})
	if len(neighbours) == 0 {
		return nil
	}

	return FindPath(grid, from, neighbours[0].Position)
}

func FindPath(grid *PathfindingGrid, from sdlutils.Vector3, to sdlutils.Vector3) *Path {
	if to.Z != from.Z {
		// TODO: finish this
		teleporters := filter.Choose(FindObjectsOfType(&GameInstance.Map, from.Z, "teleporter"), func(object *Object) bool {
			return object.Data.(*ObjectTeleporterData).Target.Z == to.Z
		}).([]*Object)
		if len(teleporters) == 0 {
			return nil
		}

		return JoinPaths(FindPath(grid, from, teleporters[0].Tile.Position), FindPath(grid, teleporters[0].Data.(*ObjectTeleporterData).Target, to))
	}

	fromNode := grid.Nodes[from.Base.X][from.Base.Y][from.Z]
	fromNode.Parent = nil
	toNode := grid.Nodes[to.Base.X][to.Base.Y][to.Z]
	toNode.Parent = nil

	closed := map[*PathfindingNode]bool{fromNode: true}
	open := &PathfindingQueue{}
	heap.Init(open)
	heap.Push(open, fromNode)
	for open.Len() > 0 {
		currentNode := heap.Pop(open).(*PathfindingNode)

		if currentNode == toNode {
			path := Path{
				Cost:  currentNode.Cost,
				Nodes: []*PathfindingNode{},
			}
			pathNode := currentNode
			for pathNode != nil {
				path.Nodes = append(path.Nodes, pathNode)
				pathNode = pathNode.Parent
			}
			slices.Reverse(path.Nodes)

			return &path
		}

		for _, neighbourNode := range GetPathfindingNodeNeighbours(grid, currentNode, true) {
			isClosed, _ := closed[neighbourNode]
			if !isClosed {
				closed[neighbourNode] = true
			}

			cost := currentNode.Cost + GetPathfindingNodeNeighbourCost(currentNode, neighbourNode)
			if !isClosed || cost < neighbourNode.Cost {
				neighbourNode.Cost = cost
				neighbourNode.Rank = cost + GetPathfindingNodeEstimatedCost(neighbourNode, toNode)
				neighbourNode.Parent = currentNode
				if isClosed {
					heap.Fix(open, neighbourNode.Index)
				} else {
					heap.Push(open, neighbourNode)
				}
			}
		}
	}

	return nil
}
