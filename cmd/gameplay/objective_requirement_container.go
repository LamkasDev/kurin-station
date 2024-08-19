package gameplay

var ObjectiveRequirementContainer = map[string]*ObjectiveRequirementTemplate{}

func RegisterObjectiveRequirements() {
	ObjectiveRequirementContainer["create"] = NewObjectiveRequirementTemplateCreate()
	ObjectiveRequirementContainer["credits"] = NewObjectiveRequirementTemplateCredits()
	ObjectiveRequirementContainer["destroy"] = NewObjectiveRequirementTemplateDestroy()
}

func NewObjectiveRequirement(requirementType string, data interface{}) *ObjectiveRequirement {
	requirement := &ObjectiveRequirement{
		Type:     requirementType,
		Template: ObjectiveRequirementContainer[requirementType],
		Data:     data,
	}

	return requirement
}
