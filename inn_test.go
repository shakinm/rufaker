package rufaker

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInn(t *testing.T) {
	t.Run("Самопроверка валидатора на реальных ИНН", func(t *testing.T) {
		testInn := []string{
			"500100732259",
		}

		for _, v := range testInn {
			assert.True(t, validateInn(Inn(v)))
		}
	})
	t.Run("Проверка генерации 12 значного ИНН", func(t *testing.T) {
		assert.True(t, validateInn(PersonInn()))
	})
}

func validateInn(v Inn) bool {
	var k1 = [10]int{7, 2, 4, 10, 3, 5, 9, 4, 6, 8}
	var k2 = [11]int{3, 7, 2, 4, 10, 3, 5, 9, 4, 6, 8}

	var numsInn [12]int

	for i := 0; i <= 11; i++ {
		n, err := strconv.Atoi(string(v[i]))
		if err != nil {
			return false
		}
		numsInn[i] = n
	}

	var crc1, crc2 int

	for i := 0; i <= 9; i++ {
		crc1 = crc1 + numsInn[i]*k1[i]
	}

	if crc1%11 != numsInn[10] {
		return false
	}

	for i := 0; i <= 10; i++ {
		crc2 = crc2 + numsInn[i]*k2[i]
	}
	crc2 = crc2 % 11

	if crc2 == 10 {
		crc2 = 0
	}

	if crc2 != numsInn[11] {
		return false
	}
	return true
}
