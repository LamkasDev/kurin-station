package gameplay

import "github.com/LamkasDev/kurin/cmd/common/sdlutils"

func NewObjectTemplateDisplaced() *ObjectTemplate {
	template := NewObjectTemplate[interface{}]("displaced", false)
	template.OnInteraction = func(object *Object, item *Item) bool {
		if item != nil && item.Type == "rod" {
			if !RemoveItemFromCharacterRaw(item, item.Mob) {
				return false
			}
			DestroyObjectRaw(GameInstance.Map, object)
			CreateObjectRaw(GameInstance.Map, object.Tile, "wall")
			return true
		}

		return false
	}
	template.OnDestroy = func(object *Object) {
		AddItemToMapRaw(
			NewItem("rod", 1),
			GameInstance.Map,
			&sdlutils.Transform{Position: sdlutils.Vector3ToFVector3Center(object.Tile.Position)},
		)
	}
	template.MaxHealth = 5

	return template
}
