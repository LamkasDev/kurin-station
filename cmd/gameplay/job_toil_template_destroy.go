package gameplay

func NewJobToilTemplateDestroy() *JobToilTemplate {
	template := NewJobToilTemplate[interface{}]("destroy")
	template.Start = ProcessJobToilDestroy
	template.Process = ProcessJobToilDestroy

	return template
}

func ProcessJobToilDestroy(driver *JobDriver, toil *JobToil) JobToilStatus {
	object := GetObjectAtTile(driver.Tile)
	if object == nil {
		return JobToilStatusComplete
	}
	if driver.Character.Fatigue == 0 {
		CharacterHitObject(driver.Character, object)
	}

	return JobToilStatusWorking
}
