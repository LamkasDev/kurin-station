package turf

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

type KurinTurfGraphic struct {
	Template templates.KurinTurfTemplate
	Textures []sdlutils.TextureWithSize
}

func NewKurinTurfGraphic(renderer *gfx.KurinRenderer, tileId string) (*KurinTurfGraphic, *error) {
	graphicDirectory := path.Join(constants.TexturesPath, "turfs", tileId)
	graphic := KurinTurfGraphic{
		Textures: make([]sdlutils.TextureWithSize, 4),
	}

	templateBytes, templateErr := os.ReadFile(path.Join(constants.DataPath, "templates", "turfs", fmt.Sprintf("%s.json", tileId)))
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
		partPath := path.Join(graphicDirectory, fmt.Sprintf("%s_%d.png", tileId, i))
		if graphic.Textures[i], err = sdlutils.LoadTexture(renderer.Renderer, partPath); err != nil {
			return &graphic, err
		}
	}

	return &graphic, nil
}
