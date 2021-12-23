package rufaker

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPositiveIntN(t *testing.T) {
	numbers := len(strconv.Itoa(math.MaxInt))
	for i := 1; i <= numbers; i++ {
		randomInt, err := PositiveIntN(i)
		assert.NoError(t, err)
		assert.EqualValues(t, i, len(strconv.Itoa(randomInt)))
	}
}
