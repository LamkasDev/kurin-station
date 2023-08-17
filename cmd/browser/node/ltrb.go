package node

type KitsuneElementRectLTRB struct {
	Left   KitsuneElementValueVariable
	Top    KitsuneElementValueVariable
	Right  KitsuneElementValueVariable
	Bottom KitsuneElementValueVariable
}

func NewKitsuneElementRect(margin int32) KitsuneElementRectLTRB {
	return NewKitsuneElementRectLTRB(margin, margin, margin, margin)
}

func NewKitsuneElementRectHV(horizonal int32, vertical int32) KitsuneElementRectLTRB {
	return NewKitsuneElementRectLTRB(horizonal, vertical, horizonal, vertical)
}

func NewKitsuneElementRectLTRB(left int32, top int32, right int32, bottom int32) KitsuneElementRectLTRB {
	return KitsuneElementRectLTRB{
		Left: KitsuneElementValueVariable{
			Type:  KitsuneElementValueVariableFixed,
			Value: left,
		},
		Right: KitsuneElementValueVariable{
			Type:  KitsuneElementValueVariableFixed,
			Value: right,
		},
		Top: KitsuneElementValueVariable{
			Type:  KitsuneElementValueVariableFixed,
			Value: top,
		},
		Bottom: KitsuneElementValueVariable{
			Type:  KitsuneElementValueVariableFixed,
			Value: bottom,
		},
	}
}
