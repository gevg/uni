package gamma

import (
	"fmt"
	"testing"
)

type testpair struct {
	in  []uint32
	out []uint64
}

var tests = testpair{
	[]uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
	[]uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
}

func TestEncode1(t *testing.T) {
	out := Encode(tests.in)
	fmt.Printf("Length out: %d\n", len(out))
	for i := range out {
		fmt.Printf("Code %d: %064b\n", i, out[i])
	}

	// for i, result := range tests.out {
	// 	if out[i] != result {
	// 		t.Error(
	// 			"For id: ", i,
	// 			"expected result: ", result,
	// 			"got: ", out[i],
	// 		)
	// 	}
	// }
}

func TestEncode2(t *testing.T) {
	out := Decode(Encode(tests.in))

	for i, result := range tests.in {
		if out[i] != result {
			t.Error(
				"For id: ", i,
				"expected result: ", result,
				"got: ", out[i],
			)
		}
	}
}

func BenchmarkEncode(b *testing.B) { // Takes 130ns/op.
	b.ReportAllocs()
	b.ResetTimer()

	// Test the performance of the Encode process
	for i := 0; i < b.N; i++ {
		out := Encode(tests.in)
		_ = out
	}
}

func BenchmarkDecode(b *testing.B) { // Takes 340ns/op.
	in := Encode(tests.in)
	b.ReportAllocs()
	b.ResetTimer()

	// Test the performance of the Decode process
	for i := 0; i < b.N; i++ {
		out := Decode(in)
		_ = out
	}
}
