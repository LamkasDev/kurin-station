package structure

import (
	"io/fs"
	"path"
	"path/filepath"

	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type KurinRendererLayerObjectData struct {
	Structures map[string]*KurinStructureGraphic
}

func NewKurinRendererLayerObject() *gfx.RendererLayer {
	return &gfx.RendererLayer{
		Load:   LoadKurinRendererLayerObject,
		Render: RenderKurinRendererLayerObject,
		Data: &KurinRendererLayerObjectData{
			Structures: map[string]*KurinStructureGraphic{},
		},
	}
}

func LoadKurinRendererLayerObject(layer *gfx.RendererLayer) error {
	return filepath.WalkDir(path.Join(constants.TexturesPath, "structures"), func(path string, d fs.DirEntry, err error) error {
		if d.Name() == "structures" || !d.IsDir() {
			return nil
		}
		if layer.Data.(*KurinRendererLayerObjectData).Structures[d.Name()], err = NewKurinStructureGraphic(d.Name()); err != nil {
			return err
		}

		return nil
	})
}

func RenderKurinRendererLayerObject(layer *gfx.RendererLayer) error {
	for _, obj := range gameplay.GameInstance.Map.Objects {
		RenderKurinObject(layer, obj)
	}

	return nil
}
