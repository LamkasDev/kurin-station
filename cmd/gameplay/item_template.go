package gameplay

import "github.com/kelindar/binary"

type ItemTemplate struct {
	Type              string
	Process           ItemProcess
	GetTextures       ItemGetTextures
	GetTextureHand    ItemGetTextureHand
	OnHandInteraction ItemOnHandInteraction
	OnTileInteraction ItemOnTileInteraction
	EncodeData        ItemEncodeData
	DecodeData        ItemDecodeData
	GetDefaultData    ItemGetDefaultData
	CanHit            bool
	MaxCount          uint16
}

type (
	ItemProcess           func(item *Item)
	ItemGetTextures       func(item *Item) []int
	ItemGetTextureHand    func(item *Item) int
	ItemOnHandInteraction func(item *Item)
	ItemOnTileInteraction func(item *Item, tile *Tile) bool
	ItemEncodeData        func(item *Item) []byte
	ItemDecodeData        func(item *Item, data []byte)
	ItemGetDefaultData    func() interface{}
)

func NewItemTemplate[D any](itemType string, maxCount uint16) *ItemTemplate {
	return &ItemTemplate{
		Type:    itemType,
		Process: func(item *Item) {},
		GetTextures: func(item *Item) []int {
			return []int{0}
		},
		GetTextureHand: func(item *Item) int {
			return 0
		},
		OnHandInteraction: func(item *Item) {},
		OnTileInteraction: func(item *Item, tile *Tile) bool {
			return false
		},
		EncodeData: func(item *Item) []byte {
			if item.Data == nil {
				return []byte{}
			}

			itemData := item.Data.(D)
			data, _ := binary.Marshal(&itemData)
			return data
		},
		DecodeData: func(item *Item, data []byte) {
			if len(data) == 0 {
				return
			}

			var itemData D
			binary.Unmarshal(data, &itemData)
			item.Data = itemData
		},
		GetDefaultData: func() interface{} {
			return nil
		},
		CanHit:   true,
		MaxCount: maxCount,
	}
}
