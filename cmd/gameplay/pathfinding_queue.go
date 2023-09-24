package gameplay

type KurinPathfindingQueue []*KurinPathfindingNode

func (queue KurinPathfindingQueue) Len() int {
	return len(queue)
}

func (queue KurinPathfindingQueue) Less(i, j int) bool {
	return queue[i].Rank < queue[j].Rank
}

func (queue KurinPathfindingQueue) Swap(i, j int) {
	queue[i], queue[j] = queue[j], queue[i]
	queue[i].Index = i
	queue[j].Index = j
}

func (queue *KurinPathfindingQueue) Push(x interface{}) {
	n := len(*queue)
	no := x.(*KurinPathfindingNode)
	no.Index = n
	*queue = append(*queue, no)
}

func (queue *KurinPathfindingQueue) Pop() interface{} {
	old := *queue
	n := len(old)
	no := old[n-1]
	no.Index = -1
	*queue = old[0 : n-1]

	return no
}
