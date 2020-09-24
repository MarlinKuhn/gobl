package gobl

import (
	"fmt"
	"strconv"
	"strings"
)

// Amount represents a quantity with decimal places that will not suffer
// rounding errors like traditional floats.
// Use cases are assumed to be within the "human managable domain", i.e.
// for dealing with counts, money, rates, short distances, etc.
// Implementation is inspired by https://github.com/shopspring/decimal, but
// simplified to account for the expectations of GoBL.
type Amount struct {
	value int64
	exp   uint32
}

// NewAmount instantiates a new Amount object with the provided value
// and exponential.
func NewAmount(val int64, exp uint32) Amount {
	return Amount{val, exp}
}

// NewAmountFromString takes the provided string and tries to convert it
// into an amount object. Strings must be in a simplified format with no
// commas and a single `.` to seperate the decimal places. Numbers are
// expected to have a fixed number of decimal places, so if your dealing
// with a string like `"12.000"`, the accuracy will be assumed to be 3
// decimal places.
// If you're dealing with numbers from humans which may contain symbols,
// commans, european style fullstops, undescores, etc. then you should use
// the `NewAmountFromHumanString` method.
func NewAmountFromString(val string) (Amount, error) {
	a := Amount{}
	x := strings.Split(val, ".")
	l := len(x)
	if l > 2 {
		return a, fmt.Errorf("amount must contain 0 or 1 decimal separators: %v", val)
	}

	// Parse the "major" part
	v, err := strconv.ParseInt(x[0], 10, 64)
	if err != nil {
		return a, fmt.Errorf("invalid major number '%v', %w", val, err)
	}
	e := uint32(0)
	v2 := int64(0)

	// Parse the decimal places (if present)
	if l == 2 {
		v2, err = strconv.ParseInt(x[1], 10, 64)
		if err != nil {
			return a, fmt.Errorf("invalid decimal number '%v', %w", val, err)
		}
		e = uint32(len(x[1]))
		v = v * intPow(10, e)
		v += v2
	}

	// Prepare the result
	a.value = v
	a.exp = e
	return a, nil
}

// Add will add the two amounts together using the base's exponential
// value for the resulting new amount.
func (a Amount) Add(a2 Amount) Amount {
	a2 = a2.Rescale(a.exp)
	return NewAmount(a.value+a2.value, a.exp)
}

// Subtract takes away the amount provided from the base.
func (a Amount) Subtract(a2 Amount) Amount {
	a2 = a2.Rescale(a.exp)
	return NewAmount(a.value-a2.value, a.exp)
}

// Multiply our base amount by the provided amount.
func (a Amount) Multiply(a2 Amount) Amount {
	return Amount{
		value: (a.value * a2.value) / intPow(10, a2.exp),
		exp:   a.exp,
	}
}

// Divide our base amount by the provided amount.
func (a Amount) Divide(a2 Amount) Amount {
	return Amount{
		value: (a.value * intPow(10, a2.exp)) / a2.value,
		exp:   a.exp,
	}
}

// Compare two amounts and return an integer value according to the
// sign of the difference:
//
//   -1 if a <  a2
//    0 if a == a2
//    1 if a >  a2
//
func (a Amount) Compare(a2 Amount) int {
	a, a2 = rescaleAmountPair(a, a2)
	if a.value < a2.value {
		return -1
	}
	if a.value > a2.value {
		return 1
	}
	return 0
}

// Equals returns true if the two amounts represent the same value,
// regardless of the exponential.
func (a Amount) Equals(a2 Amount) bool {
	return a.Compare(a2) == 0
}

// Rescale will multiply or divide the amount's value to match the
// provided exponential. This method will not round values, value
// could be lost during conversion.
func (a Amount) Rescale(exp uint32) Amount {
	if a.exp > exp {
		// need to divide
		e := a.exp - exp
		v := a.value / intPow(10, e)
		return Amount{v, exp}
	}
	if a.exp < exp {
		// need to multiply
		e := exp - a.exp
		v := a.value * intPow(10, e)
		return Amount{v, exp}
	}
	return a
}

// String returns the simplified string amount.
func (a Amount) String() string {
	if a.exp == 0 {
		return fmt.Sprintf("%d", a.value)
	}
	p := intPow(10, a.exp)
	v1 := a.value / p
	v2 := a.value - (v1 * p)
	fmt.Printf("AV: %v, V1: %v, V2: %v\n", a.value, v1, v2)
	return fmt.Sprintf("%d.%0*d", v1, a.exp, v2)
}

// MarshalJSON provides the byte value of the amount. See also the
// String() method.
// We always add quotes around values as JSON numbers do not
// guarantee that tailing 0s will be maintained. It's important
// to remember that amounts are typically for humans, and thus
// it makes sense to consider them as strings.
func (a Amount) MarshalJSON() ([]byte, error) {
	str := "\"" + a.String() + "\""
	return []byte(str), nil
}

// UnmarshalJSON will decode the amount value, even if it is quoted
// as a string.
func (a *Amount) UnmarshalJSON(value []byte) error {
	if string(value) == "null" {
		return nil
	}

	str, err := unquote(value)
	if err != nil {
		return fmt.Errorf("decoding string `%s`: %w", value, err)
	}

	amount, err := NewAmountFromString(string(str))
	if err != nil {
		return fmt.Errorf("decoding string `%s`: %w", str, err)
	}
	*a = amount

	return nil
}

func unquote(value []byte) ([]byte, error) {
	// If the amount is quoted, strip the quotes
	if len(value) > 2 && value[0] == '"' && value[len(value)-1] == '"' {
		value = value[1 : len(value)-1]
	}
	return value, nil
}

func rescaleAmountPair(a, a2 Amount) (Amount, Amount) {
	// Take the largest exp
	exp := a.exp
	if a2.exp > exp {
		exp = a2.exp
	}
	return a.Rescale(exp), a2.Rescale(exp)
}

func intPow(base int, exp uint32) int64 {
	out := int64(1)
	for exp != 0 {
		out *= int64(base)
		exp--
	}
	return out
}
