package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	if planet == "Earth" {
		return seconds / 31557600
	}
	if planet == "Mercury" {
		return (seconds / 31557600) / 0.2408467
	}
	if planet == "Venus" {
		return seconds / 31557600 / 0.61519726
	}
	if planet == "Mars" {
		return seconds / 31557600 / 1.8808158
	}
	if planet == "Jupiter" {
		return (seconds / 31557600) / 11.862615
	}
	if planet == "Saturn" {
		return (seconds / 31557600) / 29.447498
	}
	if planet == "Uranus" {
		return seconds / 31557600 / 84.016846
	}
	if planet == "Neptune" {
		return (seconds / 31557600) / 164.79132
	}
	return 0
}
