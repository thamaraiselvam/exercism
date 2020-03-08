package space

type Planet string

type YearInDays map[Planet]float64
	
type PlanetAge struct {
	seconds float64
	YearInDays
}

func newPlanetAge(seconds float64) PlanetAge {
	return PlanetAge{
		seconds: seconds,
		YearInDays: YearInDays{
			"Earth":   365.25 * 1,
			"Mercury": 365.25 * 0.2408467,
			"Venus":   365.25 * 0.61519726,
			"Mars":    365.25 * 1.8808158,
			"Jupiter": 365.25 * 11.862615,
			"Saturn":  365.25 * 29.447498,
			"Uranus":  365.25 * 84.016846,
			"Neptune": 365.25 * 164.79132,
		},
	}
}

func Age(seconds float64, planet Planet) float64 {
	planetAge := newPlanetAge(seconds)
	return planetAge.convertToYears(planet)
}

func (p PlanetAge) convertToYears(planet Planet) float64 {
	return secondsToDays(p) / p.YearInDays[planet]
}

func secondsToDays(p PlanetAge) float64 {
	minutes := p.seconds / 60
	days := (minutes / 60) / 24
	return days
}