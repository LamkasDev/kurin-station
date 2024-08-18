package gameplay

type KurinNarratorObjectiveRequirement struct {
	Type       string
	IsDone     KurinNarratorObjectiveRequirementIsDone
	EncodeData KurinNarratorObjectiveRequirementEncodeData
	DecodeData KurinNarratorObjectiveRequirementDecodeData
	Data       interface{}
}

type (
	KurinNarratorObjectiveRequirementIsDone     func(requirement *KurinNarratorObjectiveRequirement) bool
	KurinNarratorObjectiveRequirementEncodeData func(requirement *KurinNarratorObjectiveRequirement) []byte
	KurinNarratorObjectiveRequirementDecodeData func(requirement *KurinNarratorObjectiveRequirement, data []byte)
)
