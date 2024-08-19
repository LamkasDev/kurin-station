package gameplay

type Objective struct {
	Text         string
	Requirements []*ObjectiveRequirement
	Ticks        uint64
}

func NewObjectiveGettingStarted() *Objective {
	return &Objective{
		Text: "Getting Started",
		Requirements: []*ObjectiveRequirement{
			NewObjectiveRequirement("credits", &ObjectiveRequirementDataCredits{Count: 5}),
		},
	}
}

func NewObjectiveCleaningUp() *Objective {
	return &Objective{
		Text: "Cleaning Up",
		Requirements: []*ObjectiveRequirement{
			NewObjectiveRequirement("destroy", &ObjectiveRequirementDataDestroy{ObjectType: "grille", Count: 3}),
		},
	}
}

func NewObjectiveFreshStart() *Objective {
	return &Objective{
		Text: "Fresh Start",
		Requirements: []*ObjectiveRequirement{
			NewObjectiveRequirement("create", &ObjectiveRequirementDataCreate{ObjectType: "displaced", Count: 3}),
		},
	}
}

func NewObjectiveToTheMoon() *Objective {
	return &Objective{
		Text: "To The MOON",
		Requirements: []*ObjectiveRequirement{
			NewObjectiveRequirement("credits", &ObjectiveRequirementDataCredits{Count: 999}),
		},
	}
}
