//Package weather package is responsible for forescasting the weather.
package weather

var (
	//CurrentCondition required to forecast the weather.
	CurrentCondition string
	//CurrentLocation are required to forecast the weather.
	CurrentLocation  string
)
//Forecast  forecasts the weather It takes city and condition which is the currentlocation and currentcondition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
