// Пакет предоставляет функции генерации случаных значений для общих типов
package number

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

// PositeiveIntN Генерирует случайное полодительное число, параметр n - разрядность числа
func PositeiveIntN(n int) int {
	rand.Seed(time.Now().UnixNano())
	var interval int
	if n == len(strconv.Itoa(math.MaxInt)) {
		interval = math.MaxInt
	} else {
		interval = int(math.Pow(10, float64(n)) - 1)
	}
	return rand.Intn(interval)
}
