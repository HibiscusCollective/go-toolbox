package bench

import (
	"fmt"

	"github.com/gofrs/uuid/v5"
	"github.com/jaswdr/faker/v2"
)

type character struct {
	Name      string
	Level     int
	Health    float64
	Inventory []string
	Skills    map[string]int
	IsAlive   bool
}

type generate[K comparable, V any] func(num uint) map[K]V
type testData[K comparable, V any] map[string]map[K]V

// generateTestData creates a map of test cases for a given map with an increasing number of elements
// A min and max exponent must be specified. The length of the map will be a power of 2. Ex: 2^4 = 16
func generateTestData[K comparable, V any](minExp, maxExp uint, fn generate[K, V]) testData[K, V] {
	data := make(testData[K, V], maxExp-minExp+1)

	for exp := minExp; exp <= maxExp; exp++ {
		numChars := 1 << exp

		data[fmt.Sprintf("given %d items", numChars)] = fn(uint(numChars))
	}

	return data
}

func generateFakesMap[K comparable, V any](key func() K, val func() V) generate[K, V] {
	return func(num uint) map[K]V {
		fakeMap := make(map[K]V)

		for i := uint(0); i < num; i++ {
			fakeMap[key()] = val()
		}

		return fakeMap
	}
}

func generateCharacters(num uint) map[uuid.UUID]character {
	faker := faker.New()

	chars := make(map[uuid.UUID]character)
	for i := uint(0); i < num; i++ {
		chars[uuid.Must(uuid.NewV4())] = character{
			Name:      faker.Person().Name(),
			Level:     faker.IntBetween(1, 100),
			Health:    faker.Float(2, 0, 100),
			Inventory: []string{"Sword", "Shield", "Potion"},
			Skills: map[string]int{
				"Strength":     faker.IntBetween(1, 20),
				"Intelligence": faker.IntBetween(1, 20),
				"Spirit":       faker.IntBetween(1, 20),
			},
			IsAlive: faker.Bool(),
		}
	}

	return chars
}
