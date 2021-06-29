package vehicle

type Vehicle interface {
	Structure() []string // Common Method
	Speed() string
}

type Car string

func (c Car) Structure() []string {
	var parts = []string{"ECU", "Engine", "Air filter", "Wipers", "Gas Task"}
	return parts
}

func (c Car) Speed() string {
	return "200 Km/hrs"
}
