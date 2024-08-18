package gameplay

func NewKurinObjectGrille(tile *KurinTile) *KurinObject {
	obj := NewKurinObjectRaw[interface{}](tile, "grille")
	obj.OnInteraction = func(object *KurinObject, item *KurinItem) bool {
		if item != nil {
			switch data := item.Data.(type) {
			case KurinItemWelderData:
				if data.Enabled && object.Health < 3 {
					PlaySoundVolume(&GameInstance.SoundController, "welder", 0.5)
					object.Health++
				}
				return true
			}
		}

		return false
	}
	obj.OnDestroy = func(object *KurinObject) {
		CreateKurinObject(object.Tile, "broken_grille")
	}

	return obj
}
