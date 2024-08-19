package gameplay

import "slices"

type Narrator struct {
	Objectives []*Objective
}

func NewNarrator() *Narrator {
	return &Narrator{
		Objectives: []*Objective{},
	}
}

func AddNarratorObjective(narrator *Narrator, objective *Objective) {
	narrator.Objectives = append(narrator.Objectives, objective)
	if len(narrator.Objectives) == 1 {
		StartNarratorObjective(narrator, objective)
	}
}

func StartNarratorObjective(narrator *Narrator, objective *Objective) {
	CreateRunechatMessage(&GameInstance.RunechatController, NewRunechat(objective.Text))
}

func CompleteNarratorObjective(narrator *Narrator) {
	narrator.Objectives = slices.Delete(narrator.Objectives, 0, 1)
	if len(narrator.Objectives) > 0 {
		StartNarratorObjective(narrator, narrator.Objectives[0])
	}
}

func ProcessNarrator() {
	if GameInstance.Ticks == 120 {
		AddNarratorObjective(GameInstance.Narrator, NewObjectiveGettingStarted())
		AddNarratorObjective(GameInstance.Narrator, NewObjectiveCleaningUp())
		AddNarratorObjective(GameInstance.Narrator, NewObjectiveFreshStart())
		AddNarratorObjective(GameInstance.Narrator, NewObjectiveToTheMoon())
	}
	if len(GameInstance.Narrator.Objectives) > 0 {
		objective := GameInstance.Narrator.Objectives[0]
		objective.Ticks++
		done := true
		for _, requirement := range objective.Requirements {
			if !requirement.Template.IsDone(requirement) {
				done = false
				break
			}
		}
		if done {
			CompleteNarratorObjective(GameInstance.Narrator)
		}
	}
}

func NarratorOnCreateObject(object *Object) {
	for _, objective := range GameInstance.Narrator.Objectives {
		for _, requirement := range objective.Requirements {
			switch data := requirement.Data.(type) {
			case *ObjectiveRequirementDataCreate:
				if data.ObjectType != object.Type {
					continue
				}
				data.Progress++
				requirement.Data = data
			}
		}
	}
}

func NarratorOnDestroyObject(object *Object) {
	for _, objective := range GameInstance.Narrator.Objectives {
		for _, requirement := range objective.Requirements {
			switch data := requirement.Data.(type) {
			case *ObjectiveRequirementDataDestroy:
				if data.ObjectType != object.Type {
					continue
				}
				data.Progress++
				requirement.Data = data
			}
		}
	}
}
