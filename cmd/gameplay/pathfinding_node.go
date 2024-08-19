package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
)

type PathfindingNode struct {
	Position sdlutils.Vector3
	Parent   *PathfindingNode
	Cost     float32
	Rank     float32
	Index    int
}

func NewPathfindingNode(position sdlutils.Vector3) *PathfindingNode {
	return &PathfindingNode{
		Position: position,
	}
}

func GetPathfindingNodeNeighbourCost(from *PathfindingNode, to *PathfindingNode) float32 {
	return 1
}

func GetPathfindingNodeEstimatedCost(from *PathfindingNode, to *PathfindingNode) float32 {
	return float32(sdlutils.GetDistance(from.Position.Base, to.Position.Base))
}
