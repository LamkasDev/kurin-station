package gameplay

import "github.com/kelindar/binary"

func NewKurinItem(itemType string, count uint16) *KurinItem {
	switch itemType {
	case "welder":
		return NewKurinItemWelder()
	case "rod":
		item := NewKurinItemRaw[interface{}](itemType, count)
		item.MaxCount = 3
		return item
	}

	return NewKurinItemRaw[interface{}](itemType, count)
}

func NewKurinItemRaw[D any](itemType string, count uint16) *KurinItem {
	return &KurinItem{
		Id:    GetNextId(),
		Type:  itemType,
		Count: count,
		GetTextures: func(item *KurinItem) []int {
			return []int{0}
		},
		GetTextureHand: func(item *KurinItem) int {
			return 0
		},
		OnHandInteraction: func(item *KurinItem) {},
		OnTileInteraction: func(item *KurinItem, tile *KurinTile) bool {
			return false
		},
		EncodeData: func(item *KurinItem) []byte {
			if item.Data == nil {
				return []byte{}
			}

			itemData := item.Data.(D)
			data, _ := binary.Marshal(&itemData)
			return data
		},
		DecodeData: func(item *KurinItem, data []byte) {
			if len(data) == 0 {
				return
			}

			var itemData D
			binary.Unmarshal(data, &itemData)
			item.Data = itemData
		},
		Process:  func(item *KurinItem) {},
		CanHit:   true,
		MaxCount: 1,
		Data:     nil,
	}
}
