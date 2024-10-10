package item

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gameplay/templates"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type ItemGraphic struct {
	Template templates.ItemTemplate
	Textures []*sdlutils.TextureWithSizeAndSurface
	Outline  *sdlutils.TextureWithSize
	Hands    map[gameplay.Hand][][]*sdlutils.TextureWithSize
}

func NewItemGraphic(itemId string) (*ItemGraphic, error) {
	graphicDirectory := path.Join(constants.TexturesPath, "items", itemId)
	graphic := ItemGraphic{
		Hands: map[gameplay.Hand][][]*sdlutils.TextureWithSize{},
	}

	templateBytes, err := os.ReadFile(path.Join(constants.DataPath, "templates", "items", fmt.Sprintf("%s.json", itemId)))
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
	graphic.Textures = make([]*sdlutils.TextureWithSizeAndSurface, textures)

	for i := 0; i < textures; i++ {
		graphicPath := path.Join(graphicDirectory, fmt.Sprintf("%s_%d.png", itemId, i))
		if graphic.Textures[i], err = sdlutils.LoadTextureWithSurface(gfx.RendererInstance.Renderer, graphicPath); err != nil {
			return &graphic, err
		}
	}

	if graphic.Template.Outline == nil || *graphic.Template.Outline {
		graphicOutlinePath := path.Join(graphicDirectory, fmt.Sprintf("%s_0_outline.png", itemId))
		if outline, err := sdlutils.LoadTexture(gfx.RendererInstance.Renderer, graphicOutlinePath); err == nil {
			graphic.Outline = outline
		}
	}

	if graphic.Template.Hand != nil && *graphic.Template.Hand {
		textures := 1
		if graphic.Template.StatesHand != nil {
			textures = *graphic.Template.StatesHand
		}

		graphic.Hands[gameplay.HandLeft] = make([][]*sdlutils.TextureWithSize, textures)
		for j := 0; j < textures; j++ {
			graphic.Hands[gameplay.HandLeft][j] = make([]*sdlutils.TextureWithSize, 4)
			for i := 0; i < 4; i++ {
				handPath := path.Join(graphicDirectory, fmt.Sprintf("%s_left_%d_%d.png", itemId, j, i))
				if handTexture, err := sdlutils.LoadTexture(gfx.RendererInstance.Renderer, handPath); err == nil {
					graphic.Hands[gameplay.HandLeft][j][i] = handTexture
				}
			}
		}

		graphic.Hands[gameplay.HandRight] = make([][]*sdlutils.TextureWithSize, textures)
		for j := 0; j < textures; j++ {
			graphic.Hands[gameplay.HandRight][j] = make([]*sdlutils.TextureWithSize, 4)
			for i := 0; i < 4; i++ {
				handPath := path.Join(graphicDirectory, fmt.Sprintf("%s_right_%d_%d.png", itemId, j, i))
				if handTexture, err := sdlutils.LoadTexture(gfx.RendererInstance.Renderer, handPath); err == nil {
					graphic.Hands[gameplay.HandRight][j][i] = handTexture
				}
			}
		}
	}

	return &graphic, nil
}
