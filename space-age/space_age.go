package space

type Planet string

const (
	Mercury Planet = "Mercury"
	Venus   Planet = "Venus"
	Earth   Planet = "Earth"
	Mars    Planet = "Mars"
	Jupiter Planet = "Jupiter"
	Saturn  Planet = "Saturn"
	Uranus  Planet = "Uranus"
	Neptune Planet = "Neptune"
)

var planetSeconds = map[Planet]float64{
	Mercury: 0.2408467,
	Venus:   0.61519726,
	Earth:   31557600,
	Mars:    1.8808158,
	Jupiter: 11.862615,
	Saturn:  29.447498,
	Uranus:  84.016846,
	Neptune: 164.79132,
}

func Age(seconds float64, planet Planet) float64 {
	orbitalPeriod, valid := planetSeconds[planet]
	if !valid {
		return -1.00
	}
	if planet == Earth {
		return seconds / planetSeconds[Earth]
	}
	inSecondes := planetSeconds[Earth] * orbitalPeriod
	return seconds / inSecondes
}
