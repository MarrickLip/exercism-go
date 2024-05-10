package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

var BATCH_SIZE = 10
var BATCH_COST = 95000
var INDIVIDUAL_COST = 10000

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	nBatches := carsCount / BATCH_SIZE
	nIndividuals := carsCount - (nBatches * BATCH_SIZE)

	return uint(nBatches*BATCH_COST + nIndividuals*INDIVIDUAL_COST)
}
