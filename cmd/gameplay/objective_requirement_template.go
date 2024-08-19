package gameplay

import "github.com/kelindar/binary"

type ObjectiveRequirementTemplate struct {
	Type       string
	IsDone     ObjectiveRequirementIsDone
	EncodeData ObjectiveRequirementEncodeData
	DecodeData ObjectiveRequirementDecodeData
}

type (
	ObjectiveRequirementIsDone     func(requirement *ObjectiveRequirement) bool
	ObjectiveRequirementEncodeData func(requirement *ObjectiveRequirement) []byte
	ObjectiveRequirementDecodeData func(requirement *ObjectiveRequirement, data []byte)
)

func NewObjectiveRequirementTemplate[D any](requirementType string) *ObjectiveRequirementTemplate {
	return &ObjectiveRequirementTemplate{
		Type: requirementType,
		IsDone: func(requirement *ObjectiveRequirement) bool {
			return false
		},
		EncodeData: func(requirement *ObjectiveRequirement) []byte {
			if requirement.Data == nil {
				return []byte{}
			}

			requirementData := requirement.Data.(D)
			data, _ := binary.Marshal(&requirementData)
			return data
		},
		DecodeData: func(requirement *ObjectiveRequirement, data []byte) {
			if len(data) == 0 {
				return
			}

			var requirementData D
			binary.Unmarshal(data, &requirementData)
			requirement.Data = requirementData
		},
	}
}
