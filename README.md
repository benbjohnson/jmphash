jmphash
=======

Implementation of the [Jump Consistent Hash algorithm](http://arxiv.org/pdf/1406.2294v1.pdf) in Go.
This algorithm performs consistent hashing on integer keys and maps them
to integer buckets.


## Usage

To use jmphash, simply create a `Hasher` with the number of buckets you want
to map to and then call the `Hash()` function with your key. This function
will return the bucket that your key is mapped to.

```go
import "github.com/benbjohnson/jmphash"

func main() {
	// Create a hash with 100 buckets.
	h := jmphash.NewHasher(100)

	// Map keys to their appropriate buckets.
	bucket := h.Hash(12387)
}
```

