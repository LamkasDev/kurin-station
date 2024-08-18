package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
)

type KurinJobToilGotoData struct {
	Target sdlutils.Vector3
	Path   *KurinPath
}

func NewKurinJobToilGoto(target sdlutils.Vector3) *KurinJobToil {
	toil := NewKurinJobToilRaw[KurinJobToilGotoData]("goto")
	toil.Start = StartKurinJobToilGoto
	toil.Process = ProcessKurinJobToilGoto
	toil.Data = &KurinJobToilGotoData{
		Target: target,
	}

	return toil
}

func StartKurinJobToilGoto(driver *KurinJobDriver, toil *KurinJobToil) KurinJobToilStatus {
	data := toil.Data.(*KurinJobToilGotoData)
	data.Path = FindKurinPathAdjacent(&GameInstance.Map.Pathfinding, driver.Character.Position, data.Target)
	if data.Path == nil {
		return KurinJobToilStatusFailed
	}

	return KurinJobToilStatusWorking
}

func ProcessKurinJobToilGoto(driver *KurinJobDriver, toil *KurinJobToil) KurinJobToilStatus {
	data := toil.Data.(*KurinJobToilGotoData)
	if FollowKurinPath(driver.Character, data.Path) {
		return KurinJobToilStatusComplete
	}

	return KurinJobToilStatusWorking
}
