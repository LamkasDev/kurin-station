package gameplay

import "github.com/kelindar/binary"

type ObjectTemplate struct {
	Type           string
	MaxHealth      uint16
	Requirements   []ItemRequirement
	Smooth         bool
	Process        ObjectProcess
	IsPassable     ObjectIsPassable
	GetTexture     ObjectGetTexture
	OnInteraction  ObjectOnInteraction
	OnCreate       ObjectOnDestroy
	OnDestroy      ObjectOnDestroy
	EncodeData     ObjectEncodeData
	DecodeData     ObjectDecodeData
	GetDefaultData ObjectGetDefaultData
}

type (
	ObjectProcess        func(object *Object)
	ObjectIsPassable     func(object *Object) bool
	ObjectGetTexture     func(object *Object) int
	ObjectOnInteraction  func(object *Object, item *Item) bool
	ObjectOnCreate       func(object *Object)
	ObjectOnDestroy      func(object *Object)
	ObjectEncodeData     func(object *Object) []byte
	ObjectDecodeData     func(object *Object, data []byte)
	ObjectGetDefaultData func() interface{}
)

func NewObjectTemplate[D any](objectType string, smooth bool) *ObjectTemplate {
	return &ObjectTemplate{
		Type:      objectType,
		MaxHealth: 3,
		Requirements: []ItemRequirement{
			{
				Type:  "rod",
				Count: 1,
			},
		},
		Smooth:  smooth,
		Process: func(object *Object) {},
		IsPassable: func(object *Object) bool {
			return false
		},
		GetTexture: func(object *Object) int {
			return 0
		},
		OnInteraction: func(object *Object, item *Item) bool {
			return false
		},
		OnCreate:  func(object *Object) {},
		OnDestroy: func(object *Object) {},
		EncodeData: func(object *Object) []byte {
			if object.Data == nil {
				return []byte{}
			}

			objData := object.Data.(D)
			data, _ := binary.Marshal(&objData)
			return data
		},
		DecodeData: func(object *Object, data []byte) {
			if len(data) == 0 {
				return
			}

			var objData D
			binary.Unmarshal(data, &objData)
			object.Data = objData
		},
		GetDefaultData: func() interface{} {
			return nil
		},
	}
}
