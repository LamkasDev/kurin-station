package gameplay

type CharacterThinktree struct {
	Ticks int32
}

func NewCharacterThinktree() CharacterThinktree {
	return CharacterThinktree{
		Ticks: 0,
	}
}

func ProcessCharacterThinktree(character *Character) {
}
