package gameplay

import "github.com/LamkasDev/kurin/cmd/common/sdlutils"

func NewKurinItem(itemType string, transform *sdlutils.Transform) *KurinItem {
	switch itemType {
	case "welder":
		return NewKurinItemWelder(transform)
	}

	return &KurinItem{
		Type: itemType,
		Transform: transform,
		GetTextures: func(item *KurinItem, game *KurinGame) []int {
			return []int{0};
		},
		GetTextureHand: func(item *KurinItem, game *KurinGame) int {
			return 0;
		},
		Interact: func(item *KurinItem, game *KurinGame) {},
		Process: func(item *KurinItem, game *KurinGame) {},
		Data: nil,
	}
}
