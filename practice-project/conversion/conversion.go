package conversion

import (
	"errors"
	"strconv"
)

func StringToFloat(strings []string) ([]float64, error) {
	var floats []float64

	for _, stringVal := range strings {
		floatVal, err := strconv.ParseFloat(stringVal, 64)

		// Check if there's an error converting string to float64
		if err != nil {
			return nil, errors.New("failed to convert string to float64")
		}

		floats = append(floats, floatVal)
	}

	return floats, nil
}
