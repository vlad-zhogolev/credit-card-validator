package validator

import (
	"errors"
	"unicode"
)

func NumbersFromString(str string) ([]uint8, error) {
	var res []uint8
	for _, v := range str {
		if v > unicode.MaxASCII || v < '0' || v > '9' {
			return nil, errors.New("string must contain only ascii numbers")
		}
		res = append(res, byte(v)-'0')
	}
	return res, nil
}

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
