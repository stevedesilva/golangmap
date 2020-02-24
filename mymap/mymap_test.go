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
	assert.Equal(12, productsBought)

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
	assert.Equal("33", bruce2ndNo)

}

func TestMyMap_Student_should_show_message_when_no_args(t *testing.T) {
	m := mymap.New()
	data, err := m.Students(nil)
	if err == nil {
		t.Errorf("Should have received ErrMyMapNoArg: %s",data)
	}
}

func TestMyMap_Student_should_show_missing_message(t *testing.T) {

	m := mymap.New()
	data, err := m.Students([]string{"bobo"})
	if err != nil {
		t.Errorf("Received ErrMyMapNoArg")
	}
	expected := []string(nil)
	assert := assert.New(t)
	assert.Equal(expected, data)
}

func TestMyMap_Student_should_show_data(t *testing.T) {

	m := mymap.New()
	data, err := m.Students([]string{"slytherin"})
	if err != nil {
		t.Errorf("Received ErrMyMapNoArg")
	}
	expected := []string{"higgs", "horace", "nigellus", "scorpius"}
	assert := assert.New(t)
	assert.Equal(expected, data)
}
