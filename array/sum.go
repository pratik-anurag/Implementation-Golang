package array

import (
	"errors"
	"reflect"
)

func sum(numericArray []interface{}) (float64, error) {

	// if size of slice is empty then return 0, with an error
	if len(numericArray) == 0 {
		return 0, errors.New("slice is empty")
	}

	// if slice's first element is string return 0, with an error
	if reflect.TypeOf(numericArray[0]).String() == "string" {
		return 0, errors.New("slice cannot contain non numeric data types")
	}

	floatSum := 0.0
	intSum := 0
	for _, num := range numericArray {
		// if a type other than int or float appears in the slice, then return 0 with an error
		dataType := reflect.TypeOf(num).String()
		if dataType != "int" && dataType != "float64" {
			return 0, errors.New("slice cannot contain non numeric data types")
		}
		if dataType == "int" {
			intSum += num.(int)
		} else {
			floatSum += num.(float64)
		}
	}

	return floatSum + float64(intSum), nil
}
