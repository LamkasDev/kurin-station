package runechat

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type RendererLayerRunechatData struct{}

func NewRendererLayerRunechat() *gfx.RendererLayer {
	return &gfx.RendererLayer{
		Load:   LoadRendererLayerRunechat,
		Render: RenderRendererLayerRunechat,
		Data:   &RendererLayerRunechatData{},
	}
}

func LoadRendererLayerRunechat(layer *gfx.RendererLayer) error {
	return nil
}

func RenderRendererLayerRunechat(layer *gfx.RendererLayer) error {
	if len(gameplay.GameInstance.RunechatController.Messages) > 0 {
		characterTally := map[*gameplay.Mob]int32{}
		for i := len(gameplay.GameInstance.RunechatController.Messages) - 1; i >= 0; i-- {
			runechat := gameplay.GameInstance.RunechatController.Messages[i]
			switch val := runechat.Data.(type) {
			case gameplay.RunechatMobData:
				if err := RenderRunechatCharacter(layer, runechat, characterTally[val.Mob]); err != nil {
					return err
				}
				characterTally[val.Mob]++
			}

			gameplay.ProcessRunechat(&gameplay.GameInstance.RunechatController, runechat)
		}
	}

	return nil
}
