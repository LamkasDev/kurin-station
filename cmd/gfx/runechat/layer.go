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
			rawRunechat := gameplay.GameInstance.RunechatController.Messages[i]
			switch runechat := rawRunechat.Data.(type) {
			case gameplay.RunechatMobData:
				if runechat.Mob.Position.Z == gameplay.GameInstance.SelectedCharacter.Position.Z {
					if err := RenderRunechatCharacter(layer, rawRunechat, characterTally[runechat.Mob]); err != nil {
						return err
					}
				}
				characterTally[runechat.Mob]++
			}

			gameplay.ProcessRunechat(&gameplay.GameInstance.RunechatController, rawRunechat)
		}
	}

	return nil
}
