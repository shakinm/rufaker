// Пакет предоставляет инструменты для генерации данных связанных
// с персональной информацией о человеке
package person

import (
	"fmt"
	"math/rand"
	number "rufaker/common"
	"strconv"
)

type snils string

// fromated вывод текущего значения в формате ХХХ-ХХХ-ХХХ-YY
func (t snils) fromated() string {
	s := string(t)
	return fmt.Sprintf("%s-%s-%s-%s", s[:3], s[3:6], s[6:9], s[9:])
}

// Snils Генерирует случайный СНИЛС
func Snils() (snils, error) {

	var snilsNum string
	var numsSum int
	randomString := shuffleChar(strconv.Itoa(int(number.PositeiveIntN(9))))
	for k := 9; k >= 1; k-- {

		n, err := strconv.Atoi(string(randomString[9-k]))

		if err != nil {
			return snils(""), fmt.Errorf("rufaker/person Snils: error - %s", err.Error())
		}

		numsSum += k * n
		snilsNum = snilsNum + (strconv.Itoa(n))
	}
	if numsSum < 100 {
		snilsNum = snilsNum + fmt.Sprintf("%02d", numsSum)
	} else if numsSum <= 101 {
		snilsNum = snilsNum + "00"
	} else {
		snilsNum = snilsNum + fmt.Sprintf("%02d", numsSum%101)
	}
	return snils(snilsNum), nil
}

func shuffleChar(s string) string {
	output := make([]byte, len(s))
	for i := len(s) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		output[i], output[j] = s[j], s[i]
	}

	return string(output)
}
