package jmphash_test

import (
	"math"
	"testing"
	"testing/quick"

	. "github.com/benbjohnson/jmphash"
)

func TestHasher_Hash_MaxK1_N1(t *testing.T)  { testHasher_Hash(t, 1, 1) }
func TestHasher_Hash_MaxK1_N10(t *testing.T) { testHasher_Hash(t, 1, 10) }

func TestHasher_Hash_MaxK10_N10(t *testing.T)   { testHasher_Hash(t, 10, 10) }
func TestHasher_Hash_MaxK100_N10(t *testing.T)  { testHasher_Hash(t, 100, 10) }
func TestHasher_Hash_MaxK1000_N10(t *testing.T) { testHasher_Hash(t, 1000, 10) }

func TestHasher_Hash_MaxK100_N100(t *testing.T) { testHasher_Hash(t, 100, 100) }

// Ensure that the hash assign to a bucket greater than zero and less than the
// maximum bucket count.
func testHasher_Hash(t *testing.T, maxKey uint64, maxN uint32) {
	// Randomly generate bucket counts and key sets.
	err := quick.Check(func(bucketN uint32, keys []uint64) bool {
		bucketN %= maxN

		// Create a new hasher.
		h := New(int(bucketN))

		// Hash and verify each key is within the appropriate range.
		for _, key := range keys {
			key %= maxKey
			bucket := h.Hash(key)

			if bucket < 0 || bucket >= h.N() {
				t.Errorf("invalid bucket: %d; k=%d; n=%d", bucket, key, h.N())
			}
		}
		return true
	}, nil)

	if err != nil {
		t.Error(err)
	}
}

func TestHasher_Hash_Move_N1(t *testing.T)    { testHasher_Hash_Move(t, 1) }
func TestHasher_Hash_Move_N10(t *testing.T)   { testHasher_Hash_Move(t, 10) }
func TestHasher_Hash_Move_N100(t *testing.T)  { testHasher_Hash_Move(t, 100) }
func TestHasher_Hash_Move_N1000(t *testing.T) { testHasher_Hash_Move(t, 1000) }

// Ensure that changing the shard count will redistribute the appropriate number of keys.
func testHasher_Hash_Move(t *testing.T, maxN uint32) {
	err := quick.Check(func(n0 uint32, n1 uint32) bool {
		// Create a new hasher.
		n0 = (n0 % maxN) + 1
		n1 = (n1 % maxN) + 1
		h0, h1 := New(int(n0)), New(int(n1))

		// Determine the number of keys that have to move.
		var moved int
		total := 10000
		for i := 0; i < total; i++ {
			b0, b1 := h0.Hash(uint64(i)), h1.Hash(uint64(i))
			if b0 != b1 {
				moved++
			}
		}

		// Verify that the appropriate percentage of keys have moved.
		pct := float64(moved) / float64(total)
		exp := float64(1) - (math.Min(float64(n0), float64(n1)) / math.Max(float64(n0), float64(n1)))
		if math.Abs(pct-exp) > 0.1 {
			t.Errorf("invalid move: %0.2f%%; expected %0.2f%%; <n=%d→%d>", pct*100, exp*100, n0, n1)
		}

		return true
	}, nil)

	if err != nil {
		t.Error(err)
	}

}

func BenchmarkHasherHashN1(b *testing.B)     { benchmarkHasherHash(b, 1) }
func BenchmarkHasherHashN5(b *testing.B)     { benchmarkHasherHash(b, 5) }
func BenchmarkHasherHashN10(b *testing.B)    { benchmarkHasherHash(b, 10) }
func BenchmarkHasherHashN100(b *testing.B)   { benchmarkHasherHash(b, 100) }
func BenchmarkHasherHashN1000(b *testing.B)  { benchmarkHasherHash(b, 1000) }
func BenchmarkHasherHashN10000(b *testing.B) { benchmarkHasherHash(b, 10000) }

func benchmarkHasherHash(b *testing.B, n int) {
	h := New(n)
	for i := 0; i < b.N; i++ {
		if x := h.Hash(uint64(i)); x > n {
			b.Fatal("invalid hash:", x)
		}
	}
}
