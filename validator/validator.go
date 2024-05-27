package validator

import (
	"errors"
)

func Validate(digits []uint8) (bool, error) {
	if len(digits) < 2 {
		return false, errors.New("digits must have at least 2 elements")
	}

	sum := 0
	shouldTakeMod := false
	for i := len(digits) - 1; i >= 0; i-- {
		if shouldTakeMod {
			digit := digits[i] * 2
			if digit > 9 {
				digit -= 9
			}
			sum += int(digit)
		} else {
			sum += int(digits[i])
		}
		shouldTakeMod = !shouldTakeMod
	}

	return sum%10 == 0, nil
}
