package gameplay

import "github.com/kelindar/binary"

func NewKurinNarratorObjectiveRequirementRaw[D any](requirementType string) *KurinNarratorObjectiveRequirement {
	return &KurinNarratorObjectiveRequirement{
		Type: requirementType,
		IsDone: func(requirement *KurinNarratorObjectiveRequirement) bool {
			return false
		},
		EncodeData: func(requirement *KurinNarratorObjectiveRequirement) []byte {
			if requirement.Data == nil {
				return []byte{}
			}

			requirementData := requirement.Data.(D)
			data, _ := binary.Marshal(&requirementData)
			return data
		},
		DecodeData: func(requirement *KurinNarratorObjectiveRequirement, data []byte) {
			if len(data) == 0 {
				return
			}

			var requirementData D
			binary.Unmarshal(data, &requirementData)
			requirement.Data = requirementData
		},
	}
}

func NewKurinNarratorObjectiveRequirement(requirementType string) *KurinNarratorObjectiveRequirement {
	switch requirementType {
	case "create":
		return NewKurinNarratorObjectiveRequirementCreate("", 0)
	case "credits":
		return NewKurinNarratorObjectiveRequirementCredits(0)
	case "destroy":
		return NewKurinNarratorObjectiveRequirementDestroy("", 0)
	}

	return NewKurinNarratorObjectiveRequirementRaw[interface{}](requirementType)
}
