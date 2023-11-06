package hashmap

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"testing"
)

func Test_HasChaining(t *testing.T) {
	m := NewStrWithChaining[int]()
	m.Insert("first", 1)
	m.Insert("second", 2)
	m.Insert("third", 3)
	m.Insert("fourth", 4)
	m.Insert("fifth", 5)
	m.Insert("six", 6)
	m.Insert("seven", 7)
	m.Insert("eight", 8)
	m.Insert("nine", 9)
	m.Insert("ten", 5)

	fmt.Printf("%+v\n", m)

	m.Insert("ten", 10)

	fmt.Printf("%+v\n", m)

	m.Delete("eight")
	m.Delete("six")

	fmt.Printf("%+v\n", m)

	keys := make([]string, 100)
	for i := 0; i < 100; i++ {
		key := uuid.NewString()
		m.Insert(key, rand.Intn(1000000))
		keys = append(keys, key)
	}

	fmt.Printf("%+v\n", m)

	for _, key := range keys {
		m.Delete(key)
	}

	fmt.Printf("%+v\n", m)

	val, ok := m.Find("first")
	fmt.Printf("%+v, %v\n", val, ok)

	val, ok = m.Find("second")
	fmt.Printf("%+v, %v\n", val, ok)
}
