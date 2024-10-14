package gameplay

func NewMobTemplateCat() *MobTemplate {
	template := NewMobTemplateRaw[interface{}]("cat")
	template.Initialize = func(mob *Mob) {
		mob.Thinktree = NewThinktreeBasic()
	}
	template.Process = ProcessCat

	return template
}

func ProcessCat(mob *Mob) {
	ProcessMob(mob)
	if GameInstance.Ticks%900 == 0 {
		PlaySoundVolume(&GameInstance.SoundController, "meow", 0.1)
	}
}
