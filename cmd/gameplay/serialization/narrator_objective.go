package serialization

import "github.com/LamkasDev/kurin/cmd/gameplay"

type NarratorObjectiveData struct {
	Text         string
	Requirements []NarratorObjectiveRequirementData
	Ticks        uint64
}

func EncodeNarratorObjective(objective *gameplay.Objective) NarratorObjectiveData {
	data := NarratorObjectiveData{
		Text:         objective.Text,
		Requirements: []NarratorObjectiveRequirementData{},
		Ticks:        objective.Ticks,
	}
	for _, requirement := range objective.Requirements {
		data.Requirements = append(data.Requirements, EncodeNarratorObjectiveRequirement(requirement))
	}

	return data
}

func DecodeNarratorObjective(data NarratorObjectiveData) *gameplay.Objective {
	objective := &gameplay.Objective{
		Text:         data.Text,
		Requirements: []*gameplay.ObjectiveRequirement{},
		Ticks:        data.Ticks,
	}
	for _, requirementData := range data.Requirements {
		objective.Requirements = append(objective.Requirements, DecodeNarratorObjectiveRequirement(requirementData))
	}

	return objective
}
