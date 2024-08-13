package gameplay

type KurinSound struct {
	Type   string
	Volume float32
}

func NewKurinSound(soundType string, volume float32) *KurinSound {
	return &KurinSound{
		Type:   soundType,
		Volume: volume,
	}
}
