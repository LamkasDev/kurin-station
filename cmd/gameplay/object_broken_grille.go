package gameplay

import "github.com/LamkasDev/kurin/cmd/common/sdlutils"

func NewKurinObjectBrokenGrille(tile *KurinTile) *KurinObject {
	obj := NewKurinObjectRaw(tile, "broken_grille")
	obj.Health = 1
	obj.OnItemInteraction = func(object *KurinObject, item *KurinItem) bool {
		if item.Type == "rod" {
			if !RemoveKurinItemFromCharacterRaw(item, item.Character) {
				return false
			}
			DestroyKurinObjectRaw(&KurinGameInstance.Map, object)
			CreateKurinObjectRaw(&KurinGameInstance.Map, object.Tile, "grille")
			return true
		}

		return false
	}
	obj.OnDestroy = func(object *KurinObject) {
		AddKurinItemToMapRaw(NewKurinItem("rod"), &KurinGameInstance.Map, &sdlutils.Transform{Position: sdlutils.Vector3ToFVector3Center(object.Tile.Position)})
	}

	return obj
}
