package gameplay

func NewThinktreeNodeTemplateRevenge() *ThinktreeNodeTemplate {
	template := NewThinktreeNodeTemplate[interface{}]("revenge")
	template.Process = func(mob *Mob, node *ThinktreeNode) bool {
		if IsMobAggro(mob) {
			switch attacker := mob.Health.LastDamageSource.(type) {
			case *Mob:
				if mob.JobTracker.Job != nil && mob.JobTracker.Job.Type != "attack" {
					UnassignTrackerJob(mob.JobTracker)
				}
				if mob.JobTracker.Job == nil && GameInstance.Ticks >= mob.JobTracker.LastTimeoutTicks {
					job := NewJobDriver("attack", nil)
					job.Data = &JobDriverAttackData{
						Target: attacker,
					}
					job.Template.Initialize(job)
					AssignTrackerJob(mob.JobTracker, job)
				}

				return true
			}
		}

		return false
	}

	return template
}
