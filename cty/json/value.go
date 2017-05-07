package json

import (
	"bytes"

	"github.com/apparentlymart/go-cty/cty"
	"github.com/apparentlymart/go-cty/cty/convert"
)

// Marshal produces a JSON representation of the given value that can later
// be decoded into a value of the given type.
//
// A type is specified separately to allow for the given type to include
// cty.PseudoTypeDynamic to represent situations where any type is permitted
// and so type information must be included to allow recovery of the stored
// structure when decoding.
//
// The given type will also be used to attempt automatic conversions of any
// non-conformant types in the given value, although this will not always
// be possible. If the value cannot be made to be conformant then an error is
// returned, which may be a cty.PathError.
func Marshal(val cty.Value, t cty.Type) ([]byte, error) {
	errs := val.Type().TestConformance(t)
	if errs != nil {
		// Attempt a conversion
		var err error
		val, err = convert.Convert(val, t)
		if err != nil {
			return nil, err
		}
	}

	// From this point onward, val can be assumed to be conforming to t.

	buf := &bytes.Buffer{}
	var path cty.Path
	err := marshal(val, t, path, buf)

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Unmarshal decodes a JSON representation of the given value into a cty Value
// conforming to the given type.
//
// While decoding, type conversions will be done where possible to make
// the result conformant even if the types given in JSON are not exactly
// correct. If conversion isn't possible then an error is returned, which
// may be a cty.PathError.
func Unmarshal(buf []byte, t cty.Type) (cty.Value, error) {
	return cty.NilVal, nil
}