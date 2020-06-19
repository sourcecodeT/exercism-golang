package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	earthAge := seconds / 31557600

	switch planet {
	case "Mercury":
		return earthAge / 0.2408467
	case "Venus":
		return earthAge / 0.61519726
	case "Mars":
		return earthAge / 1.8808158
	case "Jupiter":
		return earthAge / 11.862615
	case "Saturn":
		return earthAge / 29.447498
	case "Uranus":
		return earthAge / 84.016846
	case "Neptune":
		return earthAge / 164.79132
	default:
		return earthAge
	}
}
