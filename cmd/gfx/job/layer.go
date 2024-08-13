package job

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/structure"
)

type KurinRendererLayerJobData struct {
	ObjectLayer *gfx.KurinRendererLayer
}

func NewKurinRendererLayerJob(objectLayer *gfx.KurinRendererLayer) *gfx.KurinRendererLayer {
	return &gfx.KurinRendererLayer{
		Load:   LoadKurinRendererLayerJob,
		Render: RenderKurinRendererLayerJob,
		Data: KurinRendererLayerJobData{
			ObjectLayer: objectLayer,
		},
	}
}

func LoadKurinRendererLayerJob(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer) error {
	return nil
}

func RenderKurinRendererLayerJob(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer) error {
	for _, job := range gameplay.KurinGameInstance.JobController.Jobs.Rep {
		if job == nil {
			continue
		}
		if err := RenderKurinJob(renderer, layer, job.(*gameplay.KurinJobDriver)); err != nil {
			return err
		}
	}

	for _, character := range gameplay.KurinGameInstance.Characters {
		if character.JobTracker.Job == nil {
			continue
		}
		if err := RenderKurinJob(renderer, layer, character.JobTracker.Job); err != nil {
			return err
		}
	}

	return nil
}

func RenderKurinJob(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer, job *gameplay.KurinJobDriver) error {
	data := layer.Data.(KurinRendererLayerJobData)
	switch val := job.Data.(type) {
	case gameplay.KurinJobDriverBuildData:
		if err := structure.RenderKurinObjectBlueprint(renderer, data.ObjectLayer, val.Prefab, sdlutils.White); err != nil {
			return err
		}
	}

	return nil
}
