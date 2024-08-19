package gameplay

var JobDriverContainer = map[string]*JobDriverTemplate{}

func RegisterJobDrivers() {
	JobDriverContainer["build_floor"] = NewJobDriverTemplateBuildFloor()
	JobDriverContainer["build"] = NewJobDriverTemplateBuild()
	JobDriverContainer["destroy_floor"] = NewJobDriverTemplateDestroyFloor()
	JobDriverContainer["destroy"] = NewJobDriverTemplateDestroy()
	JobDriverContainer["manufacture"] = NewJobDriverTemplateManufacture()
	JobDriverContainer["panic"] = NewJobDriverTemplatePanic()
}

func NewJobDriver(jobType string, tile *Tile) *JobDriver {
	driver := &JobDriver{
		Type:     jobType,
		Tile:     tile,
		Template: JobDriverContainer[jobType],
	}

	return driver
}
