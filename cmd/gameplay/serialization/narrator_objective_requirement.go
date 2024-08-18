package serialization

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
)

type KurinNarratorObjectiveRequirementData struct {
	Type string
	Data []byte
}

func EncodeKurinNarratorObjectiveRequirement(requirement *gameplay.KurinNarratorObjectiveRequirement) KurinNarratorObjectiveRequirementData {
	return KurinNarratorObjectiveRequirementData{
		Type: requirement.Type,
		Data: requirement.EncodeData(requirement),
	}
}

func DecodeKurinNarratorObjectiveRequirement(data KurinNarratorObjectiveRequirementData) *gameplay.KurinNarratorObjectiveRequirement {
	requirement := gameplay.NewKurinNarratorObjectiveRequirement(data.Type)
	requirement.DecodeData(requirement, data.Data)

	return requirement
}
