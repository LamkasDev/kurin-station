package gameplay

type KurinObject struct {
	Type      string
	Direction KurinDirection
}

func NewKurinObject(otype string) *KurinObject {
	return &KurinObject{
		Type:      otype,
		Direction: KurinDirectionSouth,
	}
}
