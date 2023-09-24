package gameplay

type KurinCharacterThinktree struct {
	Ticks int32
}

func NewKurinCharacterThinktree() KurinCharacterThinktree {
	return KurinCharacterThinktree{
		Ticks: 0,
	}
}

func ProcessKurinCharacterThinktree(character *KurinCharacter) {
}
