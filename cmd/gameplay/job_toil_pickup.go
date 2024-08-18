package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
)

type KurinJobToilPickupData struct {
	ItemType  string
	ItemCount uint16

	Item     *KurinItem
	GotoToil *KurinJobToil
}

func NewKurinJobToilPickup(itemType string, itemCount uint16) *KurinJobToil {
	toil := NewKurinJobToilRaw[KurinJobToilPickupData]("pickup")
	toil.Start = StartKurinJobToilPickup
	toil.Process = ProcessKurinJobToilPickup
	toil.End = EndKurinJobToilPickup
	toil.Data = &KurinJobToilPickupData{
		ItemType:  itemType,
		ItemCount: itemCount,
	}

	return toil
}

func StartKurinJobToilPickup(driver *KurinJobDriver, toil *KurinJobToil) KurinJobToilStatus {
	data := toil.Data.(*KurinJobToilPickupData)
	if IsKurinJobToilPickupComplete(driver, toil) {
		return KurinJobToilStatusComplete
	}
	data.Item = FindClosestItemOfType(&GameInstance.Map, driver.Character.Position, data.ItemType, true)
	if data.Item == nil {
		return KurinJobToilStatusFailed
	}
	ReserveKurinItem(data.Item)
	data.GotoToil = NewKurinJobToilGoto(sdlutils.FVector3ToVector3(data.Item.Transform.Position))

	return data.GotoToil.Start(driver, data.GotoToil)
}

func ProcessKurinJobToilPickup(driver *KurinJobDriver, toil *KurinJobToil) KurinJobToilStatus {
	data := toil.Data.(*KurinJobToilPickupData)
	if data.Item.Character != nil {
		return KurinJobToilStatusFailed
	}
	switch data.GotoToil.Process(driver, data.GotoToil) {
	case KurinJobToilStatusComplete:
		if !TransferKurinItemToCharacter(data.Item, driver.Character) {
			return KurinJobToilStatusFailed
		}
		if !IsKurinJobToilPickupComplete(driver, toil) {
			return StartKurinJobToilPickup(driver, toil)
		}

		return KurinJobToilStatusComplete
	case KurinJobToilStatusFailed:
		return KurinJobToilStatusFailed
	}

	return KurinJobToilStatusWorking
}

func IsKurinJobToilPickupComplete(driver *KurinJobDriver, toil *KurinJobToil) bool {
	data := toil.Data.(*KurinJobToilPickupData)
	item := FindItemInInventory(&driver.Character.Inventory, data.ItemType)
	if item != nil && item.Count >= data.ItemCount {
		return true
	}

	return false
}

func EndKurinJobToilPickup(driver *KurinJobDriver, toil *KurinJobToil) {
	data := toil.Data.(*KurinJobToilPickupData)
	if data.Item != nil {
		UnreserveKurinItem(data.Item)
	}
}
