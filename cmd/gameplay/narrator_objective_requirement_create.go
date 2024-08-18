package gameplay

type KurinNarratorObjectiveRequirementDataCreate struct {
	ObjectType string
	Count      uint8
	Progress   uint8
}

func NewKurinNarratorObjectiveRequirementCreate(objectType string, count uint8) *KurinNarratorObjectiveRequirement {
	requirement := NewKurinNarratorObjectiveRequirementRaw[KurinNarratorObjectiveRequirementDataCreate]("create")
	requirement.IsDone = func(requirement *KurinNarratorObjectiveRequirement) bool {
		data := requirement.Data.(KurinNarratorObjectiveRequirementDataCreate)
		return data.Progress >= data.Count
	}
	requirement.Data = KurinNarratorObjectiveRequirementDataCreate{
		ObjectType: objectType,
		Count:      count,
	}

	return requirement
}
