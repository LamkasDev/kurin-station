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

type RendererLayerHUDData struct {
	Icons       map[string]*HUDGraphic
	ItemLayer   *gfx.RendererLayer
	HoveredItem *gameplay.Item
}

func NewRendererLayerHUD(itemLayer *gfx.RendererLayer) *gfx.RendererLayer {
	return &gfx.RendererLayer{
		Load:   LoadRendererLayerHUD,
		Render: RenderRendererLayerHUD,
		Data: &RendererLayerHUDData{
			Icons:     map[string]*HUDGraphic{},
			ItemLayer: itemLayer,
		},
	}
}

func LoadRendererLayerHUD(layer *gfx.RendererLayer) error {
	textures := []string{"hand_l", "lhandactive", "hand_r", "rhandactive", "act_equip", "swap_1", "swap_2", "pda", "selector", "template", "template_active", "credit", "objective_window", "cat", "healthdoll", "crit_overlay"}
	var err error
	for _, texture := range textures {
		if layer.Data.(*RendererLayerHUDData).Icons[texture], err = NewHUDGraphic(texture); err != nil {
			return err
		}
	}

	return nil
}

func RenderRendererLayerHUD(layer *gfx.RendererLayer) error {
	if gfx.RendererInstance.Context.CameraMode != gfx.RendererCameraModeCharacter {
		return nil
	}
	data := layer.Data.(*RendererLayerHUDData)
	itemData := data.ItemLayer.Data.(*item.RendererLayerItemData)
	windowSize := gfx.RendererInstance.Context.WindowSize

	if gameplay.GameInstance.SelectedCharacter.Health.Dead {
		sdlutils.RenderTexture(gfx.RendererInstance.Renderer, data.Icons["crit_overlay"].Texture, sdl.Point{}, gfx.RendererInstance.Context.WindowScale)
	}

	// Bottom Left - Hands
	lhandRect, _ := gfx.RenderUITexture(data.Icons["hand_l"].Texture, HUDElementHandLeft.GetPosition(windowSize), HUDElementHandLeft.Scale, HUDElementHandLeft.Anchor)
	if gameplay.GetActiveHand(gameplay.GameInstance.SelectedCharacter) == gameplay.HandLeft {
		gfx.RenderUITexture(data.Icons["lhandactive"].Texture, HUDElementHandLeft.GetPosition(windowSize), HUDElementHandLeft.Scale, HUDElementHandLeft.Anchor)
	}
	lhand := gameplay.GameInstance.SelectedCharacter.Data.(*gameplay.MobCharacterData).Inventory.Hands[gameplay.HandLeft]
	if lhand != nil {
		graphic := itemData.Items[lhand.Type]
		if graphic.Outline != nil && data.HoveredItem == lhand {
			gfx.RenderUITexture(graphic.Outline, HUDElementHandLeft.GetPosition(windowSize), HUDElementHandLeft.Scale, HUDElementHandLeft.Anchor)
		}
		for _, i := range lhand.Template.GetTextures(lhand) {
			gfx.RenderUITexture(graphic.Textures[i].Base, HUDElementHandLeft.GetPosition(windowSize), HUDElementHandLeft.Scale, HUDElementHandLeft.Anchor)
		}
	}

	rhandRect, _ := gfx.RenderUITexture(data.Icons["hand_r"].Texture, HUDElementHandRight.GetPosition(windowSize), HUDElementHandRight.Scale, HUDElementHandRight.Anchor)
	if gameplay.GetActiveHand(gameplay.GameInstance.SelectedCharacter) == gameplay.HandRight {
		gfx.RenderUITexture(data.Icons["rhandactive"].Texture, HUDElementHandRight.GetPosition(windowSize), HUDElementHandRight.Scale, HUDElementHandRight.Anchor)
	}
	rhand := gameplay.GameInstance.SelectedCharacter.Data.(*gameplay.MobCharacterData).Inventory.Hands[gameplay.HandRight]
	if rhand != nil {
		graphic := itemData.Items[rhand.Type]
		if graphic.Outline != nil && data.HoveredItem == rhand {
			gfx.RenderUITexture(graphic.Outline, HUDElementHandRight.GetPosition(windowSize), HUDElementHandRight.Scale, HUDElementHandRight.Anchor)
		}
		for _, i := range rhand.Template.GetTextures(rhand) {
			gfx.RenderUITexture(graphic.Textures[i].Base, HUDElementHandRight.GetPosition(windowSize), HUDElementHandRight.Scale, HUDElementHandRight.Anchor)
		}
	}

	gfx.RenderUITexture(data.Icons["act_equip"].Texture, sdl.Point{X: -int32(float32(lhandRect.W) / 2), Y: HUDElementHandLeft.GetPosition(windowSize).Y + lhandRect.H}, sdl.FPoint{X: 2, Y: 2}, gfx.UIAnchorBottomCenter)
	gfx.RenderUITexture(data.Icons["swap_1"].Texture, sdl.Point{X: -int32(float32(lhandRect.W) / 2), Y: HUDElementHandLeft.GetPosition(windowSize).Y + lhandRect.H}, sdl.FPoint{X: 2, Y: 2}, gfx.UIAnchorBottomCenter)
	gfx.RenderUITexture(data.Icons["swap_2"].Texture, sdl.Point{X: int32(float32(rhandRect.W) / 2), Y: HUDElementHandRight.GetPosition(windowSize).Y + rhandRect.H}, sdl.FPoint{X: 2, Y: 2}, gfx.UIAnchorBottomCenter)

	// Top Left - PDA
	gfx.RenderUITexture(data.Icons["pda"].Texture, HUDElementPDA.GetPosition(windowSize), HUDElementPDA.Scale, HUDElementPDA.Anchor)
	if HUDElementPDA.Hovered {
		layer.Data.(*RendererLayerHUDData).Icons["selector"].Texture.Texture.SetColorMod(0, 255, 0)
		gfx.RenderUITexture(data.Icons["selector"].Texture, HUDElementPDA.GetPosition(windowSize), HUDElementPDA.Scale, HUDElementPDA.Anchor)
	}

	// Top Right - Objectives
	goalsRect, _ := gfx.RenderUITexture(data.Icons["objective_window"].Texture, HUDElementGoals.GetPosition(windowSize), sdl.FPoint{X: 1, Y: 1}, gfx.UIAnchorTopLeft)
	if len(gameplay.GameInstance.Narrator.Objectives) > 0 {
		lineHeight := (goalsRect.H / 10)
		lineWidth := (goalsRect.W / 20)
		objective := gameplay.GameInstance.Narrator.Objectives[0]
		text := objective.Text[0:mathutils.MinInt(int(math.Floor(float64(objective.Ticks)/6)), len(objective.Text))]
		catRect, _ := gfx.RenderUITexture(data.Icons["cat"].Texture, sdl.Point{X: goalsRect.X + 6, Y: goalsRect.Y}, sdl.FPoint{X: 2, Y: 2}, gfx.UIAnchorTopLeft)
		_, titleRect := sdlutils.RenderLabel(gfx.RendererInstance.Renderer, "hud.goals", gfx.RendererInstance.Fonts.Default, sdlutils.White, text, sdl.Point{X: catRect.X + catRect.W + lineWidth, Y: catRect.Y + lineHeight*2}, gfx.RendererInstance.Context.WindowScale)
		for i, requirement := range objective.Requirements {
			pos := sdl.Point{X: titleRect.X, Y: titleRect.Y + (int32(i+3) * lineHeight)}
			text := "??"
			switch data := requirement.Data.(type) {
			case *gameplay.ObjectiveRequirementDataCredits:
				text = fmt.Sprintf("Earn %d credits (%d/%d)", data.Count, data.Progress, data.Count)
			case *gameplay.ObjectiveRequirementDataCreate:
				text = fmt.Sprintf("Create %d %s (%d/%d)", data.Count, data.ObjectType, data.Progress, data.Count)
			case *gameplay.ObjectiveRequirementDataDestroy:
				text = fmt.Sprintf("Destroy %d %s (%d/%d)", data.Count, data.ObjectType, data.Progress, data.Count)
			}

			sdlutils.SetDrawColor(gfx.RendererInstance.Renderer, sdlutils.White)
			if requirement.Template.IsDone(requirement) {
				gfx.RendererInstance.Renderer.FillRect(&sdl.Rect{X: pos.X, Y: pos.Y, W: lineHeight, H: lineHeight})
			} else {
				gfx.RendererInstance.Renderer.DrawRect(&sdl.Rect{X: pos.X, Y: pos.Y, W: lineHeight, H: lineHeight})
			}
			sdlutils.RenderLabel(gfx.RendererInstance.Renderer, fmt.Sprintf("hud.goals.%d", i), gfx.RendererInstance.Fonts.DefaultSmall, sdlutils.White, text, sdl.Point{X: pos.X + lineWidth, Y: pos.Y - int32(gfx.RendererInstance.Context.WindowScale.Y*3)}, gfx.RendererInstance.Context.WindowScale)
		}
	}

	creditRect, _ := gfx.RenderUITexture(data.Icons["credit"].Texture, sdl.Point{X: goalsRect.X + int32(gfx.RendererInstance.Context.WindowScale.X*8), Y: goalsRect.Y + goalsRect.H + int32(gfx.RendererInstance.Context.WindowScale.Y*16)}, sdl.FPoint{X: 2, Y: 2}, gfx.UIAnchorTopLeft)
	sdlutils.RenderLabel(gfx.RendererInstance.Renderer, "hud.credits", gfx.RendererInstance.Fonts.Default, sdlutils.White, fmt.Sprint(gameplay.GameInstance.Credits), sdl.Point{X: creditRect.X + creditRect.W + int32(gfx.RendererInstance.Context.WindowScale.X*8), Y: creditRect.Y}, gfx.RendererInstance.Context.WindowScale)

	// Center Right - Health
	gfx.RenderUITexture(data.Icons["healthdoll"].Texture, sdl.Point{X: int32(gfx.RendererInstance.Context.WindowScale.X * 8), Y: 0}, sdl.FPoint{X: 2, Y: 2}, gfx.UIAnchorCenterRight)

	return nil
}
