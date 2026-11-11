package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	// 
	n:=successRate/100
	return float64(productionRate)*n
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	// 
	n:=CalculateWorkingCarsPerHour(productionRate,successRate)
	q:=int(n)
	return q/60

}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	n:=uint(carsCount)
	//21=2*10+1
	//37=3*10+7
	d:=n/10
	r:=n%10
	return d*95000 + r*10000
	
}
