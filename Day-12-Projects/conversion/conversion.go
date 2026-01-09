package conversion

import (
	"errors"
	"strconv"
)

func StringToFloat(strings []string) ([]float64, error) {
	var floats []float64
	for _, str := range strings {

		floatPrice, err := strconv.ParseFloat(str, 64)
		if err != nil {
			
			return nil,errors.New("Failed to convert string to float")
		}
		floats = append(floats, floatPrice)
	}
	return  floats,nil
	
}