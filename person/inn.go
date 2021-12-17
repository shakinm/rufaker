package person

import (
	"math/rand"
	"strconv"
	"time"
)

var k1 = [10]int{7, 2, 4, 10, 3, 5, 9, 4, 6, 8}
var k2 = [11]int{3, 7, 2, 4, 10, 3, 5, 9, 4, 6, 8}

type inn12 string

// Inn Генерирует ИНН физического лица (12 знаков)
func Inn() inn12 {

	rand.Seed(time.Now().UnixNano())

	randomNumber := rand.Perm(10)
	randomNumber = append(randomNumber, calculateCRC1(randomNumber))
	randomNumber = append(randomNumber, calculateCRC2(randomNumber))
	var str string

	for i := 0; i <= 11; i++ {
		str = str + strconv.Itoa(randomNumber[i])
	}
	return inn12(str)
}

func calculateCRC1(num []int) (crc int) {
	for i := 0; i <= 9; i++ {
		crc = crc + num[i]*k1[i]
	}
	return crc % 11
}

func calculateCRC2(num []int) (crc int) {
	for i := 0; i <= 10; i++ {
		crc = crc + num[i]*k2[i]
	}

	crc = crc % 11

	if crc == 10 {
		crc = 0
	}

	return crc
}
