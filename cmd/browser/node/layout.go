package node

type KitsuneLayoutDirection bool

var KitsuneLayoutDirectionRow = KitsuneLayoutDirection(false)
var KitsuneLayoutDirectionColumn = KitsuneLayoutDirection(true)

func GetKitsuneElementLayoutDirection(element *KitsuneElement) KitsuneLayoutDirection {
	switch element.Data.(type) {
	case *KitsuneElementTextData:
		return KitsuneLayoutDirectionColumn
	}

	return KitsuneLayoutDirectionRow
}
