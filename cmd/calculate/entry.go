package calculate

import "errors"

// Calculate returns the calculated result requested
func Calculate(a float32, b float32, operation string) (result float32, err error) {

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

func add(a float32, b float32) float32 {
	return a + b
}

func subtract(a float32, b float32) float32 {
	return a - b
}

func multiply(a float32, b float32) float32 {
	return a * b
}

func divide(a float32, b float32) float32 {
	return a / b
}
