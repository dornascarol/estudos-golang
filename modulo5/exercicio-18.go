// Weather Forecast

// Package weather é uma biblioteca para retornar a previsão do tempo.
package weather

// CurrentCondition descreve a condição climática atual.
var CurrentCondition string

// CurrentLocation indica o local para qual a previsão é desejada.
var CurrentLocation string

// Forecast retorna a previsão do tempo com base na cidade e nas condições.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
