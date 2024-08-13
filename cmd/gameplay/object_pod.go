package gameplay

import "github.com/kelindar/binary"

type KurinObjectPodData struct {
	Enabled bool
}

func NewKurinObjectPod(tile *KurinTile) *KurinObject {
	obj := NewKurinObjectRaw(tile, "pod")
	obj.OnItemInteraction = func(object *KurinObject, item *KurinItem) bool {
		if item.Type == "credit" {
			if !RemoveKurinItemFromCharacterRaw(item, item.Character) {
				return false
			}
			PlaySound(&KurinGameInstance.SoundController, "jingle")
			KurinGameInstance.Credits++
			return true
		}

		return false
	}
	obj.EncodeData = func(object *KurinObject) []byte {
		objData := object.Data.(KurinObjectPodData)
		data, _ := binary.Marshal(&objData)
		return data
	}
	obj.DecodeData = func(object *KurinObject, data []byte) {
		var objData KurinObjectPodData 
		binary.Unmarshal(data, &objData)
		object.Data = objData
	}
	obj.Data = KurinObjectPodData{
		Enabled: false,
	}
	
	return obj
}
