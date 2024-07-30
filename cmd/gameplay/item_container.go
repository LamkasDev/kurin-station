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
		Data: nil,
	}
}
