package gameplay

type KurinNarratorObjective struct {
	Text         string
	Requirements []*KurinNarratorObjectiveRequirement
	Ticks        uint64
}

func NewKurinNarratorObjectiveGettingStarted() *KurinNarratorObjective {
	return &KurinNarratorObjective{
		Text: "Getting Started",
		Requirements: []*KurinNarratorObjectiveRequirement{
			NewKurinNarratorObjectiveRequirementCredits(5),
		},
	}
}

func NewKurinNarratorObjectiveCleaningUp() *KurinNarratorObjective {
	return &KurinNarratorObjective{
		Text: "Cleaning Up",
		Requirements: []*KurinNarratorObjectiveRequirement{
			NewKurinNarratorObjectiveRequirementDestroy("grille", 3),
		},
	}
}

func NewKurinNarratorObjectiveFreshStart() *KurinNarratorObjective {
	return &KurinNarratorObjective{
		Text: "Fresh Start",
		Requirements: []*KurinNarratorObjectiveRequirement{
			NewKurinNarratorObjectiveRequirementCreate("displaced", 3),
		},
	}
}

func NewKurinNarratorObjectiveToTheMoon() *KurinNarratorObjective {
	return &KurinNarratorObjective{
		Text: "To The MOON",
		Requirements: []*KurinNarratorObjectiveRequirement{
			NewKurinNarratorObjectiveRequirementCredits(999),
		},
	}
}
