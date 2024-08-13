package gameplay

type KurinNarratorObjectiveRequirement struct {
	Type   string
	IsDone KurinNarratorObjectiveRequirementIsDone
	Data   interface{}
}

type KurinNarratorObjectiveRequirementIsDone func(requirement *KurinNarratorObjectiveRequirement) bool

type KurinNarratorObjectiveRequirementDataCredits struct {
	Count uint32
}

func NewKurinNarratorObjectiveRequirementCredits(count uint32) *KurinNarratorObjectiveRequirement {
	return &KurinNarratorObjectiveRequirement{
		Type: "credits",
		IsDone: func(requirement *KurinNarratorObjectiveRequirement) bool {
			return KurinGameInstance.Credits >= requirement.Data.(KurinNarratorObjectiveRequirementDataCredits).Count
		},
		Data: KurinNarratorObjectiveRequirementDataCredits{
			Count: count,
		},
	}
}

type KurinNarratorObjectiveRequirementDataCreate struct {
	ObjectType string
	Count      uint8
	Progress   uint8
}

func NewKurinNarratorObjectiveRequirementCreate(objectType string, count uint8) *KurinNarratorObjectiveRequirement {
	return &KurinNarratorObjectiveRequirement{
		Type: "create",
		IsDone: func(requirement *KurinNarratorObjectiveRequirement) bool {
			data := requirement.Data.(KurinNarratorObjectiveRequirementDataCreate)
			return data.Progress >= data.Count
		},
		Data: KurinNarratorObjectiveRequirementDataCreate{
			ObjectType: objectType,
			Count:      count,
		},
	}
}

type KurinNarratorObjectiveRequirementDataDestroy struct {
	ObjectType string
	Count      uint8
	Progress   uint8
}

func NewKurinNarratorObjectiveRequirementDestroy(objectType string, count uint8) *KurinNarratorObjectiveRequirement {
	return &KurinNarratorObjectiveRequirement{
		Type: "destroy",
		IsDone: func(requirement *KurinNarratorObjectiveRequirement) bool {
			data := requirement.Data.(KurinNarratorObjectiveRequirementDataDestroy)
			return data.Progress >= data.Count
		},
		Data: KurinNarratorObjectiveRequirementDataDestroy{
			ObjectType: objectType,
			Count:      count,
		},
	}
}
