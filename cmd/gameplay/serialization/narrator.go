package serialization

import "github.com/LamkasDev/kurin/cmd/gameplay"

type KurinNarratorData struct {
	Objectives []KurinNarratorObjectiveData
}

func EncodeKurinNarrator(narrator *gameplay.KurinNarrator) KurinNarratorData {
	data := KurinNarratorData{
		Objectives: []KurinNarratorObjectiveData{},
	}
	for _, objective := range narrator.Objectives {
		data.Objectives = append(data.Objectives, EncodeKurinNarratorObjective(objective))
	}

	return data
}

func DecodeKurinNarrator(data KurinNarratorData) *gameplay.KurinNarrator {
	narrator := gameplay.NewKurinNarrator()
	for _, objectiveData := range data.Objectives {
		gameplay.AddKurinNarratorObjective(narrator, DecodeKurinNarratorObjective(objectiveData))
	}

	return narrator
}
