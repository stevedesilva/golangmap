package mymap_test

import (
	// "fmt"
	// "reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stevedesilva/golangmap/mymap"
)

func TestMyMap_init_maps(t *testing.T) {
	assert := assert.New(t)

	m := mymap.New()
	assert.NotNil(m)

	// number
	n := m.NumbersMap()
	t.Log(n)

	who, phone := "max", "N/A"
	if v, ok := n[who]; ok {
		phone = v
	}
	assert.Equal(4, len(n))
	assert.Equal("123566777", phone)

	// customer
	c := m.CustomeridMap()
	t.Log(c)

	assert.Equal(2, len(c))
	productsBought := c[2]["a1"]
	// 	2: {"a1": 12},
	assert.Equal(12,productsBought)

	// product
	p := m.ProductMap()
	t.Log(p)

	assert.Equal(2, len(p))
	assert.False(p[1])
	assert.True(p[2])

	// multi
	mm := m.MultinumbersMap()
	t.Log(mm)
	// bruce second number
	bruce2ndNo := mm["bruce"][1]
	assert.Equal(3, len(mm))
	assert.Equal("33",bruce2ndNo)

}

// func TestInsertionSort_should_sort_sequence(t *testing.T) {
// 	type test struct {
// 		input []int
// 		want  []int
// 	}

// 	tests := []test{
// 		{
// 			input: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
// 			want:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
// 		},
// 		{
// 			input: []int{4, 1, 3, 2, 5, 0},
// 			want:  []int{0, 1, 2, 3, 4, 5},
// 		},
// 		{
// 			input: []int{1, 9, 2, 5},
// 			want:  []int{1, 2, 5, 9},
// 		},
// 	}

// 	for _, tc := range tests {
// 		actual := insertion.Sort(tc.input)
// 		if !reflect.DeepEqual(tc.want, actual) {
// 			t.Errorf("expected: %v, got: %v", tc.want, actual)
// 		}
// 	}

// }

// func TestMyMap_init() {
// 	// Hint: Store phone numbers as text
// 	m := mymap.New()

// 	// #1
// 	// Key        : Last name
// 	// Element    : Phone number
// 	numbers := map[string]string{
// 		"max":   "123566777",
// 		"bruce": "123566774",
// 		"saul":  "123566776",
// 		"mike":  "123566737",
// 	}
// 	fmt.Printf("numbers %#v \n", numbers)
// 	fmt.Printf("numbers %v \n", numbers)
// 	// #2
// 	// Key        : Product ID
// 	// Element    : Available / Unavailable
// 	products := map[int]bool{
// 		1: false,
// 		2: true,
// 	}
// 	fmt.Printf("numbers %#v \n", products)

// 	// #3
// 	// Key        : Last name
// 	// Element    : Phone numbers
// 	multinumbers := map[string][]string{
// 		"max":   {"123566777", "988"},
// 		"bruce": {"123566777", "33"},
// 		"saul":  {"123566777", "111"}}
// 	fmt.Printf("numbers %#v \n", multinumbers)

// 	// #4
// 	// Key        : Customer ID
// 	// Element Key:
// 	//   Key: Product ID Element: Quantity
// 	customerid := map[int]map[string]int{
// 		1: {"a2": 1},
// 		2: {"a1": 12},
// 	}
// 	fmt.Printf("numbers %#v \n", customerid)
// }
