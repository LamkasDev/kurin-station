package animation

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/LamkasDev/kurin/cmd/gameplay/templates"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type KurinAnimationGraphic struct {
	Template templates.KurinAnimationTemplate
}

func NewKurinAnimationGraphic(_ *gfx.KurinRenderer, animationId string) (*KurinAnimationGraphic, *error) {
	graphic := KurinAnimationGraphic{}

	templateBytes, templateErr := os.ReadFile(path.Join(constants.DataPath, "templates", "animations", fmt.Sprintf("%s.json", animationId)))
	if templateErr != nil {
		return &graphic, &templateErr
	}
	if err := json.Unmarshal(templateBytes, &graphic.Template); err != nil {
		return &graphic, &err
	}

	return &graphic, nil
}
