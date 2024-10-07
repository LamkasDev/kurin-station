package gameplay

type Hand uint8

const (
	HandLeft  = Hand(0)
	HandRight = Hand(1)
)

type Inventory struct {
	ActiveHand Hand
	Hands      map[Hand]*Item
}

func NewInventory() *Inventory {
	return &Inventory{
		ActiveHand: HandLeft,
		Hands: map[Hand]*Item{
			HandLeft:  nil,
			HandRight: nil,
		},
	}
}

func GetInventory(character *Mob) *Inventory {
	return character.Data.(*MobCharacterData).Inventory
}

func GetItemInHand(character *Mob, hand Hand) *Item {
	return GetInventory(character).Hands[hand]
}

func GetActiveHand(character *Mob) Hand {
	return GetInventory(character).ActiveHand
}

func GetHeldItem(character *Mob) *Item {
	return GetItemInHand(character, GetActiveHand(character))
}

func FindItemInInventory(inventory *Inventory, itemType string) *Item {
	for _, item := range inventory.Hands {
		if item != nil && item.Type == itemType {
			return item
		}
	}

	return nil
}
