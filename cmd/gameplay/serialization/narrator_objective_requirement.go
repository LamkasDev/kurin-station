package serialization

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
)

type NarratorObjectiveRequirementData struct {
	Type string
	Data []byte
}

func EncodeNarratorObjectiveRequirement(requirement *gameplay.ObjectiveRequirement) NarratorObjectiveRequirementData {
	return NarratorObjectiveRequirementData{
		Type: requirement.Type,
		Data: requirement.Template.EncodeData(requirement),
	}
}

func DecodeNarratorObjectiveRequirement(data NarratorObjectiveRequirementData) *gameplay.ObjectiveRequirement {
	requirement := gameplay.NewObjectiveRequirement(data.Type, nil)
	requirement.Template.DecodeData(requirement, data.Data)

	return requirement
}
