package gameplay

import "github.com/LamkasDev/kurin/cmd/common/sdlutils"

func NewThinktreeNodeTemplateAttack() *ThinktreeNodeTemplate {
	template := NewThinktreeNodeTemplate[interface{}]("attack")
	template.Process = func(mob *Mob, node *ThinktreeNode) bool {
		target := FindMob(GameInstance.Map, func(target *Mob) bool {
			return target.Faction != mob.Faction && target.Position.Z == mob.Position.Z && !target.Health.Dead && sdlutils.GetDistance(target.Position.Base, mob.Position.Base) <= 6
		})
		if target != nil {
			job := NewJobDriver("attack", nil)
			job.Data = &JobDriverAttackData{
				Target: target,
			}
			job.Template.Initialize(job)
			AssignTrackerJob(mob.JobTracker, job)

			return true
		}

		return false
	}

	return template
}
