package animation

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/LamkasDev/kurin/cmd/gameplay/templates"
)

type AnimationGraphic struct {
	Template templates.AnimationTemplate
}

func NewAnimationGraphic(animationId string) (*AnimationGraphic, error) {
	graphic := AnimationGraphic{}

	templateBytes, err := os.ReadFile(path.Join(constants.DataPath, "templates", "animations", fmt.Sprintf("%s.json", animationId)))
	if err != nil {
		return &graphic, err
	}
	if err := json.Unmarshal(templateBytes, &graphic.Template); err != nil {
		return &graphic, err
	}

	return &graphic, nil
}
