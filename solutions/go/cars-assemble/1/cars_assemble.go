package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	successRatePercent := successRate/100
    productionRate_f := float64(productionRate)
    return productionRate_f*successRatePercent
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	carsPerMinute := int(CalculateWorkingCarsPerHour(productionRate, successRate)/60)
    return carsPerMinute
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	singles := carsCount%10
    singlesCost := singles*10000
    batchCars := carsCount-singles
    batchCarsCost := batchCars*9500
    return uint(singlesCost+batchCarsCost)
}
