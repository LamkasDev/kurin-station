package gameplay

type KurinObjectPodData struct {
	Enabled bool
}

func NewKurinObjectPod(tile *KurinTile) *KurinObject {
	obj := NewKurinObjectRaw[*KurinObjectPodData](tile, "pod")
	obj.Health = 0
	obj.OnInteraction = func(object *KurinObject, item *KurinItem) bool {
		if item != nil && item.Type == "credit" {
			if !RemoveKurinItemFromCharacterRaw(item, item.Character) {
				return false
			}
			PlaySound(&GameInstance.SoundController, "jingle")
			GameInstance.Credits++
			return true
		}

		OpenKurinDialog(&KurinDialogRequest{Type: "pod", Data: &KurinDialogPodData{Pod: object}})
		return true
	}
	obj.Data = &KurinObjectPodData{
		Enabled: false,
	}

	return obj
}
