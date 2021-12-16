// Пакет предоставляет функции генерации случайных значений для общих типов
package number

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
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

func randInt10(min int, seed int) int {
	rand.Seed(time.Now().UnixNano() + int64(seed))
	return min + rand.Intn(9-min)
}

func pow10n(n int) int {
	b := 10

	if n == 0 {
		return 1
	}

	if n == 1 {
		return b
	}

	for i := 0; i <= n-2; i++ {
		b = b * 10
	}

	return b
}
