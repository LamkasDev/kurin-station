package gameplay

import (
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
)

type KurinJobToilGotoData struct {
	Target sdlutils.Vector3
	Path *KurinPath
}

func NewKurinJobToilGoto(target sdlutils.Vector3) *KurinJobToil {
	return &KurinJobToil{
		Data: KurinJobToilGotoData{
			Target: target,
		},
		Start:   StartKurinJobToilGoto,
		Process: ProcessKurinJobToilGoto,
	}
}

func StartKurinJobToilGoto(driver *KurinJobDriver, toil *KurinJobToil) {
	data := toil.Data.(KurinJobToilGotoData)
	data.Path = FindKurinPathAdjacent(&KurinGameInstance.Map.Pathfinding, driver.Character.Position, data.Target)
	toil.Data = data
}

func ProcessKurinJobToilGoto(driver *KurinJobDriver, toil *KurinJobToil) bool {
	data := toil.Data.(KurinJobToilGotoData)
	return FollowKurinPath(driver.Character, data.Path)
}
