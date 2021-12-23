package rufaker

import (
	"fmt"
	"math"
	"strconv"
)

// PositiveIntN Генерирует случайное положительное число, параметр n - разрядность числа.
// Разрядность числа не должна превышать разрядность числа math.MaxInt
func PositiveIntN(n int) (result int, err error) {

	if n > len(strconv.Itoa(math.MaxInt)) {
		return result, fmt.Errorf("rufaker/number: argument n MUST be less oe equal then %d", len(strconv.Itoa(math.MaxInt)))
	}

	for i := n; i >= 1; i-- {

		if i == n {
			result = result + (randInt10(1, n+i) * pow10n(i-1))
			continue
		}
		result = result + (randInt10(0, n+i) * pow10n(i-1))
	}

	return result, nil
}
