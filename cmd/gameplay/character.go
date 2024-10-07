package gameplay

func PopulateCharacter(character *Mob) {
	GetInventory(character).Hands[HandLeft] = NewItem("survivalknife", 1)
	GetInventory(character).Hands[HandRight] = NewItem("welder", 1)
}
