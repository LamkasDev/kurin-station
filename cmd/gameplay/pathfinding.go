package gameplay

import (
	"container/heap"

	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/exp/slices"
)

type KurinPathfindingNode struct {
	Tile   *KurinTile
	Parent *KurinPathfindingNode
	Cost   float32
	Rank   float32
	Index  int
}

func NewKurinPathfindingNode(tile *KurinTile) *KurinPathfindingNode {
	return &KurinPathfindingNode{
		Tile: tile,
	}
}

func GetKurinPathfindingNodeNeighbourCost(from *KurinPathfindingNode, to *KurinPathfindingNode) float32 {
	return 1
}

func GetKurinPathfindingNodeEstimatedCost(from *KurinPathfindingNode, to *KurinPathfindingNode) float32 {
	return float32(sdlutils.GetDistance(from.Tile.Position.Base, to.Tile.Position.Base))
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
	for x := int32(0); x < kmap.Size.Base.X; x++ {
		grid.Nodes[x] = make([]*KurinPathfindingNode, kmap.Size.Base.Y)
		for y := int32(0); y < kmap.Size.Base.Y; y++ {
			grid.Nodes[x][y] = NewKurinPathfindingNode(kmap.Tiles[x][y][0])
		}
	}

	return grid
}

func GetNodeAt(grid *KurinPathfindingGrid, position sdl.Point) *KurinPathfindingNode {
	if position.X < 0 || position.Y < 0 || position.X >= grid.Size.X || position.Y >= grid.Size.Y {
		return nil
	}
	tile := grid.Nodes[position.X][position.Y]
	if !CanEnterKurinTile(tile.Tile) {
		return nil
	}

	return tile
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

	return FindKurinPath(grid, from, neighbours[0].Tile.Position)
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
