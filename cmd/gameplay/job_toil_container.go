package gameplay

var JobToilContainer = map[string]*JobToilTemplate{}

func RegisterJobToils() {
	JobToilContainer["build_floor"] = NewJobToilTemplateBuildFloor()
	JobToilContainer["build"] = NewJobToilTemplateBuild()
	JobToilContainer["destroy"] = NewJobToilTemplateDestroy()
	JobToilContainer["destroy_floor"] = NewJobToilTemplateDestroyFloor()
	JobToilContainer["goto"] = NewJobToilTemplateGoto()
	JobToilContainer["manufacture"] = NewJobToilTemplateManufacture()
	JobToilContainer["panic"] = NewJobToilTemplatePanic()
	JobToilContainer["pickup"] = NewJobToilTemplatePickup()
}

func NewJobToil(toilType string, data interface{}) *JobToil {
	toil := &JobToil{
		Type:     toilType,
		Template: JobToilContainer[toilType],
		Data:     data,
	}

	return toil
}
