package gameplay

var BodypartContainer = map[string]*BodypartTemplate{}

func RegisterBodyparts() {
	BodypartContainer["head"] = NewBodypartTemplate[interface{}]("head", 3)
	BodypartContainer["l_arm"] = NewBodypartTemplate[interface{}]("l_arm", 2)
	BodypartContainer["r_arm"] = NewBodypartTemplate[interface{}]("r_arm", 2)
	BodypartContainer["chest"] = NewBodypartTemplate[interface{}]("chest", 5)
	BodypartContainer["l_leg"] = NewBodypartTemplate[interface{}]("l_leg", 2)
	BodypartContainer["r_leg"] = NewBodypartTemplate[interface{}]("r_leg", 2)
}

func NewBodypart(bodypartType string) *Bodypart {
	bodypart := &Bodypart{
		Type:     bodypartType,
		Template: BodypartContainer[bodypartType],
	}
	bodypart.Points = bodypart.Template.MaxPoints

	return bodypart
}
