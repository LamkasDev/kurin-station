package gameplay

type ObjectiveRequirementDataCredits struct {
	Count    uint32
	Progress uint32
}

func NewObjectiveRequirementTemplateCredits() *ObjectiveRequirementTemplate {
	template := NewObjectiveRequirementTemplate[*ObjectiveRequirementDataCredits]("credits")
	template.IsDone = func(requirement *ObjectiveRequirement) bool {
		data := requirement.Data.(*ObjectiveRequirementDataCredits)
		return data.Progress >= data.Count
	}

	return template
}
