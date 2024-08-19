package gameplay

type RunechatCharacterData struct {
	Character *Character
}

func NewRunechatCharacter(character *Character, message string) *Runechat {
	runechat := NewRunechat(message)
	runechat.Data = RunechatCharacterData{
		Character: character,
	}

	return runechat
}
