package annalyn

func CanFastAttack(knightIsAwake bool) bool {
	return !knightIsAwake
}

func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	return knightIsAwake || archerIsAwake || prisonerIsAwake
}

func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	return prisonerIsAwake && !archerIsAwake
}

func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	switch {
	case petDogIsPresent && !archerIsAwake:
		return true
	case !petDogIsPresent && !archerIsAwake && !knightIsAwake && prisonerIsAwake:
		return true
	}

	return false
}
