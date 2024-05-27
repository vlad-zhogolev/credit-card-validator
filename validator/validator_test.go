package validator

import (
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
		isValid, err := Validate(data.digits)
		if isValid != data.isValid || err != nil {
			t.Fatalf(`Test failed for digits %v, expected validity: %v with error: %v`, data.digits, data.isValid, err)
		}
	}
}
