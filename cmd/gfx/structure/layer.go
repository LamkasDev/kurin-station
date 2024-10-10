package structure

import (
	"io/fs"
	"path"
	"path/filepath"

	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type RendererLayerObjectData struct {
	Structures map[string]*StructureGraphic
}

func NewRendererLayerObject() *gfx.RendererLayer {
	return &gfx.RendererLayer{
		Load:   LoadRendererLayerObject,
		Render: RenderRendererLayerObject,
		Data: &RendererLayerObjectData{
			Structures: map[string]*StructureGraphic{},
		},
	}
}

func LoadRendererLayerObject(layer *gfx.RendererLayer) error {
	return filepath.WalkDir(path.Join(constants.TexturesPath, "structures"), func(path string, d fs.DirEntry, err error) error {
		if d.Name() == "structures" || !d.IsDir() {
			return nil
		}
		if layer.Data.(*RendererLayerObjectData).Structures[d.Name()], err = NewStructureGraphic(d.Name()); err != nil {
			return err
		}

		return nil
	})
}

func RenderRendererLayerObject(layer *gfx.RendererLayer) error {
	for _, obj := range gameplay.GameInstance.Map.Objects {
		if obj.Tile.Position.Z != gameplay.GameInstance.SelectedZ {
			continue
		}
		RenderObject(layer, obj)
	}

	return nil
}
