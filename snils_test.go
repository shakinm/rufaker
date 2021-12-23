package rufaker

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPersonSnils(t *testing.T) {
	t.Run("Самопроверка валидатора на реальных СНИЛС", func(t *testing.T) {
		testSnils := []string{
			"03170525210",
			"81361336680",
			"17268189000",
			"82194470004",
		}

		for _, v := range testSnils {
			assert.True(t, validateSnils(Snils(v)))
		}
	})
	t.Run("Генерация 11 значного числа", func(t *testing.T) {
		fakeSnils, err := PersonSnils()
		assert.NoError(t, err)
		assert.True(t, validateSnils(fakeSnils))
	})
	t.Run("Генерация в формате ХХХ-ХХХ-ХХХ-YY", func(t *testing.T) {
		fakeSnils, err := PersonSnils()
		assert.NoError(t, err)
		assert.NotEmpty(t, fakeSnils.Fromated())
	})
}

func validateSnils(v Snils) bool {

	crc, err := strconv.Atoi(string(v[9:]))

	if err != nil {
		return false
	}

	nums := (string(v[:9]))
	var numsSum int
	for k := 9; k >= 1; k-- {
		n, err := strconv.Atoi(string(nums[9-k]))

		if err != nil {
			return false
		}

		numsSum += k * n
	}

	return crc == getCrc(numsSum)
}

func getCrc(number int) int {
	if number == 100 || number == 101 {
		number = 0
	} else if number > 101 {
		number = number % 101
	}
	if number >= 100 {
		number = getCrc(number)
	}
	return number
}
