package item

import (
	"io/fs"
	"path"
	"path/filepath"

	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type RendererLayerItemData struct {
	Items map[string]*ItemGraphic
}

func NewRendererLayerItem() *gfx.RendererLayer {
	return &gfx.RendererLayer{
		Load:   LoadRendererLayerItem,
		Render: RenderRendererLayerItem,
		Data: &RendererLayerItemData{
			Items: map[string]*ItemGraphic{},
		},
	}
}

func LoadRendererLayerItem(layer *gfx.RendererLayer) error {
	return filepath.WalkDir(path.Join(constants.TexturesPath, "items"), func(path string, d fs.DirEntry, err error) error {
		if d.Name() == "items" || !d.IsDir() {
			return nil
		}
		graphic, err := NewItemGraphic(d.Name())
		if err != nil {
			return err
		}
		layer.Data.(*RendererLayerItemData).Items[d.Name()] = graphic

		return nil
	})
}

func RenderRendererLayerItem(layer *gfx.RendererLayer) error {
	for _, item := range gameplay.GameInstance.Map.Items {
		if item.Transform.Position.Z != gameplay.GameInstance.SelectedCharacter.Position.Z {
			continue
		}
		RenderItem(layer, item)
	}

	return nil
}
