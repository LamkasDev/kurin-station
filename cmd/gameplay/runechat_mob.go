package gameplay

type RunechatMobData struct {
	Mob *Mob
}

func NewRunechatMob(mob *Mob, message string) *Runechat {
	runechat := NewRunechat(message)
	runechat.Data = RunechatMobData{
		Mob: mob,
	}

	return runechat
}
