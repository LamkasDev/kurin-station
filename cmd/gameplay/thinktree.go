package gameplay

type Thinktree struct {
	Nodes []*ThinktreeNode
}

func NewThinktree() Thinktree {
	return Thinktree{
		Nodes: []*ThinktreeNode{},
	}
}

func NewThinktreeBasic() Thinktree {
	thinktree := NewThinktree()
	thinktree.Nodes = append(thinktree.Nodes, NewThinktreeNode("revenge"))
	thinktree.Nodes = append(thinktree.Nodes, NewThinktreeNode("attack"))
	thinktree.Nodes = append(thinktree.Nodes, NewThinktreeNode("panic"))
	thinktree.Nodes = append(thinktree.Nodes, NewThinktreeNode("wander"))

	return thinktree
}

func ProcessThinktree(mob *Mob) {
	for _, node := range mob.Thinktree.Nodes {
		if node.Template.Process(mob, node) {
			break
		}
	}
}
