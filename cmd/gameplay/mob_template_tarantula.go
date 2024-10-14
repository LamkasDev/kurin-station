package gameplay

func NewMobTemplateTarantula() *MobTemplate {
	template := NewMobTemplateRaw[interface{}]("tarantula")
	template.Initialize = func(mob *Mob) {
		mob.Thinktree = NewThinktreeBasic()
	}
	template.Process = ProcessTarantula

	return template
}

func ProcessTarantula(mob *Mob) {
	ProcessMob(mob)
}
