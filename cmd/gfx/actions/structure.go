package actions

import (
	"slices"
	"sort"
	"strings"

	"github.com/LamkasDev/kurin/cmd/gfx/structure"
	"github.com/LamkasDev/kurin/cmd/gfx/turf"
	"golang.org/x/exp/maps"
	"robpike.io/filter"
)

func GetMenuGraphics(data *RendererLayerActionsData) []interface{} {
	var graphics []interface{}
	graphics = GetMenuStructureGraphics(data)
	graphics = slices.Concat(graphics, GetMenuTurfGraphics(data))

	return graphics
}

func GetMenuStructureGraphics(data *RendererLayerActionsData) []interface{} {
	structures := maps.Values(data.ObjectLayer.Data.(*structure.KurinRendererLayerObjectData).Structures)
	sort.Slice(structures, func(i, j int) bool {
		return structures[i].Template.Id < structures[j].Template.Id
	})
	structures = filter.Choose(structures, func(item *structure.KurinStructureGraphic) bool {
		return item.Blueprint != nil && strings.Contains(item.Template.Name, data.Input)
	}).([]*structure.KurinStructureGraphic)

	newStructures := []interface{}{}
	for _, structure := range structures {
		newStructures = append(newStructures, structure)
	}

	return newStructures
}

func GetMenuTurfGraphics(data *RendererLayerActionsData) []interface{} {
	turfs := maps.Values(data.TurfLayer.Data.(*turf.KurinRendererLayerTileData).Turfs)
	sort.Slice(turfs, func(i, j int) bool {
		return turfs[i].Template.Id < turfs[j].Template.Id
	})
	turfs = filter.Choose(turfs, func(item *turf.KurinTurfGraphic) bool {
		return item.Blueprint != nil && strings.Contains(item.Template.Name, data.Input)
	}).([]*turf.KurinTurfGraphic)

	newTurfs := []interface{}{}
	for _, turf := range turfs {
		newTurfs = append(newTurfs, turf)
	}

	return newTurfs
}
