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
	Template     templates.TurfTemplate
	Textures     []*sdlutils.TextureWithSize
	BlurTextures []*sdlutils.TextureWithSize
	Blueprint    *sdlutils.TextureWithSize
}

func NewTurfGraphic(tileId string) (*TurfGraphic, error) {
	graphicDirectory := path.Join(constants.TexturesPath, "turfs", tileId)
	graphic := TurfGraphic{
		BlurTextures: make([]*sdlutils.TextureWithSize, 1),
	}

	templateBytes, err := os.ReadFile(path.Join(constants.DataPath, "templates", "turfs", fmt.Sprintf("%s.json", tileId)))
	if err != nil {
		return &graphic, err
	}
	if err := json.Unmarshal(templateBytes, &graphic.Template); err != nil {
		return &graphic, err
	}

	textures := 1
	if graphic.Template.States != nil {
		textures = *graphic.Template.States
	}
	graphic.Textures = make([]*sdlutils.TextureWithSize, textures)

	for i := 0; i < textures; i++ {
		graphicPath := path.Join(graphicDirectory, fmt.Sprintf("%s_%d.png", tileId, i))
		if graphic.Textures[i], err = sdlutils.LoadTexture(gfx.RendererInstance.Renderer, graphicPath); err != nil {
			return &graphic, err
		}
	}

	if graphic.BlurTextures[0], err = sdlutils.LoadTexture(gfx.RendererInstance.Renderer, path.Join(graphicDirectory, fmt.Sprintf("%s_blur.png", tileId))); err != nil {
		return &graphic, err
	}

	if blueprint, err := sdlutils.LoadTexture(gfx.RendererInstance.Renderer, path.Join(graphicDirectory, fmt.Sprintf("%s_blueprint.png", tileId))); err == nil {
		graphic.Blueprint = blueprint
	}

	return &graphic, nil
}
