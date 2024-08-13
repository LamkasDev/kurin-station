package gameplay

func NewKurinObjectGrille(tile *KurinTile) *KurinObject {
	obj := NewKurinObjectRaw(tile, "grille")
	obj.OnItemInteraction = func(object *KurinObject, item *KurinItem) bool {
		switch data := item.Data.(type) {
		case KurinItemWelderData:
			if data.Enabled && object.Health < 3 {
				PlaySoundVolume(&KurinGameInstance.SoundController, "welder", 0.5)
				object.Health++
			}
			return true
		}

		return false
	}
	obj.OnDestroy = func(object *KurinObject) {
		CreateKurinObject(object.Tile, "broken_grille")
	}

	return obj
}
