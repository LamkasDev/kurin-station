package gameplay

type MobCharacterData struct {
	Inventory *Inventory
}

func NewMobTemplateCharacter() *MobTemplate {
	template := NewMobTemplateRaw[interface{}]("character")
	template.Initialize = func(mob *Mob) {
		mob.Data = &MobCharacterData{
			Inventory: NewInventory(),
		}
	}
	template.Process = ProcessCharacter

	return template
}

func ProcessCharacter(character *Mob) {
	if character.Fatigue > 0 {
		character.Fatigue--
	}
	if GameInstance.SelectedCharacter != character {
		if !ProcessJobTracker(character.JobTracker) {
			ProcessThinktree(character)
		}
	}
}
