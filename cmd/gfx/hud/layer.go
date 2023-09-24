package hud

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/item"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinRendererLayerHUDData struct {
	Icons     map[string]*KurinHUDGraphic
	ItemLayer *gfx.KurinRendererLayer
}

func NewKurinRendererLayerHUD(itemLayer *gfx.KurinRendererLayer) *gfx.KurinRendererLayer {
	return &gfx.KurinRendererLayer{
		Load:   LoadKurinRendererLayerHUD,
		Render: RenderKurinRendererLayerHUD,
		Data: KurinRendererLayerHUDData{
			Icons:     map[string]*KurinHUDGraphic{},
			ItemLayer: itemLayer,
		},
	}
}

func LoadKurinRendererLayerHUD(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer) *error {
	return LoadKurinRendererLayerHUDArray(renderer, layer, []string{"hand_l", "lhandactive", "hand_r", "rhandactive", "act_equip", "swap_1", "swap_2", "pda", "selector", "template", "template_active"})
}

func LoadKurinRendererLayerHUDArray(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer, textures []string) *error {
	var err *error
	for _, texture := range textures {
		if layer.Data.(KurinRendererLayerHUDData).Icons[texture], err = NewKurinHUDGraphic(renderer, texture); err != nil {
			return err
		}
	}

	return nil
}

func RenderKurinRendererLayerHUD(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer, game *gameplay.KurinGame) *error {
	if renderer.WindowContext.CameraMode != gfx.KurinRendererCameraModeCharacter {
		return nil
	}

	half := gfx.GetHalfWindowSize(&renderer.WindowContext)

	sdlutils.RenderTexture(renderer.Renderer, layer.Data.(KurinRendererLayerHUDData).Icons["hand_l"].Texture, KurinHUDElementHandLeft.GetPosition(renderer.WindowContext.WindowSize), sdl.FPoint{X: 2, Y: 2})
	if game.SelectedCharacter.ActiveHand == gameplay.KurinHandLeft {
		sdlutils.RenderTexture(renderer.Renderer, layer.Data.(KurinRendererLayerHUDData).Icons["lhandactive"].Texture, sdl.Point{X: half.X, Y: renderer.WindowContext.WindowSize.Y - 72}, sdl.FPoint{X: 2, Y: 2})
	}
	lhand := game.SelectedCharacter.Inventory.Hands[gameplay.KurinHandLeft]
	if lhand != nil {
		sdlutils.RenderTexture(renderer.Renderer, layer.Data.(KurinRendererLayerHUDData).ItemLayer.Data.(item.KurinRendererLayerItemData).Items[lhand.Type].Texture.Base, sdl.Point{X: half.X, Y: renderer.WindowContext.WindowSize.Y - 72}, sdl.FPoint{X: 2, Y: 2})
	}

	sdlutils.RenderTexture(renderer.Renderer, layer.Data.(KurinRendererLayerHUDData).Icons["hand_r"].Texture, KurinHUDElementHandRight.GetPosition(renderer.WindowContext.WindowSize), sdl.FPoint{X: 2, Y: 2})
	if game.SelectedCharacter.ActiveHand == gameplay.KurinHandRight {
		sdlutils.RenderTexture(renderer.Renderer, layer.Data.(KurinRendererLayerHUDData).Icons["rhandactive"].Texture, sdl.Point{X: half.X - 64, Y: renderer.WindowContext.WindowSize.Y - 72}, sdl.FPoint{X: 2, Y: 2})
	}
	rhand := game.SelectedCharacter.Inventory.Hands[gameplay.KurinHandRight]
	if rhand != nil {
		sdlutils.RenderTexture(renderer.Renderer, layer.Data.(KurinRendererLayerHUDData).ItemLayer.Data.(item.KurinRendererLayerItemData).Items[rhand.Type].Texture.Base, sdl.Point{X: half.X - 64, Y: renderer.WindowContext.WindowSize.Y - 72}, sdl.FPoint{X: 2, Y: 2})
	}

	sdlutils.RenderTexture(renderer.Renderer, layer.Data.(KurinRendererLayerHUDData).Icons["act_equip"].Texture, sdl.Point{X: half.X - 64, Y: renderer.WindowContext.WindowSize.Y - 136}, sdl.FPoint{X: 2, Y: 2})
	sdlutils.RenderTexture(renderer.Renderer, layer.Data.(KurinRendererLayerHUDData).Icons["swap_1"].Texture, sdl.Point{X: half.X - 64, Y: renderer.WindowContext.WindowSize.Y - 136}, sdl.FPoint{X: 2, Y: 2})
	sdlutils.RenderTexture(renderer.Renderer, layer.Data.(KurinRendererLayerHUDData).Icons["swap_2"].Texture, sdl.Point{X: half.X, Y: renderer.WindowContext.WindowSize.Y - 136}, sdl.FPoint{X: 2, Y: 2})

	sdlutils.RenderTexture(renderer.Renderer, layer.Data.(KurinRendererLayerHUDData).Icons["pda"].Texture, KurinHUDElementPDA.GetPosition(renderer.WindowContext.WindowSize), sdl.FPoint{X: 2, Y: 2})
	if KurinHUDElementPDA.Hovered {
		layer.Data.(KurinRendererLayerHUDData).Icons["selector"].Texture.Texture.SetColorMod(0, 255, 0)
		sdlutils.RenderTexture(renderer.Renderer, layer.Data.(KurinRendererLayerHUDData).Icons["selector"].Texture, KurinHUDElementPDA.GetPosition(renderer.WindowContext.WindowSize), sdl.FPoint{X: 2, Y: 2})
	}

	return nil
}
