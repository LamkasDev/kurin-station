package gameplay

import (
	"slices"

	"github.com/LamkasDev/kurin/cmd/common/mathutils"
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/gameplay/common"
	"github.com/arl/math32"
)

type ObjectLatheData struct {
	Energy    uint32
	MaxEnergy uint32
	Orders    []*LatheOrder
}

func NewObjectTemplateLathe() *ObjectTemplate {
	template := NewObjectTemplate[*ObjectLatheData]("lathe", false)
	template.Process = func(object *Object) {
		data := object.Data.(*ObjectLatheData)
		data.Energy = mathutils.MinUint32(data.Energy+1, data.MaxEnergy)
	}
	template.GetTexture = func(object *Object) int {
		data := object.Data.(*ObjectLatheData)
		if len(data.Orders) == 0 {
			return 0
		}
		if data.Orders[0].TicksLeft > 90 {
			if int(math32.Floor(float32(data.Orders[0].TicksLeft)/20))%2 == 0 {
				return 0
			}
			return 1
		}

		return int(math32.Floor((90 - float32(data.Orders[0].TicksLeft)) / 10))
	}
	template.OnInteraction = func(object *Object, item *Item) bool {
		OpenDialog(&DialogRequest{Type: "lathe", Data: &DialogLatheData{Lathe: object}})
		return true
	}
	template.GetDefaultData = func() interface{} {
		return &ObjectLatheData{
			MaxEnergy: 8000,
			Orders:    []*LatheOrder{},
		}
	}
	template.MaxHealth = 15

	return template
}

func CreateNewOrderAtLathe(lathe *Object, order *LatheOrder) bool {
	data := lathe.Data.(*ObjectLatheData)
	if data.Energy < order.Energy {
		PlaySound(&GameInstance.SoundController, "synth_no")
		return false
	}
	data.Energy -= order.Energy
	data.Orders = append(data.Orders, order)
	if len(data.Orders) == 1 {
		job := NewJobDriver("manufacture", lathe.Tile)
		job.Template.Initialize(job)
		PushJobToController(GameInstance.JobController[FactionPlayer], job)
	}
	PlaySound(&GameInstance.SoundController, "servostep")

	return true
}

func ProgressOrdersAtLathe(lathe *Object) bool {
	data := lathe.Data.(*ObjectLatheData)
	if len(data.Orders) == 0 {
		return true
	}
	data.Orders[0].TicksLeft--
	if data.Orders[0].TicksLeft == 0 {
		AddItemToMapRaw(
			NewItem(data.Orders[0].ItemType, data.Orders[0].ItemCount),
			GameInstance.Map,
			&sdlutils.Transform{
				Position: common.GetPositionInDirectionFVCenter(lathe.Tile.Position, common.DirectionEast),
				Rotation: 0,
			},
		)
		PlaySound(&GameInstance.SoundController, "servostep")
		data.Orders = slices.Delete(data.Orders, 0, 1)
	}

	return false
}
