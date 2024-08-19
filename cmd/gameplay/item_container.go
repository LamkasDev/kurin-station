package gameplay

var ItemContainer = map[string]*ItemTemplate{}

func RegisterItems() {
	ItemContainer["rod"] = NewItemTemplate[interface{}]("rod", 3)
	ItemContainer["credit"] = NewItemTemplate[interface{}]("credit", 1)
	ItemContainer["survivalknife"] = NewItemTemplate[interface{}]("survivalknife", 1)
	ItemContainer["welder"] = NewItemTemplateWelder()
}

func NewItem(itemType string, count uint16) *Item {
	item := &Item{
		Id:       GetNextId(),
		Type:     itemType,
		Count:    count,
		Template: ItemContainer[itemType],
	}
	item.Data = item.Template.GetDefaultData()

	return item
}
