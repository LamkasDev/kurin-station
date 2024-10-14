package gameplay

type JobToilAttackData struct {
	Target *Mob
}

func NewJobToilTemplateAttack() *JobToilTemplate {
	template := NewJobToilTemplate[*JobToilAttackData]("attack")
	template.Process = ProcessJobToilAttack

	return template
}

func ProcessJobToilAttack(driver *JobDriver, toil *JobToil) JobToilStatus {
	data := toil.Data.(*JobToilAttackData)
	if data.Target.Health.Dead {
		return JobToilStatusComplete
	}
	if !CanMobInteractWithMob(driver.Mob, data.Target) {
		return JobToilStatusComplete
	}
	if driver.Mob.Fatigue == 0 {
		MobHitMob(driver.Mob, data.Target)
		return JobToilStatusComplete
	}

	return JobToilStatusWorking
}
