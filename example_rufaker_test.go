package rufaker_test

import (
	"fmt"

	"github.com/shakinm/rufaker"
)

func Example_fakeData() {

	// ИНН
	rufaker.PersonInn() // => 500100732259

	// СНИЛС
	snils, _ := rufaker.PersonSnils() // => 82194470004
	fmt.Printf("%s", snils)
	// Result: 82194470004

	formatedSnils := snils.Fromated()
	fmt.Printf("%s", formatedSnils)
	// Result: 821-944-700-04

	// Имя и пол
	rufaker.FirstNameMale()    // => Евгений
	rufaker.FirstNameFemale()  // => Жанна
	rufaker.MiddleNameMale()   // => Фёдорович
	rufaker.MiddleNameFemale() // => Эммануиловна
	rufaker.LastNameMale()     // => Бирюков
	rufaker.LastNameFemale()   // => Одинцова
	fnm := rufaker.FullNameMale()
	fmt.Printf("%+v", fnm)
	// Result:
	/*
		   {
			   firstName: Евгений
			   middleName: Эммануиловна
			   lastName: Одинцова
		   }
	*/
	fnf := rufaker.FullNameFemale()
	fmt.Printf("%+v", fnf)
	// Result:
	/*
		   {
			   firstName: Жанна
			   middleName: Фёдорович
			   lastName: Бирюков
		   }
	*/
	rufaker.FirstName()  // => Снежана
	rufaker.MiddleName() // => Тимофеевич
	rufaker.LastName()   // => Носков
	rfn := rufaker.RandomFullName()
	fmt.Printf("%+v", rfn)
	// Result:
	/*
		   {
			   firstName: Дарья
			   middleName: Алексеевна
			   lastName: Крылова
		   }
	*/
	rufaker.Gender() // => Женский
}
