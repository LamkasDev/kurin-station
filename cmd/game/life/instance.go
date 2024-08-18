package life

import (
	"github.com/LamkasDev/kurin/cmd/event"
	eventActions "github.com/LamkasDev/kurin/cmd/event/actions"
	"github.com/LamkasDev/kurin/cmd/event/base"
	"github.com/LamkasDev/kurin/cmd/event/camera"
	eventContext "github.com/LamkasDev/kurin/cmd/event/context"
	eventDebug "github.com/LamkasDev/kurin/cmd/event/debug"
	eventDialog "github.com/LamkasDev/kurin/cmd/event/dialog"
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
	"github.com/LamkasDev/kurin/cmd/gfx/background"
	"github.com/LamkasDev/kurin/cmd/gfx/context"
	"github.com/LamkasDev/kurin/cmd/gfx/debug"
	"github.com/LamkasDev/kurin/cmd/gfx/dialog"
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
	"github.com/LamkasDev/kurin/cmd/sound/music"
	"github.com/LamkasDev/kurin/cmd/sound/voice"
)

type KurinInstance struct {
	Game         gameplay.KurinGame
	EventManager event.EventManager
	SoundManager sound.KurinSoundManager
}

func NewKurinInstance() (KurinInstance, error) {
	instance := KurinInstance{
		Game: gameplay.NewKurinGame(),
	}

	var err error
	if err = gfx.InitializeKurinRenderer(); err != nil {
		return instance, err
	}

	turfLayer := turf.NewKurinRendererLayerTile()
	objectLayer := structure.NewKurinRendererLayerObject()
	itemLayer := item.NewKurinRendererLayerItem()
	toolLayer := tool.NewKurinRendererLayerTool(turfLayer, objectLayer)
	actionsLayer := actions.NewKurinRendererLayerActions(turfLayer, objectLayer, itemLayer)
	hudLayer := hud.NewKurinRendererLayerHUD(itemLayer)
	contextLayer := context.NewKurinRendererLayerContext()
	debugLayer := debug.NewKurinRendererLayerDebug()
	dialogLayer := dialog.NewKurinRendererLayerDialog(itemLayer)

	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, background.NewKurinRendererLayerBackground())
	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, turfLayer)
	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, objectLayer)
	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, itemLayer)
	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, job.NewKurinRendererLayerJob(objectLayer))
	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, species.NewKurinRendererLayerCharacter(itemLayer))
	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, particle.NewKurinRendererLayerParticle())
	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, runechat.NewKurinRendererLayerRunechat())
	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, toolLayer)
	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, actionsLayer)
	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, tooltip.NewKurinRendererLayerTooltip())
	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, hudLayer)
	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, contextLayer)
	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, debugLayer)
	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, animation.NewKurinRendererLayerAnimation())
	gfx.RendererInstance.Layers = append(gfx.RendererInstance.Layers, dialogLayer)

	if err := gfx.LoadKurinRenderer(); err != nil {
		return instance, err
	}

	if err = event.InitializeEventManager(); err != nil {
		return instance, err
	}

	instance.EventManager.Layers = append(instance.EventManager.Layers, base.NewKurinEventLayerBase())
	instance.EventManager.Layers = append(instance.EventManager.Layers, eventDialog.NewKurinEventLayerDialog(dialogLayer))
	instance.EventManager.Layers = append(instance.EventManager.Layers, eventContext.NewKurinEventLayerContext(contextLayer))
	instance.EventManager.Layers = append(instance.EventManager.Layers, eventTool.NewKurinEventLayerTool(toolLayer))
	instance.EventManager.Layers = append(instance.EventManager.Layers, eventActions.NewEventLayerActions(actionsLayer, toolLayer))
	instance.EventManager.Layers = append(instance.EventManager.Layers, keybinds.NewKurinEventLayerKeybinds())
	instance.EventManager.Layers = append(instance.EventManager.Layers, movement.NewKurinEventLayerMovement())
	instance.EventManager.Layers = append(instance.EventManager.Layers, force.NewKurinEventLayerForce())
	instance.EventManager.Layers = append(instance.EventManager.Layers, camera.NewKurinEventLayerCamera())
	instance.EventManager.Layers = append(instance.EventManager.Layers, eventHud.NewKurinEventLayerHUD(hudLayer, itemLayer))
	instance.EventManager.Layers = append(instance.EventManager.Layers, interaction.NewKurinEventLayerInteraction(itemLayer))
	instance.EventManager.Layers = append(instance.EventManager.Layers, eventDebug.NewKurinEventLayerDebug(debugLayer))

	if err := event.LoadEventManager(); err != nil {
		return instance, err
	}

	if instance.SoundManager, err = sound.NewKurinSoundManager(); err != nil {
		return instance, err
	}

	instance.SoundManager.Layers = append(instance.SoundManager.Layers, ambient.NewKurinSoundLayerAmbient())
	instance.SoundManager.Layers = append(instance.SoundManager.Layers, voice.NewKurinSoundLayerVoice())
	instance.SoundManager.Layers = append(instance.SoundManager.Layers, music.NewKurinSoundLayerMusic())

	if err := sound.LoadKurinSoundManager(&instance.SoundManager); err != nil {
		return instance, err
	}

	return instance, nil
}

func RunKurinInstance(instance *KurinInstance) error {
	if err := event.ProcessEventManager(); err != nil {
		return err
	}

	gfx.RendererInstance.Context.CameraTileSize = render.GetCameraTileSize()
	gfx.RendererInstance.Context.CameraOffset = render.GetCameraOffset()

	if err := sound.ProcessKurinSoundManager(&instance.SoundManager); err != nil {
		return err
	}

	if err := gfx.ClearKurinRenderer(); err != nil {
		return err
	}

	if err := gfx.RenderKurinRenderer(); err != nil {
		return err
	}

	gfx.PresentKurinRenderer()

	return nil
}

func FreeKurinInstance(instance *KurinInstance) error {
	if err := gfx.FreeKurinRenderer(); err != nil {
		return err
	}

	if err := event.FreeEventManager(); err != nil {
		return err
	}

	return sound.FreeKurinSoundManager(&instance.SoundManager)
}
