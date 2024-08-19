package serialization

import "github.com/LamkasDev/kurin/cmd/gameplay"

type NarratorData struct {
	Objectives []NarratorObjectiveData
}

func EncodeNarrator(narrator *gameplay.Narrator) NarratorData {
	data := NarratorData{
		Objectives: []NarratorObjectiveData{},
	}
	for _, objective := range narrator.Objectives {
		data.Objectives = append(data.Objectives, EncodeNarratorObjective(objective))
	}

	return data
}

func DecodeNarrator(data NarratorData) *gameplay.Narrator {
	narrator := gameplay.NewNarrator()
	for _, objectiveData := range data.Objectives {
		gameplay.AddNarratorObjective(narrator, DecodeNarratorObjective(objectiveData))
	}

	return narrator
}
