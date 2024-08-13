package runechat

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
)

type KurinRendererLayerRunechatData struct{}

func NewKurinRendererLayerRunechat() *gfx.KurinRendererLayer {
	return &gfx.KurinRendererLayer{
		Load:   LoadKurinRendererLayerRunechat,
		Render: RenderKurinRendererLayerRunechat,
		Data:   KurinRendererLayerRunechatData{},
	}
}

func LoadKurinRendererLayerRunechat(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer) error {
	return nil
}

func RenderKurinRendererLayerRunechat(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer) error {
	if len(gameplay.KurinGameInstance.RunechatController.Messages) > 0 {
		characterTally := map[*gameplay.KurinCharacter]int32{}
		for i := len(gameplay.KurinGameInstance.RunechatController.Messages) - 1; i >= 0; i-- {
			runechat := gameplay.KurinGameInstance.RunechatController.Messages[i]
			switch val := runechat.Data.(type) {
			case gameplay.KurinRunechatCharacterData:
				if err := RenderKurinRunechatCharacter(renderer, layer, runechat, characterTally[val.Character]); err != nil {
					return err
				}
				characterTally[val.Character]++
			}

			gameplay.ProcessKurinRunechat(&gameplay.KurinGameInstance.RunechatController, runechat)
		}
	}

	return nil
}
