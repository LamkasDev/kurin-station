package gameplay

import "github.com/LamkasDev/kurin/cmd/common/sdlutils"

func NewKurinObjectBrokenGrille(tile *KurinTile) *KurinObject {
	obj := NewKurinObjectRaw[interface{}](tile, "broken_grille")
	obj.Health = 1
	obj.OnInteraction = func(object *KurinObject, item *KurinItem) bool {
		if item != nil && item.Type == "rod" {
			if !RemoveKurinItemFromCharacterRaw(item, item.Character) {
				return false
			}
			DestroyKurinObjectRaw(&GameInstance.Map, object)
			CreateKurinObjectRaw(&GameInstance.Map, object.Tile, "grille")
			return true
		}

		return false
	}
	obj.OnDestroy = func(object *KurinObject) {
		AddKurinItemToMapRaw(NewKurinItem("rod", 1), &GameInstance.Map, &sdlutils.Transform{Position: sdlutils.Vector3ToFVector3Center(object.Tile.Position)})
	}

	return obj
}
