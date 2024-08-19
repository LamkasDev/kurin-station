package gameplay

type Hand uint8

const (
	HandLeft  = Hand(0)
	HandRight = Hand(1)
)

type Inventory struct {
	Hands map[Hand]*Item
}

func NewInventory() Inventory {
	return Inventory{
		Hands: map[Hand]*Item{
			HandLeft:  nil,
			HandRight: nil,
		},
	}
}

func FindItemInInventory(inventory *Inventory, itemType string) *Item {
	for _, item := range inventory.Hands {
		if item != nil && item.Type == itemType {
			return item
		}
	}

	return nil
}
