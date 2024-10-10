package gameplay

type ItemCorpseData struct {
	Mob string
}

func NewItemTemplateCorpse() *ItemTemplate {
	template := NewItemTemplate[*ItemWelderData]("corpse", 1, 0)
	template.CanPickup = false
	template.GetDefaultData = func() interface{} {
		return &ItemCorpseData{}
	}

	return template
}
