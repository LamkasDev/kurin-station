package serialization

import "github.com/LamkasDev/kurin/cmd/gameplay"

type KurinNarratorObjectiveData struct {
	Text         string
	Requirements []KurinNarratorObjectiveRequirementData
	Ticks        uint64
}

func EncodeKurinNarratorObjective(objective *gameplay.KurinNarratorObjective) KurinNarratorObjectiveData {
	data := KurinNarratorObjectiveData{
		Text:         objective.Text,
		Requirements: []KurinNarratorObjectiveRequirementData{},
		Ticks:        objective.Ticks,
	}
	for _, requirement := range objective.Requirements {
		data.Requirements = append(data.Requirements, EncodeKurinNarratorObjectiveRequirement(requirement))
	}

	return data
}

func DecodeKurinNarratorObjective(data KurinNarratorObjectiveData) *gameplay.KurinNarratorObjective {
	objective := &gameplay.KurinNarratorObjective{
		Text:         data.Text,
		Requirements: []*gameplay.KurinNarratorObjectiveRequirement{},
		Ticks:        data.Ticks,
	}
	for _, requirementData := range data.Requirements {
		objective.Requirements = append(objective.Requirements, DecodeKurinNarratorObjectiveRequirement(requirementData))
	}

	return objective
}
