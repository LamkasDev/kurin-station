package gameplay

func NewKurinItem(itemType string) *KurinItem {
	switch itemType {
	case "welder":
		return NewKurinItemWelder()
	}

	return NewKurinItemRaw(itemType)
}

func NewKurinItemRaw(itemType string) *KurinItem {
	return &KurinItem{
		Id:   GetNextId(),
		Type: itemType,
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
			return []byte{}
		},
		DecodeData: func(item *KurinItem, data []byte) {},
		Process:    func(item *KurinItem) {},
		CanHit:     true,
		Data:       nil,
	}
}
