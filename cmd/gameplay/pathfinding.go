package gameplay

import (
	"container/heap"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/exp/slices"
)

type KurinPathfindingNode struct {
	Position sdlutils.Vector3
	Parent   *KurinPathfindingNode
	Cost     float32
	Rank     float32
	Index    int
}

func NewKurinPathfindingNode(position sdlutils.Vector3) *KurinPathfindingNode {
	return &KurinPathfindingNode{
		Position: position,
	}
}

func GetKurinPathfindingNodeNeighbourCost(from *KurinPathfindingNode, to *KurinPathfindingNode) float32 {
	return 1
}

func GetKurinPathfindingNodeEstimatedCost(from *KurinPathfindingNode, to *KurinPathfindingNode) float32 {
	return float32(sdlutils.GetDistance(from.Position.Base, to.Position.Base))
}

type KurinPathfindingGrid struct {
	Size  sdl.Point
	Nodes [][]*KurinPathfindingNode
}

func NewKurinPathfindingGrid(kmap *KurinMap) KurinPathfindingGrid {
	grid := KurinPathfindingGrid{
		Size:  kmap.Size.Base,
		Nodes: make([][]*KurinPathfindingNode, kmap.Size.Base.X),
	}
	for x := range kmap.Size.Base.X {
		grid.Nodes[x] = make([]*KurinPathfindingNode, kmap.Size.Base.Y)
		for y := range kmap.Size.Base.Y {
			grid.Nodes[x][y] = NewKurinPathfindingNode(sdlutils.Vector3{Base: sdl.Point{X: x, Y: y}, Z: 0})
		}
	}

	return grid
}

func GetNodeAt(grid *KurinPathfindingGrid, position sdl.Point) *KurinPathfindingNode {
	if !CanEnterMapPosition(&GameInstance.Map, sdlutils.Vector3{Base: position, Z: 0}) {
		return nil
	}

	return grid.Nodes[position.X][position.Y]
}

type KurinPath struct {
	Cost  float32
	Nodes []*KurinPathfindingNode

	Index int
	Ticks uint32
}

func FindKurinPathAdjacent(grid *KurinPathfindingGrid, from sdlutils.Vector3, to sdlutils.Vector3) *KurinPath {
	fromNode := grid.Nodes[from.Base.X][from.Base.Y]
	toNode := grid.Nodes[to.Base.X][to.Base.Y]
	neighbours := GetKurinPathfindingNodeNeighbours(grid, toNode, false)
	slices.SortFunc(neighbours, func(a, b *KurinPathfindingNode) int {
		return int(GetKurinPathfindingNodeEstimatedCost(fromNode, a) - GetKurinPathfindingNodeEstimatedCost(fromNode, b))
	})
	if len(neighbours) == 0 {
		return nil
	}

	return FindKurinPath(grid, from, neighbours[0].Position)
}

func FindKurinPath(grid *KurinPathfindingGrid, from sdlutils.Vector3, to sdlutils.Vector3) *KurinPath {
	fromNode := grid.Nodes[from.Base.X][from.Base.Y]
	fromNode.Parent = nil
	toNode := grid.Nodes[to.Base.X][to.Base.Y]
	toNode.Parent = nil

	closed := map[*KurinPathfindingNode]bool{fromNode: true}
	open := &KurinPathfindingQueue{}
	heap.Init(open)
	heap.Push(open, fromNode)
	for open.Len() > 0 {
		currentNode := heap.Pop(open).(*KurinPathfindingNode)

		if currentNode == toNode {
			path := KurinPath{
				Cost:  currentNode.Cost,
				Nodes: []*KurinPathfindingNode{},
			}
			pathNode := currentNode
			for pathNode != nil {
				path.Nodes = append(path.Nodes, pathNode)
				pathNode = pathNode.Parent
			}
			slices.Reverse(path.Nodes)

			return &path
		}

		for _, neighbourNode := range GetKurinPathfindingNodeNeighbours(grid, currentNode, true) {
			isClosed, _ := closed[neighbourNode]
			if !isClosed {
				closed[neighbourNode] = true
			}

			cost := currentNode.Cost + GetKurinPathfindingNodeNeighbourCost(currentNode, neighbourNode)
			if !isClosed || cost < neighbourNode.Cost {
				neighbourNode.Cost = cost
				neighbourNode.Rank = cost + GetKurinPathfindingNodeEstimatedCost(neighbourNode, toNode)
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
