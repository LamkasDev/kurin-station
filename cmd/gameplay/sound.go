package gameplay

type KurinSound struct {
	Type string
}

func NewKurinSound(stype string) *KurinSound {
	return &KurinSound{
		Type: stype,
	}
}
