package gameplay

type ObjectiveRequirement struct {
	Type     string
	Template *ObjectiveRequirementTemplate
	Data     interface{}
}
