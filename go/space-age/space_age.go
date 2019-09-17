// The Age function is given seconds on a specific planet  and returns how old someone would be in Earth years.
package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	var earthDays = map[Planet]float64{
		"Earth" : 1.0,
		"Mercury" : 0.2408467,
		"Venus" : 0.61519726,
		"Mars" : 1.8808158,
		"Jupiter" : 11.862615,
		"Saturn" : 29.447498,
		"Uranus" : 84.016846,
		"Neptune" : 164.79132,
	}
	var earthSeconds = 31557600.0
	var age = (seconds / earthSeconds) / earthDays[planet]	
	return age
}