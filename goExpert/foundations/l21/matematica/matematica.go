package matematica

type Number interface {
	~int | ~float64
}

var PI = 3.14

// Only functions that start with first letter Upper Case are exported
func Sum[T Number](number ...T) T {
	var total T

	for _, value := range number {
		total += value
	}

	return total
}
