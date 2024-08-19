package gameplay

type PathfindingQueue []*PathfindingNode

func (queue PathfindingQueue) Len() int {
	return len(queue)
}

func (queue PathfindingQueue) Less(i, j int) bool {
	return queue[i].Rank < queue[j].Rank
}

func (queue PathfindingQueue) Swap(i, j int) {
	queue[i], queue[j] = queue[j], queue[i]
	queue[i].Index = i
	queue[j].Index = j
}

func (queue *PathfindingQueue) Push(x interface{}) {
	n := len(*queue)
	no := x.(*PathfindingNode)
	no.Index = n
	*queue = append(*queue, no)
}

func (queue *PathfindingQueue) Pop() interface{} {
	old := *queue
	n := len(old)
	no := old[n-1]
	no.Index = -1
	*queue = old[0 : n-1]

	return no
}
