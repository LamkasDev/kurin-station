package gameplay

type KurinInventory struct {
	Hands map[KurinHand]*KurinItem
}

func NewKurinInventory() KurinInventory {
	return KurinInventory{
		Hands: map[KurinHand]*KurinItem{
			KurinHandLeft:  nil,
			KurinHandRight: nil,
		},
	}
}
