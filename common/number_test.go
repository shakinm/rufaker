package number

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPositeiveIntN(t *testing.T) {
	numbers := len(strconv.Itoa(math.MaxInt))
	for i := 1; i <= numbers; i++ {
		assert.EqualValues(t, i, len(strconv.Itoa(PositeiveIntN(i))))
	}
}
