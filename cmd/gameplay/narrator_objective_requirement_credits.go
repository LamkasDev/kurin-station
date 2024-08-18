package gameplay

type KurinNarratorObjectiveRequirementDataCredits struct {
	Count uint32
}

func NewKurinNarratorObjectiveRequirementCredits(count uint32) *KurinNarratorObjectiveRequirement {
	requirement := NewKurinNarratorObjectiveRequirementRaw[KurinNarratorObjectiveRequirementDataCredits]("credits")
	requirement.IsDone = func(requirement *KurinNarratorObjectiveRequirement) bool {
		return GameInstance.Credits >= requirement.Data.(KurinNarratorObjectiveRequirementDataCredits).Count
	}
	requirement.Data = KurinNarratorObjectiveRequirementDataCredits{
		Count: count,
	}

	return requirement
}
