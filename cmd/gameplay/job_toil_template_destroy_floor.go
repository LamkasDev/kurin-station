package gameplay

func NewJobToilTemplateDestroyFloor() *JobToilTemplate {
	template := NewJobToilTemplate[interface{}]("destroy_floor")
	template.Process = ProcessJobToilDestroyFloor

	return template
}

func ProcessJobToilDestroyFloor(driver *JobDriver, toil *JobToil) JobToilStatus {
	if toil.Ticks >= 180 {
		PlaySoundVolume(&GameInstance.SoundController, "grillehit", 0.5)
		DestroyTile(driver.Tile)
		return JobToilStatusComplete
	}

	return JobToilStatusWorking
}
