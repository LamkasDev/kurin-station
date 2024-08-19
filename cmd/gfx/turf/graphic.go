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

type TurfGraphic struct {
	Template  templates.TurfTemplate
	Textures  []*sdlutils.TextureWithSize
	Blueprint *sdlutils.TextureWithSize
}

func NewTurfGraphic(tileId string) (*TurfGraphic, error) {
	graphicDirectory := path.Join(constants.TexturesPath, "turfs", tileId)
	graphic := TurfGraphic{
		Textures: make([]*sdlutils.TextureWithSize, 4),
	}

	templateBytes, err := os.ReadFile(path.Join(constants.DataPath, "templates", "turfs", fmt.Sprintf("%s.json", tileId)))
	if err != nil {
		return &graphic, err
	}
	if err := json.Unmarshal(templateBytes, &graphic.Template); err != nil {
		return &graphic, err
	}

	partPath := path.Join(graphicDirectory, fmt.Sprintf("%s.png", tileId))
	if graphic.Textures[0], err = sdlutils.LoadTexture(gfx.RendererInstance.Renderer, partPath); err != nil {
		return &graphic, err
	}

	if blueprint, err := sdlutils.LoadTexture(gfx.RendererInstance.Renderer, path.Join(graphicDirectory, fmt.Sprintf("%s_blueprint.png", tileId))); err == nil {
		graphic.Blueprint = blueprint
	}

	return &graphic, nil
}
