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

type KurinItemGraphic struct {
	Template templates.KurinItemTemplate
	Textures  []sdlutils.TextureWithSizeAndSurface
	Outline  *sdlutils.TextureWithSize
	Hands    map[gameplay.KurinHand][][]sdlutils.TextureWithSize
}

func NewKurinItemGraphic(renderer *gfx.KurinRenderer, itemId string) (*KurinItemGraphic, error) {
	graphicDirectory := path.Join(constants.TexturesPath, "items", itemId)
	graphic := KurinItemGraphic{
		Hands: map[gameplay.KurinHand][][]sdlutils.TextureWithSize{},
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
	graphic.Textures = make([]sdlutils.TextureWithSizeAndSurface, textures)
	
	for i := 0; i < textures; i++ {
		graphicPath := path.Join(graphicDirectory, fmt.Sprintf("%s_%d.png", itemId, i))
		if graphic.Textures[i], err = sdlutils.LoadTextureWithSurface(renderer.Renderer, graphicPath); err != nil {
			return &graphic, err
		}
	}

	graphicOutlinePath := path.Join(graphicDirectory, fmt.Sprintf("%s_0_outline.png", itemId))
	if outline, err := sdlutils.LoadTexture(renderer.Renderer, graphicOutlinePath); err == nil {
		graphic.Outline = &outline
	}

	if graphic.Template.Hand == nil || *graphic.Template.Hand {
		textures := 1
		if graphic.Template.StatesHand != nil {
			textures = *graphic.Template.StatesHand
		}

		graphic.Hands[gameplay.KurinHandLeft] = make([][]sdlutils.TextureWithSize, textures)
		for j := 0; j < textures; j++ {
			graphic.Hands[gameplay.KurinHandLeft][j] = make([]sdlutils.TextureWithSize, 4)
			for i := 0; i < 4; i++ {
				handPath := path.Join(graphicDirectory, fmt.Sprintf("%s_left_%d_%d.png", itemId, j, i))
				if handTexture, err := sdlutils.LoadTexture(renderer.Renderer, handPath); err == nil {
					graphic.Hands[gameplay.KurinHandLeft][j][i] = handTexture
				}
			}
		}

		graphic.Hands[gameplay.KurinHandRight] = make([][]sdlutils.TextureWithSize, textures)
		for j := 0; j < textures; j++ {
			graphic.Hands[gameplay.KurinHandRight][j] = make([]sdlutils.TextureWithSize, 4)
			for i := 0; i < 4; i++ {
				handPath := path.Join(graphicDirectory, fmt.Sprintf("%s_right_%d_%d.png", itemId, j, i))
				if handTexture, err := sdlutils.LoadTexture(renderer.Renderer, handPath); err == nil {
					graphic.Hands[gameplay.KurinHandRight][j][i] = handTexture
				}
			}
		}
	}

	return &graphic, nil
}
