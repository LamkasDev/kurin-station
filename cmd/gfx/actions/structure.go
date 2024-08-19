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
	structures := maps.Values(data.ObjectLayer.Data.(*structure.RendererLayerObjectData).Structures)
	sort.Slice(structures, func(i, j int) bool {
		return structures[i].Template.Id < structures[j].Template.Id
	})
	structures = filter.Choose(structures, func(item *structure.StructureGraphic) bool {
		return item.Blueprint != nil && strings.Contains(strings.ToLower(item.Template.Name), strings.ToLower(data.Input))
	}).([]*structure.StructureGraphic)

	newStructures := []interface{}{}
	for _, structure := range structures {
		newStructures = append(newStructures, structure)
	}

	return newStructures
}

func GetMenuTurfGraphics(data *RendererLayerActionsData) []interface{} {
	turfs := maps.Values(data.TurfLayer.Data.(*turf.RendererLayerTileData).Turfs)
	sort.Slice(turfs, func(i, j int) bool {
		return turfs[i].Template.Id < turfs[j].Template.Id
	})
	turfs = filter.Choose(turfs, func(item *turf.TurfGraphic) bool {
		return item.Blueprint != nil && strings.Contains(strings.ToLower(item.Template.Name), strings.ToLower(data.Input))
	}).([]*turf.TurfGraphic)

	newTurfs := []interface{}{}
	for _, turf := range turfs {
		newTurfs = append(newTurfs, turf)
	}

	return newTurfs
}
