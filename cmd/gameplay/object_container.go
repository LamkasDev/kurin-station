package gameplay

import "github.com/veandco/go-sdl2/sdl"

func NewKurinObject(tile *KurinTile, objectType string) *KurinObject {
	switch objectType {
	case "broken_grille":
		return NewKurinObjectBrokenGrille(tile)
	case "grille":
		return NewKurinObjectGrille(tile)
	case "pod":
		return NewKurinObjectPod(tile)
	}

	return NewKurinObjectRaw(tile, objectType)
}

func NewKurinObjectRaw(tile *KurinTile, objectType string) *KurinObject {
	return &KurinObject{
		Id: GetNextId(),
		Type:      objectType,
		Tile: tile,
		Direction: KurinDirectionSouth,
		Health: 3,
		Process: func(object *KurinObject) {},
		OnItemInteraction: func(object *KurinObject, item *KurinItem) bool {
			return false
		},
		OnCreate:  func(object *KurinObject) {},
		OnDestroy: func(object *KurinObject) {},
		EncodeData: func(object *KurinObject) []byte {
			return []byte{}
		},
		DecodeData: func(object *KurinObject, data []byte) {},
	}
}

func InteractKurinObject(character *KurinCharacter, object *KurinObject) {
	KurinCharacterHitObject(character, object)
}

func GetKurinObjectSize(object *KurinObject) sdl.Point {
	switch object.Type {
	case "pod":
		return sdl.Point{X: 2, Y: 2}
	}

	return sdl.Point{X: 1, Y: 1}
}
