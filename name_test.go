package rufaker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	t.Run("Генерация мужского имени", func(t *testing.T) {
		assert.True(t, hasElementSliceString(firstNamesMale, FirstNameMale()))
	})
	t.Run("Генерация женского имени", func(t *testing.T) {
		assert.True(t, hasElementSliceString(firstNamesFemale, FirstNameFemale()))
	})
	t.Run("Генерация мужского отчества", func(t *testing.T) {
		assert.True(t, hasElementSliceString(middleNamesMale, MiddleNameMale()))
	})
	t.Run("Генерация женского отчества", func(t *testing.T) {
		assert.True(t, hasElementSliceString(middleNamesFemale, MiddleNameFemale()))
	})
	t.Run("Генерация мужской фамилии", func(t *testing.T) {
		assert.True(t, hasElementSliceString(lastName, LastNameMale()))
	})
	t.Run("Генерация женской фамилии", func(t *testing.T) {
		ln := LastNameFemale()
		assert.True(t, hasElementSliceString(lastName, string(ln[:len(ln)-2])))
	})
	t.Run("Генерация мужских ФИО", func(t *testing.T) {
		fnm := FullNameMale()
		assert.True(t, hasElementSliceString(firstNamesMale, fnm.firstName))
		assert.True(t, hasElementSliceString(middleNamesMale, fnm.middleName))
		assert.True(t, hasElementSliceString(lastName, fnm.lastName))
	})
	t.Run("Генерация женских ФИО", func(t *testing.T) {
		fnf := FullNameFemale()
		assert.True(t, hasElementSliceString(firstNamesFemale, fnf.firstName))
		assert.True(t, hasElementSliceString(middleNamesFemale, fnf.middleName))
		assert.True(t, hasElementSliceString(lastName, string(fnf.lastName[:len(fnf.lastName)-2])))
	})
	t.Run("Генерация случайного имени", func(t *testing.T) {
		assert.True(t, hasElementSliceString(firstNames, FirstName()))
	})
	t.Run("Генерация случайного отчества", func(t *testing.T) {
		assert.True(t, hasElementSliceString(middleNames, MiddleName()))
	})
	t.Run("Генерация случайных ФИО", func(t *testing.T) {
		fn := RandomFullName()
		assert.True(t, hasElementSliceString(firstNames, fn.firstName))
		assert.True(t, hasElementSliceString(middleNames, fn.middleName))
	})
}
func TestGender(t *testing.T) {
	t.Run("Генерация пола", func(t *testing.T) {
		assert.True(t, hasElementSliceString(genders, Gender()))
	})
}
