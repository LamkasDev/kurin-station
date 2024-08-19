package gameplay

type Sound struct {
	Type   string
	Volume float32
}

func NewSound(soundType string, volume float32) *Sound {
	return &Sound{
		Type:   soundType,
		Volume: volume,
	}
}
