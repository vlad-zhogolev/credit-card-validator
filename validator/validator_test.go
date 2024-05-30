package validator

import (
	"reflect"
	"testing"
)

type TestValidateData struct {
	digits  []uint8
	isValid bool
}

func TestValidate(t *testing.T) {
	testData := []TestValidateData{
		{[]uint8{1, 7, 8, 9, 3, 7, 2, 9, 9, 7, 4}, true},
		{[]uint8{1, 7, 8, 9, 3, 7, 2, 9, 9, 7, 3}, false},
	}

	for _, data := range testData {
		isValid := Validate(data.digits)
		if isValid != data.isValid {
			t.Fatalf(`Test failed for digits %v, expected validity: %v`, data.digits, data.isValid)
		}
	}
}

func TestNumbersFromString(t *testing.T) {
	testStr := "1234"
	numbers, _ := NumbersFromString(testStr)
	if !reflect.DeepEqual(numbers, []uint8{1, 2, 3, 4}) {
		t.Fatalf(`Test failed for string %v`, testStr)
	}
}
