package gameplay

type KurinHand uint8

const (
	KurinHandLeft  = KurinHand(0)
	KurinHandRight = KurinHand(1)
)

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

func FindItemInInventory(inventory *KurinInventory, itemType string) *KurinItem {
	for _, item := range inventory.Hands {
		if item != nil && item.Type == itemType {
			return item
		}
	}

	return nil
}
