package gameplay

type ObjectiveRequirementDataDestroy struct {
	ObjectType string
	Count      uint8
	Progress   uint8
}

func NewObjectiveRequirementTemplateDestroy() *ObjectiveRequirementTemplate {
	template := NewObjectiveRequirementTemplate[*ObjectiveRequirementDataDestroy]("destroy")
	template.IsDone = func(requirement *ObjectiveRequirement) bool {
		data := requirement.Data.(*ObjectiveRequirementDataDestroy)
		return data.Progress >= data.Count
	}

	return template
}
