package mathext_test

import (
	"math/rand"
	"testing"

	"github.com/tanveerprottoy/git-ops/mathext"
)

func TestAdd(t *testing.T) {
	// run parallelly
	t.Parallel()

	tests := []struct {
		name string
		val0 int
		val1 int
		exp  int
	}{
		{"2 + 5", 2, 5, 7},
		{"9 + 5", 9, 5, 14},
		{"27 + 45", 27, 45, 72},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := mathext.Add(tc.val0, tc.val1)
			if actual != tc.exp {
				t.Errorf("Add(%d, %d) = %v; want %v", tc.val0, tc.val1, actual, tc.exp)
			}
		})
	}
}

// bencmarks
// any benchmark should be careful to avoid compiler optimisations eliminating the
// function under test and artificially lowering the run time of the benchmark.
const max = 100

var result int

func BenchmarkAdd(b *testing.B) {
	var r int

	// good practise not to use values from b.N
	for range b.N {
		r = mathext.Add(rand.Intn(max), rand.Intn(max))
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func BenchmarkAll(b *testing.B) {
	b.Run("Add", BenchmarkAdd)
}
