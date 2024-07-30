package life

import (
	"github.com/LamkasDev/kurin/cmd/event"
	eventActions "github.com/LamkasDev/kurin/cmd/event/actions"
	"github.com/LamkasDev/kurin/cmd/event/base"
	"github.com/LamkasDev/kurin/cmd/event/camera"
	eventContext "github.com/LamkasDev/kurin/cmd/event/context"
	eventDebug "github.com/LamkasDev/kurin/cmd/event/debug"
	"github.com/LamkasDev/kurin/cmd/event/force"
	eventHud "github.com/LamkasDev/kurin/cmd/event/hud"
	"github.com/LamkasDev/kurin/cmd/event/interaction"
	"github.com/LamkasDev/kurin/cmd/event/keybinds"
	"github.com/LamkasDev/kurin/cmd/event/movement"
	eventTool "github.com/LamkasDev/kurin/cmd/event/tool"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/actions"
	"github.com/LamkasDev/kurin/cmd/gfx/animation"
	"github.com/LamkasDev/kurin/cmd/gfx/context"
	"github.com/LamkasDev/kurin/cmd/gfx/debug"
	"github.com/LamkasDev/kurin/cmd/gfx/hud"
	"github.com/LamkasDev/kurin/cmd/gfx/item"
	"github.com/LamkasDev/kurin/cmd/gfx/job"
	"github.com/LamkasDev/kurin/cmd/gfx/particle"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/LamkasDev/kurin/cmd/gfx/runechat"
	"github.com/LamkasDev/kurin/cmd/gfx/species"
	"github.com/LamkasDev/kurin/cmd/gfx/structure"
	"github.com/LamkasDev/kurin/cmd/gfx/tool"
	"github.com/LamkasDev/kurin/cmd/gfx/tooltip"
	"github.com/LamkasDev/kurin/cmd/gfx/turf"
	"github.com/LamkasDev/kurin/cmd/sound"
	"github.com/LamkasDev/kurin/cmd/sound/ambient"
	"github.com/LamkasDev/kurin/cmd/sound/voice"
)

type KurinInstance struct {
	Game         gameplay.KurinGame
	Renderer     *gfx.KurinRenderer
	EventManager event.KurinEventManager
	SoundManager sound.KurinSoundManager
}

func NewKurinInstance() (KurinInstance, *error) {
	instance := KurinInstance{
		Game: gameplay.NewKurinGame(),
	}

	var err *error
	if instance.Renderer, err = gfx.NewKurinRenderer(); err != nil {
		return instance, err
	}
	instance.Renderer.Layers = append(instance.Renderer.Layers, turf.NewKurinRendererLayerTile())
	objectLayer := structure.NewKurinRendererLayerObject()
	instance.Renderer.Layers = append(instance.Renderer.Layers, objectLayer)
	itemLayer := item.NewKurinRendererLayerItem()
	instance.Renderer.Layers = append(instance.Renderer.Layers, itemLayer)
	instance.Renderer.Layers = append(instance.Renderer.Layers, job.NewKurinRendererLayerJob(objectLayer))
	instance.Renderer.Layers = append(instance.Renderer.Layers, species.NewKurinRendererLayerCharacter(itemLayer))
	instance.Renderer.Layers = append(instance.Renderer.Layers, particle.NewKurinRendererLayerParticle())
	instance.Renderer.Layers = append(instance.Renderer.Layers, runechat.NewKurinRendererLayerRunechat())
	toolLayer := tool.NewKurinRendererLayerTool(objectLayer)
	instance.Renderer.Layers = append(instance.Renderer.Layers, toolLayer)
	actionsLayer := actions.NewKurinRendererLayerActions(objectLayer)
	instance.Renderer.Layers = append(instance.Renderer.Layers, actionsLayer)
	instance.Renderer.Layers = append(instance.Renderer.Layers, tooltip.NewKurinRendererLayerTooltip())
	instance.Renderer.Layers = append(instance.Renderer.Layers, hud.NewKurinRendererLayerHUD(itemLayer))
	contextLayer := context.NewKurinRendererLayerContext()
	instance.Renderer.Layers = append(instance.Renderer.Layers, contextLayer)
	debugLayer := debug.NewKurinRendererLayerDebug()
	instance.Renderer.Layers = append(instance.Renderer.Layers, debugLayer)
	instance.Renderer.Layers = append(instance.Renderer.Layers, animation.NewKurinRendererLayerAnimation())
	if err := gfx.LoadKurinRenderer(instance.Renderer); err != nil {
		return instance, err
	}

	if instance.EventManager, err = event.NewKurinEventManager(instance.Renderer); err != nil {
		return instance, err
	}
	instance.EventManager.Layers = append(instance.EventManager.Layers, base.NewKurinEventLayerBase())
	instance.EventManager.Layers = append(instance.EventManager.Layers, eventContext.NewKurinEventLayerContext(contextLayer))
	instance.EventManager.Layers = append(instance.EventManager.Layers, eventTool.NewKurinEventLayerTool(toolLayer))
	instance.EventManager.Layers = append(instance.EventManager.Layers, eventActions.NewKurinEventLayerActions(actionsLayer, toolLayer))
	instance.EventManager.Layers = append(instance.EventManager.Layers, keybinds.NewKurinEventLayerKeybinds())
	instance.EventManager.Layers = append(instance.EventManager.Layers, movement.NewKurinEventLayerMovement())
	instance.EventManager.Layers = append(instance.EventManager.Layers, force.NewKurinEventLayerForce())
	instance.EventManager.Layers = append(instance.EventManager.Layers, camera.NewKurinEventLayerCamera())
	instance.EventManager.Layers = append(instance.EventManager.Layers, eventHud.NewKurinEventLayerHUD())
	instance.EventManager.Layers = append(instance.EventManager.Layers, interaction.NewKurinEventLayerInteraction(itemLayer))
	instance.EventManager.Layers = append(instance.EventManager.Layers, eventDebug.NewKurinEventLayerDebug(debugLayer))
	if err := event.LoadKurinEventManager(&instance.EventManager); err != nil {
		return instance, err
	}

	if instance.SoundManager, err = sound.NewKurinSoundManager(); err != nil {
		return instance, err
	}
	instance.SoundManager.Layers = append(instance.SoundManager.Layers, ambient.NewKurinSoundLayerAmbient())
	instance.SoundManager.Layers = append(instance.SoundManager.Layers, voice.NewKurinSoundLayerVoice())
	if err := sound.LoadKurinSoundManager(&instance.SoundManager); err != nil {
		return instance, err
	}

	return instance, nil
}

func RunKurinInstance(instance *KurinInstance) *error {
	if err := event.ProcessKurinEventManager(&instance.EventManager, &instance.Game); err != nil {
		return err
	}
	instance.Renderer.RendererContext.CameraTileSize = render.GetCameraTileSize(instance.Renderer)
	instance.Renderer.RendererContext.CameraOffset = render.GetCameraOffset(instance.Renderer)

	if err := sound.ProcessKurinSoundManager(&instance.SoundManager, &instance.Game); err != nil {
		return err
	}

	if err := gfx.ClearKurinRenderer(instance.Renderer); err != nil {
		return err
	}
	if err := gfx.RenderKurinRenderer(instance.Renderer, &instance.Game); err != nil {
		return err
	}
	gfx.PresentKurinRenderer(instance.Renderer)

	return nil
}

func FreeKurinInstance(instance *KurinInstance) *error {
	if err := gfx.FreeKurinRenderer(instance.Renderer); err != nil {
		return err
	}
	if err := event.FreeKurinEventManager(&instance.EventManager); err != nil {
		return err
	}

	return sound.FreeKurinSoundManager(&instance.SoundManager)
}
