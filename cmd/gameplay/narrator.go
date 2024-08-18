package gameplay

import "slices"

type KurinNarrator struct {
	Objectives []*KurinNarratorObjective
}

func NewKurinNarrator() *KurinNarrator {
	return &KurinNarrator{
		Objectives: []*KurinNarratorObjective{},
	}
}

func AddKurinNarratorObjective(narrator *KurinNarrator, objective *KurinNarratorObjective) {
	narrator.Objectives = append(narrator.Objectives, objective)
	if len(narrator.Objectives) == 1 {
		StartKurinNarratorObjective(narrator, objective)
	}
}

func StartKurinNarratorObjective(narrator *KurinNarrator, objective *KurinNarratorObjective) {
	CreateKurinRunechatMessage(&GameInstance.RunechatController, NewKurinRunechat(objective.Text))
}

func CompleteKurinNarratorObjective(narrator *KurinNarrator) {
	narrator.Objectives = slices.Delete(narrator.Objectives, 0, 1)
	if len(narrator.Objectives) > 0 {
		StartKurinNarratorObjective(narrator, narrator.Objectives[0])
	}
}

func ProcessKurinNarrator() {
	if GameInstance.Ticks == 120 {
		AddKurinNarratorObjective(GameInstance.Narrator, NewKurinNarratorObjectiveGettingStarted())
		AddKurinNarratorObjective(GameInstance.Narrator, NewKurinNarratorObjectiveCleaningUp())
		AddKurinNarratorObjective(GameInstance.Narrator, NewKurinNarratorObjectiveFreshStart())
		AddKurinNarratorObjective(GameInstance.Narrator, NewKurinNarratorObjectiveToTheMoon())
	}
	if len(GameInstance.Narrator.Objectives) > 0 {
		objective := GameInstance.Narrator.Objectives[0]
		objective.Ticks++
		done := true
		for _, requirement := range objective.Requirements {
			if !requirement.IsDone(requirement) {
				done = false
				break
			}
		}
		if done {
			CompleteKurinNarratorObjective(GameInstance.Narrator)
		}
	}
}

func KurinNarratorOnCreateObject(object *KurinObject) {
	for _, objective := range GameInstance.Narrator.Objectives {
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
	for _, objective := range GameInstance.Narrator.Objectives {
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
