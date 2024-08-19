package gameplay

func NewJobToilTemplateManufacture() *JobToilTemplate {
	template := NewJobToilTemplate[interface{}]("manufacture")
	template.Process = ProcessJobToilManufacture

	return template
}

func ProcessJobToilManufacture(driver *JobDriver, toil *JobToil) JobToilStatus {
	lathe := GetObjectAtTile(driver.Tile)
	if lathe == nil {
		return JobToilStatusComplete
	}
	if ProgressOrderAtLathe(lathe) {
		return JobToilStatusComplete
	}

	return JobToilStatusWorking
}
