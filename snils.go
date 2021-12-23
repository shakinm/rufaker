package rufaker

import (
	"fmt"
	"strconv"
)

// Snils Страховой номер индивидуального лицевого счёта (СНИЛС)
type Snils string

// Fromated вывод текущего значения в формате ХХХ-ХХХ-ХХХ-YY
func (t Snils) Fromated() string {
	s := string(t)
	return fmt.Sprintf("%s-%s-%s-%s", s[:3], s[3:6], s[6:9], s[9:])
}

// PersonSnils Генерирует случайный СНИЛС
func PersonSnils() (Snils, error) {

	var snilsNum string
	var crc int

	randomNumber, err := PositiveIntN(9)

	if err != nil {
		return Snils(""), fmt.Errorf("rufaker/person Snils: %s", err.Error())
	}

	randomString := strconv.Itoa(randomNumber)
	for k := 9; k >= 1; k-- {

		n, err := strconv.Atoi(string(randomString[9-k]))

		if err != nil {
			return Snils(""), fmt.Errorf("rufaker/person Snils: error - %s", err.Error())
		}

		crc += k * n
		snilsNum = snilsNum + (strconv.Itoa(n))
	}
	if crc < 100 {
		snilsNum = snilsNum + fmt.Sprintf("%02d", crc)
	} else if crc <= 101 {
		snilsNum = snilsNum + "00"
	} else {
		snilsNum = snilsNum + fmt.Sprintf("%02d", crc%101)
	}
	return Snils(snilsNum), nil
}
