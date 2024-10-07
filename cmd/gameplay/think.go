package gameplay

type Thinktree struct {
	Ticks int32
}

func NewThinktree() Thinktree {
	return Thinktree{
		Ticks: 0,
	}
}

func ProcessThinktree(mob *Mob) {
}
