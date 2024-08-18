package gameplay

type KurinNarratorObjectiveRequirementDataDestroy struct {
	ObjectType string
	Count      uint8
	Progress   uint8
}

func NewKurinNarratorObjectiveRequirementDestroy(objectType string, count uint8) *KurinNarratorObjectiveRequirement {
	requirement := NewKurinNarratorObjectiveRequirementRaw[KurinNarratorObjectiveRequirementDataDestroy]("destroy")
	requirement.IsDone = func(requirement *KurinNarratorObjectiveRequirement) bool {
		data := requirement.Data.(KurinNarratorObjectiveRequirementDataDestroy)
		return data.Progress >= data.Count
	}
	requirement.Data = KurinNarratorObjectiveRequirementDataDestroy{
		ObjectType: objectType,
		Count:      count,
	}

	return requirement
}
