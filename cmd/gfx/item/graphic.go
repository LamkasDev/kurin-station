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
	Hands    map[gameplay.KurinHand][]sdlutils.TextureWithSize
}

func NewKurinItemGraphic(renderer *gfx.KurinRenderer, itemId string) (*KurinItemGraphic, *error) {
	graphicDirectory := path.Join(constants.TexturesPath, "items", itemId)
	graphic := KurinItemGraphic{
		Hands: map[gameplay.KurinHand][]sdlutils.TextureWithSize{},
	}

	templateBytes, templateErr := os.ReadFile(path.Join(constants.DataPath, "templates", "items", fmt.Sprintf("%s.json", itemId)))
	if templateErr != nil {
		return &graphic, &templateErr
	}
	if err := json.Unmarshal(templateBytes, &graphic.Template); err != nil {
		return &graphic, &err
	}

	textures := 1
	if graphic.Template.States != nil {
		textures = *graphic.Template.States
	}
	graphic.Textures = make([]sdlutils.TextureWithSizeAndSurface, textures)
	
	var err *error
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
		graphic.Hands[gameplay.KurinHandLeft] = make([]sdlutils.TextureWithSize, 4)
		for i := 0; i < 4; i++ {
			handPath := path.Join(graphicDirectory, fmt.Sprintf("%s_left_%d.png", itemId, i))
			if handTexture, err := sdlutils.LoadTexture(renderer.Renderer, handPath); err == nil {
				graphic.Hands[gameplay.KurinHandLeft][i] = handTexture
			}
		}

		graphic.Hands[gameplay.KurinHandRight] = make([]sdlutils.TextureWithSize, 4)
		for i := 0; i < 4; i++ {
			handPath := path.Join(graphicDirectory, fmt.Sprintf("%s_right_%d.png", itemId, i))
			if handTexture, err := sdlutils.LoadTexture(renderer.Renderer, handPath); err == nil {
				graphic.Hands[gameplay.KurinHandRight][i] = handTexture
			}
		}
	}

	return &graphic, nil
}
