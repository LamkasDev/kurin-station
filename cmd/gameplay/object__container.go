package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
	"github.com/kelindar/binary"
)

func NewKurinObject(tile *KurinTile, objectType string) *KurinObject {
	switch objectType {
	case "broken_grille":
		return NewKurinObjectBrokenGrille(tile)
	case "grille":
		return NewKurinObjectGrille(tile)
	case "pod":
		return NewKurinObjectPod(tile)
	case "wall":
		return NewKurinObjectWall(tile)
	case "displaced":
		return NewKurinObjectDisplaced(tile)
	case "big_thruster":
		return NewKurinObjectBigThruster(tile)
	case "lathe":
		return NewKurinObjectLathe(tile)
	}

	return NewKurinObjectRaw[interface{}](tile, objectType)
}

func NewKurinObjectRaw[D any](tile *KurinTile, objectType string) *KurinObject {
	return &KurinObject{
		Id:        GetNextId(),
		Type:      objectType,
		Tile:      tile,
		Direction: common.KurinDirectionSouth,
		Health:    3,
		Process:   func(object *KurinObject) {},
		GetTexture: func(object *KurinObject) int {
			return 0
		},
		OnInteraction: func(object *KurinObject, item *KurinItem) bool {
			return false
		},
		OnCreate:  func(object *KurinObject) {},
		OnDestroy: func(object *KurinObject) {},
		EncodeData: func(object *KurinObject) []byte {
			if object.Data == nil {
				return []byte{}
			}

			objData := object.Data.(D)
			data, _ := binary.Marshal(&objData)
			return data
		},
		DecodeData: func(object *KurinObject, data []byte) {
			if len(data) == 0 {
				return
			}

			var objData D
			binary.Unmarshal(data, &objData)
			object.Data = objData
		},
	}
}

func InteractKurinObject(character *KurinCharacter, object *KurinObject) {
	KurinCharacterHitObject(character, object)
}
