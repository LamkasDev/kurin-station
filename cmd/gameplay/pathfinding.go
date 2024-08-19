package gameplay

import (
	"container/heap"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"golang.org/x/exp/slices"
)

type Path struct {
	Cost  float32
	Nodes []*PathfindingNode
	Index int
}

func FindPathAdjacent(grid *PathfindingGrid, from sdlutils.Vector3, to sdlutils.Vector3) *Path {
	fromNode := grid.Nodes[from.Base.X][from.Base.Y]
	toNode := grid.Nodes[to.Base.X][to.Base.Y]
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
	fromNode := grid.Nodes[from.Base.X][from.Base.Y]
	fromNode.Parent = nil
	toNode := grid.Nodes[to.Base.X][to.Base.Y]
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
