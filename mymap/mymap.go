package mymap

import "fmt"

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
