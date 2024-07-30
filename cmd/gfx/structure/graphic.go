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
	Template  templates.KurinStructureTemplate
	Textures  []sdlutils.TextureWithSize
	Blueprint *sdlutils.TextureWithSize
}

func NewKurinStructureGraphic(renderer *gfx.KurinRenderer, structureId string) (*KurinStructureGraphic, *error) {
	graphicDirectory := path.Join(constants.TexturesPath, "structures", structureId)
	graphic := KurinStructureGraphic{
		Textures: make([]sdlutils.TextureWithSize, 4),
	}

	templateBytes, templateErr := os.ReadFile(path.Join(constants.DataPath, "templates", "structures", fmt.Sprintf("%s.json", structureId)))
	if templateErr != nil {
		return &graphic, &templateErr
	}
	if err := json.Unmarshal(templateBytes, &graphic.Template); err != nil {
		return &graphic, &err
	}

	num := 4
	if graphic.Template.Rotate != nil && !*graphic.Template.Rotate {
		num = 1
	}

	var err *error
	for i := 0; i < num; i++ {
		partPath := path.Join(graphicDirectory, fmt.Sprintf("%s_%d.png", structureId, i))
		if graphic.Textures[i], err = sdlutils.LoadTexture(renderer.Renderer, partPath); err != nil {
			return &graphic, err
		}
	}
	if blueprint, err := sdlutils.LoadTexture(renderer.Renderer, path.Join(graphicDirectory, fmt.Sprintf("%s_0_blueprint.png", structureId))); err == nil {
		graphic.Blueprint = &blueprint
	}

	return &graphic, nil
}
