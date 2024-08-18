package gameplay

func NewKurinObjectWall(tile *KurinTile) *KurinObject {
	obj := NewKurinObjectRaw[interface{}](tile, "wall")
	obj.Health = 5
	obj.OnInteraction = func(object *KurinObject, item *KurinItem) bool {
		if item != nil {
			switch data := item.Data.(type) {
			case KurinItemWelderData:
				if data.Enabled && object.Health < 5 {
					PlaySoundVolume(&GameInstance.SoundController, "welder", 0.5)
					object.Health++
				}
				return true
			}
		}

		return false
	}
	obj.OnDestroy = func(object *KurinObject) {
		CreateKurinObject(object.Tile, "displaced")
	}

	return obj
}
