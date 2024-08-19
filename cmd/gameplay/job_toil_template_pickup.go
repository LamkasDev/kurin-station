package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
)

type JobToilPickupData struct {
	ItemType  string
	ItemCount uint16

	Item     *Item
	GotoToil *JobToil
}

func NewJobToilTemplatePickup() *JobToilTemplate {
	template := NewJobToilTemplate[JobToilPickupData]("pickup")
	template.Start = StartJobToilPickup
	template.Process = ProcessJobToilPickup
	template.End = EndJobToilPickup

	return template
}

func StartJobToilPickup(driver *JobDriver, toil *JobToil) JobToilStatus {
	data := toil.Data.(*JobToilPickupData)
	if IsJobToilPickupComplete(driver, toil) {
		return JobToilStatusComplete
	}
	data.Item = FindClosestItemOfType(&GameInstance.Map, driver.Character.Position, data.ItemType, true)
	if data.Item == nil {
		return JobToilStatusFailed
	}
	ReserveItem(data.Item)
	data.GotoToil = NewJobToil("goto", &JobToilGotoData{Target: sdlutils.FVector3ToVector3(data.Item.Transform.Position)})

	return data.GotoToil.Template.Start(driver, data.GotoToil)
}

func ProcessJobToilPickup(driver *JobDriver, toil *JobToil) JobToilStatus {
	data := toil.Data.(*JobToilPickupData)
	if data.Item.Character != nil {
		return JobToilStatusFailed
	}
	status := data.GotoToil.Template.Process(driver, data.GotoToil)
	switch status {
	case JobToilStatusComplete:
		if !TransferItemToCharacter(data.Item, driver.Character) {
			return JobToilStatusFailed
		}
		if !IsJobToilPickupComplete(driver, toil) {
			return StartJobToilPickup(driver, toil)
		}

		return JobToilStatusComplete
	case JobToilStatusFailed:
		return JobToilStatusFailed
	}

	return JobToilStatusWorking
}

func IsJobToilPickupComplete(driver *JobDriver, toil *JobToil) bool {
	data := toil.Data.(*JobToilPickupData)
	item := FindItemInInventory(&driver.Character.Inventory, data.ItemType)
	if item != nil && item.Count >= data.ItemCount {
		return true
	}

	return false
}

func EndJobToilPickup(driver *JobDriver, toil *JobToil) {
	data := toil.Data.(*JobToilPickupData)
	if data.Item != nil {
		UnreserveItem(data.Item)
	}
}
