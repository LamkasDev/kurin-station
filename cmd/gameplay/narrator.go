package gameplay

import "slices"

type KurinNarrator struct {
	Objectives []*KurinNarratorObjective
}

func NewKurinNarrator() KurinNarrator {
	return KurinNarrator{
		Objectives: []*KurinNarratorObjective{},
	}
}

func AddKurinNarratorObjective(objective *KurinNarratorObjective) {
	KurinGameInstance.Narrator.Objectives = append(KurinGameInstance.Narrator.Objectives, objective)
	if len(KurinGameInstance.Narrator.Objectives) == 1 {
		StartKurinNarratorObjective(objective)
	}
}

func StartKurinNarratorObjective(objective *KurinNarratorObjective) {
	CreateKurinRunechatMessage(&KurinGameInstance.RunechatController, NewKurinRunechat(objective.Text))
}

func CompleteKurinNarratorObjective() {
	KurinGameInstance.Narrator.Objectives = slices.Delete(KurinGameInstance.Narrator.Objectives, 0, 1)
	if len(KurinGameInstance.Narrator.Objectives) > 0 {
		StartKurinNarratorObjective(KurinGameInstance.Narrator.Objectives[0])
	}
}

func ProcessKurinNarrator() {
	if KurinGameInstance.Ticks == 120 {
		AddKurinNarratorObjective(NewKurinNarratorObjectiveGettingStarted())
		AddKurinNarratorObjective(NewKurinNarratorObjectiveCleaningUp())
		AddKurinNarratorObjective(NewKurinNarratorObjectiveFreshStart())
		AddKurinNarratorObjective(NewKurinNarratorObjectiveToTheMoon())
	}
	if len(KurinGameInstance.Narrator.Objectives) > 0 {
		objective := KurinGameInstance.Narrator.Objectives[0]
		objective.Ticks++
		done := true
		for _, requirement := range objective.Requirements {
			if !requirement.IsDone(requirement) {
				done = false
				break
			}
		}
		if done {
			CompleteKurinNarratorObjective()
		}
	}
}

func KurinNarratorOnCreateObject(object *KurinObject) {
	for _, objective := range KurinGameInstance.Narrator.Objectives {
		for _, requirement := range objective.Requirements {
			switch data := requirement.Data.(type) {
			case KurinNarratorObjectiveRequirementDataCreate:
				if data.ObjectType != object.Type {
					continue
				}
				data.Progress++
				requirement.Data = data
			}
		}
	}
}

func KurinNarratorOnDestroyObject(object *KurinObject) {
	for _, objective := range KurinGameInstance.Narrator.Objectives {
		for _, requirement := range objective.Requirements {
			switch data := requirement.Data.(type) {
			case KurinNarratorObjectiveRequirementDataDestroy:
				if data.ObjectType != object.Type {
					continue
				}
				data.Progress++
				requirement.Data = data
			}
		}
	}
}
