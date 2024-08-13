package hud

import (
	"fmt"
	"math"

	"github.com/LamkasDev/kurin/cmd/common/mathutils"
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/item"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinRendererLayerHUDData struct {
	Icons     map[string]*KurinHUDGraphic
	ItemLayer *gfx.KurinRendererLayer
	HoveredItem *gameplay.KurinItem
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

func LoadKurinRendererLayerHUD(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer) error {
	return LoadKurinRendererLayerHUDArray(renderer, layer, []string{"hand_l", "lhandactive", "hand_r", "rhandactive", "act_equip", "swap_1", "swap_2", "pda", "selector", "template", "template_active", "credit", "objective_window", "cat"})
}

func LoadKurinRendererLayerHUDArray(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer, textures []string) error {
	var err error
	for _, texture := range textures {
		if layer.Data.(KurinRendererLayerHUDData).Icons[texture], err = NewKurinHUDGraphic(renderer, texture); err != nil {
			return err
		}
	}

	return nil
}

func RenderKurinRendererLayerHUD(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer) error {
	if renderer.Context.CameraMode != gfx.KurinRendererCameraModeCharacter {
		return nil
	}
	half := gfx.GetHalfWindowSize(&renderer.Context)
	data := layer.Data.(KurinRendererLayerHUDData)
	itemData := data.ItemLayer.Data.(item.KurinRendererLayerItemData)

	sdlutils.RenderTexture(renderer.Renderer, data.Icons["hand_l"].Texture, KurinHUDElementHandLeft.GetPosition(renderer.Context.WindowSize), sdl.FPoint{X: 2, Y: 2})
	if gameplay.KurinGameInstance.SelectedCharacter.ActiveHand == gameplay.KurinHandLeft {
		sdlutils.RenderTexture(renderer.Renderer, data.Icons["lhandactive"].Texture, sdl.Point{X: half.X, Y: renderer.Context.WindowSize.Y - 72}, sdl.FPoint{X: 2, Y: 2})
	}
	lhand := gameplay.KurinGameInstance.SelectedCharacter.Inventory.Hands[gameplay.KurinHandLeft]
	if lhand != nil {
		graphic := itemData.Items[lhand.Type]
		if graphic.Outline != nil && data.HoveredItem == lhand {
			sdlutils.RenderTexture(renderer.Renderer, *graphic.Outline, KurinHUDElementHandLeft.GetPosition(renderer.Context.WindowSize), sdl.FPoint{X: 2, Y: 2})
		}
		for _, i := range lhand.GetTextures(lhand) {
			sdlutils.RenderTexture(renderer.Renderer, graphic.Textures[i].Base, KurinHUDElementHandLeft.GetPosition(renderer.Context.WindowSize), sdl.FPoint{X: 2, Y: 2})
		}
	}

	sdlutils.RenderTexture(renderer.Renderer, data.Icons["hand_r"].Texture, KurinHUDElementHandRight.GetPosition(renderer.Context.WindowSize), sdl.FPoint{X: 2, Y: 2})
	if gameplay.KurinGameInstance.SelectedCharacter.ActiveHand == gameplay.KurinHandRight {
		sdlutils.RenderTexture(renderer.Renderer, data.Icons["rhandactive"].Texture, sdl.Point{X: half.X - 64, Y: renderer.Context.WindowSize.Y - 72}, sdl.FPoint{X: 2, Y: 2})
	}
	rhand := gameplay.KurinGameInstance.SelectedCharacter.Inventory.Hands[gameplay.KurinHandRight]
	if rhand != nil {
		graphic := itemData.Items[rhand.Type]
		if graphic.Outline != nil && data.HoveredItem == rhand {
			sdlutils.RenderTexture(renderer.Renderer, *graphic.Outline, KurinHUDElementHandRight.GetPosition(renderer.Context.WindowSize), sdl.FPoint{X: 2, Y: 2})
		}
		for _, i := range rhand.GetTextures(rhand) {
			sdlutils.RenderTexture(renderer.Renderer, graphic.Textures[i].Base, KurinHUDElementHandRight.GetPosition(renderer.Context.WindowSize), sdl.FPoint{X: 2, Y: 2})
		}
	}

	sdlutils.RenderTexture(renderer.Renderer, data.Icons["act_equip"].Texture, sdl.Point{X: half.X - 64, Y: renderer.Context.WindowSize.Y - 136}, sdl.FPoint{X: 2, Y: 2})
	sdlutils.RenderTexture(renderer.Renderer, data.Icons["swap_1"].Texture, sdl.Point{X: half.X - 64, Y: renderer.Context.WindowSize.Y - 136}, sdl.FPoint{X: 2, Y: 2})
	sdlutils.RenderTexture(renderer.Renderer, data.Icons["swap_2"].Texture, sdl.Point{X: half.X, Y: renderer.Context.WindowSize.Y - 136}, sdl.FPoint{X: 2, Y: 2})

	sdlutils.RenderTexture(renderer.Renderer, data.Icons["pda"].Texture, KurinHUDElementPDA.GetPosition(renderer.Context.WindowSize), sdl.FPoint{X: 2, Y: 2})
	if KurinHUDElementPDA.Hovered {
		layer.Data.(KurinRendererLayerHUDData).Icons["selector"].Texture.Texture.SetColorMod(0, 255, 0)
		sdlutils.RenderTexture(renderer.Renderer, data.Icons["selector"].Texture, KurinHUDElementPDA.GetPosition(renderer.Context.WindowSize), sdl.FPoint{X: 2, Y: 2})
	}

	goals := KurinHUDElementGoals.GetPosition(renderer.Context.WindowSize)
	sdlutils.RenderTexture(renderer.Renderer, data.Icons["objective_window"].Texture, goals, sdl.FPoint{X: 1, Y: 1})
	if len(gameplay.KurinGameInstance.Narrator.Objectives) > 0 {
		objective := gameplay.KurinGameInstance.Narrator.Objectives[0]
		text := objective.Text[0:mathutils.MinInt(int(math.Floor(float64(objective.Ticks)/6)), len(objective.Text))]
		sdlutils.RenderTexture(renderer.Renderer, data.Icons["cat"].Texture, sdl.Point{X: goals.X + 6, Y: goals.Y}, sdl.FPoint{X: 2, Y: 2})
		sdlutils.RenderLabel(renderer.Renderer, "hud.goals", renderer.Fonts.Default, sdlutils.White, text, sdl.Point{X: goals.X + 72, Y: goals.Y + 16}, sdl.FPoint{X: 1, Y: 1})
		for i, requirement := range objective.Requirements {
			pos :=  sdl.Point{X: goals.X + 72, Y: goals.Y + 36 + (int32(i) * 16)}
			text := "??"
			switch data := requirement.Data.(type) {
			case gameplay.KurinNarratorObjectiveRequirementDataCredits:
				text = fmt.Sprintf("Earn %d credits (%d/%d)", data.Count, gameplay.KurinGameInstance.Credits, data.Count)
			case gameplay.KurinNarratorObjectiveRequirementDataCreate:
				text = fmt.Sprintf("Create %d %s (%d/%d)", data.Count, data.ObjectType, data.Progress, data.Count)
			case gameplay.KurinNarratorObjectiveRequirementDataDestroy:
				text = fmt.Sprintf("Destroy %d %s (%d/%d)", data.Count, data.ObjectType, data.Progress, data.Count)
			}

			sdlutils.SetDrawColor(renderer.Renderer, sdlutils.White)
			if requirement.IsDone(requirement) {
				renderer.Renderer.FillRect(&sdl.Rect{X: pos.X, Y: pos.Y + 2, W: 10, H: 10})
			} else {
				renderer.Renderer.DrawRect(&sdl.Rect{X: pos.X, Y: pos.Y + 2, W: 10, H: 10})
			}
			sdlutils.RenderLabel(renderer.Renderer, fmt.Sprintf("hud.goals.%d", i), renderer.Fonts.DefaultSmall, sdlutils.White, text, sdl.Point{X: pos.X + 16, Y: pos.Y}, sdl.FPoint{X: 1, Y: 1})
		}
	}

	credit := KurinHUDElementCredit.GetPosition(renderer.Context.WindowSize)
	sdlutils.RenderTexture(renderer.Renderer, data.Icons["credit"].Texture, credit, sdl.FPoint{X: 2, Y: 2})
	sdlutils.RenderLabel(renderer.Renderer, "hud.credits", renderer.Fonts.Default, sdlutils.White, fmt.Sprint(gameplay.KurinGameInstance.Credits), sdl.Point{X: credit.X + 58, Y: credit.Y + 28}, sdl.FPoint{X: 1, Y: 1})

	return nil
}
