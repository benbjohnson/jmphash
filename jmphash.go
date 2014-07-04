package jmphash

// Hasher represents a Jump Consistent Hasher.
type Hasher struct {
	n int32
}

// NewHasher returns a new instance of Hasher.
// If the number of of buckets is less than or equal to zero then one bucket is used.
func NewHasher(n int) *Hasher {
	if n <= 0 {
		n = 1
	}
	return &Hasher{int32(n)}
}

// N returns the number of buckets the hasher can assign to.
func (h *Hasher) N() int { return int(h.n) }

// Hash returns the integer hash for the given key.
func (h *Hasher) Hash(key uint64) int {
	b, j, n := int64(-1), int64(0), int64(h.n)
	for j < n {
		b = j
		key = key*uint64(2862933555777941757) + 1
		j = int64(float64(b+1) * (float64(int64(1)<<31) / float64((key>>33)+1)))
	}
	return int(b)
}
