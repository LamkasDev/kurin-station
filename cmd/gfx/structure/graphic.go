package structure

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay/templates"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type KurinStructureGraphic struct {
	Template       templates.KurinStructureTemplate
	Textures       [][]*sdlutils.TextureWithSize
	TexturesSmooth map[string]*sdlutils.TextureWithSize
	Blueprint      *sdlutils.TextureWithSize
}

func NewKurinStructureGraphic(structureId string) (*KurinStructureGraphic, error) {
	graphicDirectory := path.Join(constants.TexturesPath, "structures", structureId)
	graphic := KurinStructureGraphic{
		Textures:       make([][]*sdlutils.TextureWithSize, 4),
		TexturesSmooth: map[string]*sdlutils.TextureWithSize{},
	}

	templateBytes, err := os.ReadFile(path.Join(constants.DataPath, "templates", "structures", fmt.Sprintf("%s.json", structureId)))
	if err != nil {
		return &graphic, err
	}
	if err := json.Unmarshal(templateBytes, &graphic.Template); err != nil {
		return &graphic, err
	}

	if graphic.Template.Rotate != nil && *graphic.Template.Rotate {
		list := []string{
			"n", "e", "s", "w",
		}
		for i, direction := range list {
			graphic.Textures[i] = make([]*sdlutils.TextureWithSize, 1)
			partPath := path.Join(graphicDirectory, fmt.Sprintf("%s_%s.png", structureId, direction))
			if graphic.Textures[i][0], err = sdlutils.LoadTexture(gfx.RendererInstance.Renderer, partPath); err != nil {
				return &graphic, err
			}
		}
	} else if graphic.Template.States != nil {
		graphic.Textures[0] = make([]*sdlutils.TextureWithSize, *graphic.Template.States)
		for i := 0; i < *graphic.Template.States; i++ {
			partPath := path.Join(graphicDirectory, fmt.Sprintf("%s_%d.png", structureId, i))
			if graphic.Textures[0][i], err = sdlutils.LoadTexture(gfx.RendererInstance.Renderer, partPath); err != nil {
				return &graphic, err
			}
		}
	} else {
		graphic.Textures[0] = make([]*sdlutils.TextureWithSize, 1)
		partPath := path.Join(graphicDirectory, fmt.Sprintf("%s.png", structureId))
		if graphic.Textures[0][0], err = sdlutils.LoadTexture(gfx.RendererInstance.Renderer, partPath); err != nil {
			return &graphic, err
		}
	}

	if graphic.Template.Smooth != nil && *graphic.Template.Smooth {
		list := []string{
			"n", "e", "s", "w", "ne", "es", "sw", "nw", "nes", "esw", "nsw", "new", "nesw",
			"ns", "ew",
		}
		graphic.TexturesSmooth[""] = graphic.Textures[0][0]
		for _, direction := range list {
			partPath := path.Join(graphicDirectory, fmt.Sprintf("%s_%s.png", structureId, direction))
			if graphic.TexturesSmooth[direction], err = sdlutils.LoadTexture(gfx.RendererInstance.Renderer, partPath); err != nil {
				return &graphic, err
			}
		}
	}

	if blueprint, err := sdlutils.LoadTexture(gfx.RendererInstance.Renderer, path.Join(graphicDirectory, fmt.Sprintf("%s_blueprint.png", structureId))); err == nil {
		graphic.Blueprint = blueprint
	}

	return &graphic, nil
}
