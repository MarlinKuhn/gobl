package dsig

import "errors"

// DigestAlgorithm determines the name of the algorithm used to generate the digest's
// value.
type DigestAlgorithm string

// Known list of digest algorithms supported.
const (
	DigestSHA256 DigestAlgorithm = "sha256"
)

// Digest defines a structure to hold a digest value including the algorithm used
// to generate it.
type Digest struct {
	// Algorithm stores the algorithm key that was used to generate the value.
	Algorithm DigestAlgorithm `json:"alg" jsonschema:"title=Algorithm"`

	// Value contains the Hexadecimal representation of the resulting hash
	// generated by the algorithm.
	Value string `json:"val" jsonschema:"title=Value"`
}

// Equals checks to ensure the current digest result matches that of the provided
// digest object. This will fail if the algorithms are different.
func (d *Digest) Equals(d2 *Digest) error {
	if d.Algorithm != d2.Algorithm {
		return errors.New("digest algorithm mismatch")
	}
	if d.Value != d2.Value {
		return errors.New("digest mismatch")
	}
	return nil
}
