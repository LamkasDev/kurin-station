package gameplay

type KurinRunechatCharacterData struct {
	Character *KurinCharacter
}

func NewKurinRunechatCharacter(character *KurinCharacter, message string) *KurinRunechat {
	runechat := NewKurinRunechat(message)
	runechat.Data = KurinRunechatCharacterData{
		Character: character,
	}

	return runechat
}
