package bench

import (
	"testing"

	"github.com/HibiscusCollective/go-toolbox/fxmap"
	"github.com/jaswdr/faker/v2"
)

func BenchmarkInvert(b *testing.B) {
	fake := faker.New()

	testData := generateTestData(4, 16, generateFakesMap(fake.Int64, fake.Lorem().Word))

	for name, td := range testData {
		b.Run(name+": control", RunInvertControlBenchmark(td))
		b.Run(name+": experiment", RunInvertExperimentBenchmark(td))
	}

}

func RunInvertControlBenchmark(items map[int64]string) func(b *testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StartTimer()
			inverted := make(map[string]int64)
			for key, val := range items {
				inverted[val] = key
			}
			b.StopTimer()
		}

		b.ReportAllocs()
	}
}

func RunInvertExperimentBenchmark(items map[int64]string) func(b *testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StartTimer()
			fxmap.Invert(items)
			b.StopTimer()
		}

		b.ReportAllocs()
	}
}
