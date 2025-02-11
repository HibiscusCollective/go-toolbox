package bench

import (
	"testing"

	"github.com/jaswdr/faker/v2"
)

func BenchmarkInvertControl(b *testing.B) {
	fake := faker.New()

	testData := generateTestData(4, 16, generateFakesMap(fake.Int64, fake.Lorem().Word))

	for name, items := range testData {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.RunParallel(func(pb *testing.PB) {
					inverted := make(map[int64]string)
					for key, val := range items {
						inverted[key] = val
					}
				})

				b.ReportAllocs()
			}
		})
	}
}
