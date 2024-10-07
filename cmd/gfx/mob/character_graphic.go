package mob

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

type HumanGraphicContainer struct {
	Genders map[string]*HumanGraphic
}

type HumanGraphic struct {
	Template templates.SpeciesTemplate
	Textures map[string][]*sdlutils.TextureWithSize
}

func NewHumanGraphicContainer(speciesId string) (*HumanGraphicContainer, error) {
	container := HumanGraphicContainer{
		Genders: map[string]*HumanGraphic{},
	}

	graphic, err := NewHumanGraphic(speciesId, gameplay.DefaultGender)
	if err != nil {
		return &container, err
	}
	container.Genders[gameplay.DefaultGender] = graphic

	return &container, nil
}

func NewHumanGraphic(speciesId string, speciesGender string) (*HumanGraphic, error) {
	graphic := HumanGraphic{
		Textures: map[string][]*sdlutils.TextureWithSize{},
	}

	templateBytes, err := os.ReadFile(path.Join(constants.DataPath, "templates", "species", fmt.Sprintf("%s.json", speciesId)))
	if err != nil {
		return &graphic, err
	}
	if err := json.Unmarshal(templateBytes, &graphic.Template); err != nil {
		return &graphic, err
	}

	for _, part := range graphic.Template.Parts {
		partDirectory := constants.TexturesPath
		if part.Path != nil {
			partDirectory = path.Join(partDirectory, *part.Path, part.Id)
		} else {
			partDirectory = path.Join(partDirectory, "bodyparts_greyscale", part.Id)
		}

		partFile := part.Id
		if part.Type != nil && *part.Type {
			partFile = fmt.Sprintf("%s_%s", partFile, speciesGender)
		}

		graphic.Textures[part.Id] = make([]*sdlutils.TextureWithSize, 4)
		for i := range 4 {
			partPath := path.Join(partDirectory, fmt.Sprintf("%s_%d.png", partFile, i))
			if graphic.Textures[part.Id][i], err = sdlutils.LoadTexture(gfx.RendererInstance.Renderer, partPath); err != nil {
				return &graphic, err
			}
		}
	}

	return &graphic, nil
}
