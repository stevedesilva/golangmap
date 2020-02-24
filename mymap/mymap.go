package mymap

import (
	"errors"
	"fmt"
	"sort"
)

// // ---------------------------------------------------------
// // EXERCISE: Warm-up
// //
// //  Create and print the following maps.
// //
// //  1. Phone numbers by last name
// //  2. Product availability by Product ID
// //  3. Multiple phone numbers by last name
// //  4. Shopping basket by Customer ID
// //
// //     Each item in the shopping basket has a Product ID and
// //     quantity. Through the map, you can tell:
// //     "Mr. X has bought Y bananas"
// //
// // ---------------------------------------------------------

// MyMaps base
type MyMaps struct{}

// New constructor
func New() *MyMaps {
	return &MyMaps{}
}

// NumbersMap exercise
func (m *MyMaps) NumbersMap() map[string]string {

	// Key        : Last name
	// Element    : Phone number
	numbers := map[string]string{
		"max":   "123566777",
		"bruce": "123566774",
		"saul":  "123566776",
		"mike":  "123566737",
	}
	fmt.Printf("numbers %#v \n", numbers)
	fmt.Printf("numbers %v \n", numbers)
	return numbers
}

// ProductMap exercise
func (m *MyMaps) ProductMap() map[int]bool {

	// Key        : Product ID
	// Element    : Available / Unavailable
	products := map[int]bool{
		1: false,
		2: true,
	}
	fmt.Printf("numbers %#v \n", products)
	return products
}

// MultinumbersMap exercise
func (m *MyMaps) MultinumbersMap() map[string][]string {

	// #3
	// Key        : Last name
	// Element    : Phone numbers
	multinumbers := map[string][]string{
		"max":   {"123566777", "988"},
		"bruce": {"123566777", "33"},
		"saul":  {"123566777", "111"}}
	fmt.Printf("numbers %#v \n", multinumbers)

	return multinumbers
}

// CustomeridMap exercise
func (m *MyMaps) CustomeridMap() map[int]map[string]int {
	// #4
	// Key        : Customer ID
	// Element Key:
	//   Key: Product ID Element: Quantity
	customerid := map[int]map[string]int{
		1: {"a2": 1},
		2: {"a1": 12},
	}
	fmt.Printf("numbers %#v \n", customerid)
	return customerid

}

// Students example
//
//  Create a program that returns the students by the given
//  Hogwarts house name (see the data below).
//
//  Print the students sorted by name.
//
//  "bobo" doesn't belong to Hogwarts, remove it by using
//  the delete function.
//
//
// RESTRICTIONS
//
//  + Add the following data to your map as is.
//    Do not sort it manually and do not modify it.
//
//  + Slices in the map shouldn't be sorted (changed).
//    HINT: Copy them.
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//
//  Please type a Hogwarts house name.
//
//
//  go run main.go bobo
//
//  Sorry. I don't know anything about "bobo".
//
//
//  go run main.go hufflepuf
//
//  ~~~ hufflepuf students ~~~
//
//        + diggory
//        + helga
//        + scamander
//        + wenlock
//
// ---------------------------------------------------------
func (m *MyMaps) Students(args []string) ([]string, error) {

	// House        Student Name
	// ---------------------------
	// gryffindor   weasley
	// gryffindor   hagrid
	// gryffindor   dumbledore
	// gryffindor   lupin
	// hufflepuf    wenlock
	// hufflepuf    scamander
	// hufflepuf    helga
	// hufflepuf    diggory
	// ravenclaw    flitwick
	// ravenclaw    bagnold
	// ravenclaw    wildsmith
	// ravenclaw    montmorency
	// slytherin    horace
	// slytherin    nigellus
	// slytherin    higgs
	// slytherin    scorpius
	// bobo         wizardry
	// bobo         unwanted

	house := make(map[string][]string, len(args))

	house["gryffindor"] = []string{"weasley", "hagrid", "dumbledore", "lupin"}
	house["hufflepuf"] = []string{"wenlock", "scamander", "helga", "diggory"}
	house["ravenclaw"] = []string{"flitwick", "bagnold", "wildsmith", "montmorency"}
	house["slytherin"] = []string{"horace", "nigellus", "higgs", "scorpius"}
	house["bobo"] = []string{"wizardry", "unwanted"}

	delete(house, "bobo")

	if len(args) < 1 {
		return nil, ErrMyMapNoArg
	}

	in := args[0]

	clone := append([]string(nil), house[in]...)

	sort.Strings(clone)

	return clone, nil
}

// ErrMyMapNoArg error
var ErrMyMapNoArg = errors.New("Please type a Hogwarts house name")
