//Package rufaker предоставляет инструменты для генерации данных случайных данных
package rufaker

import (
	"math/rand"
	"time"
)

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

func randomElementFromSliceString(v []string) string {
	return v[rand.Int()%len(v)]
}

func randomBoolean() bool {
	return time.Now().Nanosecond()%2 == 0
}

func hasElementSliceString(slice []string, search string) bool {
	for _, s := range slice {
		if s == search {
			return true
		}
	}
	return false
}
