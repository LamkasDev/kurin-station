package runechat

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type KurinRendererLayerRunechatData struct{}

func NewKurinRendererLayerRunechat() *gfx.RendererLayer {
	return &gfx.RendererLayer{
		Load:   LoadKurinRendererLayerRunechat,
		Render: RenderKurinRendererLayerRunechat,
		Data:   &KurinRendererLayerRunechatData{},
	}
}

func LoadKurinRendererLayerRunechat(layer *gfx.RendererLayer) error {
	return nil
}

func RenderKurinRendererLayerRunechat(layer *gfx.RendererLayer) error {
	if len(gameplay.GameInstance.RunechatController.Messages) > 0 {
		characterTally := map[*gameplay.KurinCharacter]int32{}
		for i := len(gameplay.GameInstance.RunechatController.Messages) - 1; i >= 0; i-- {
			runechat := gameplay.GameInstance.RunechatController.Messages[i]
			switch val := runechat.Data.(type) {
			case gameplay.KurinRunechatCharacterData:
				if err := RenderKurinRunechatCharacter(layer, runechat, characterTally[val.Character]); err != nil {
					return err
				}
				characterTally[val.Character]++
			}

			gameplay.ProcessKurinRunechat(&gameplay.GameInstance.RunechatController, runechat)
		}
	}

	return nil
}
