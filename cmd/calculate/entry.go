package calculate

import "errors"

// Calculate returns the calculated result requested
func Calculate(a int, b int, operation string) (result float64, err error) {

	switch operation {
	case "ADD":
		result = add(a, b)
	case "SUBTRACT":
		result = subtract(a, b)
	case "MULTIPLY":
		result = multiply(a, b)
	case "DIVIDE":
		result = divide(a, b)
	default:
		return 0, errors.New(`Math operation is not recognized`)
	}
	return result, nil
}

func add(a int, b int) float64 {
	return float64(a) + float64(b)
}

func subtract(a int, b int) float64 {
	return float64(a) - float64(b)
}

func multiply(a int, b int) float64 {
	return float64(a) * float64(b)
}

func divide(a int, b int) float64 {
	return float64(a) / float64(b)
}
