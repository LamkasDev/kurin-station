package item

import (
	"io/fs"
	"path"
	"path/filepath"

	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type KurinRendererLayerItemData struct {
	Items map[string]*KurinItemGraphic
}

func NewKurinRendererLayerItem() *gfx.RendererLayer {
	return &gfx.RendererLayer{
		Load:   LoadKurinRendererLayerItem,
		Render: RenderKurinRendererLayerItem,
		Data: &KurinRendererLayerItemData{
			Items: map[string]*KurinItemGraphic{},
		},
	}
}

func LoadKurinRendererLayerItem(layer *gfx.RendererLayer) error {
	return filepath.WalkDir(path.Join(constants.TexturesPath, "items"), func(path string, d fs.DirEntry, err error) error {
		if d.Name() == "items" || !d.IsDir() {
			return nil
		}
		graphic, err := NewKurinItemGraphic(d.Name())
		if err != nil {
			return err
		}
		layer.Data.(*KurinRendererLayerItemData).Items[d.Name()] = graphic

		return nil
	})
}

func RenderKurinRendererLayerItem(layer *gfx.RendererLayer) error {
	for _, item := range gameplay.GameInstance.Map.Items {
		RenderKurinItem(layer, item)
	}

	return nil
}
