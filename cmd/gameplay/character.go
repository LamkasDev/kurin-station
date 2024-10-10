package gameplay

func PopulateCharacter(character *Mob) {
	GetInventory(character).ActiveHand = HandLeft
	AddItemToCharacterRaw(NewItem("gun", 1), character)
	GetInventory(character).ActiveHand = HandRight
	AddItemToCharacterRaw(NewItem("welder", 1), character)
}
