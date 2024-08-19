package gameplay

type ObjectiveRequirementDataCreate struct {
	ObjectType string
	Count      uint8
	Progress   uint8
}

func NewObjectiveRequirementTemplateCreate() *ObjectiveRequirementTemplate {
	template := NewObjectiveRequirementTemplate[*ObjectiveRequirementDataCreate]("create")
	template.IsDone = func(requirement *ObjectiveRequirement) bool {
		data := requirement.Data.(*ObjectiveRequirementDataCreate)
		return data.Progress >= data.Count
	}

	return template
}
