// ex13.1 provides a deep equivalence relation for arbitrary values.
package equalish

import (
	"reflect"
	"unsafe"
)

// Consider numbers equalish if their difference is less than one part in
// <multiplier>.
const multiplier = 1000000000

func numbersEqualish(x, y float64) bool {
	// This isn't just a shortcut. We need a special case for zero.
	if x == y {
		return true
	}
	var diff float64
	if x > y {
		diff = x - y
	} else {
		diff = y - x
	}
	d := diff * multiplier
	if d < x && d < y {
		return true
	}
	return false
}

func equalish(x, y reflect.Value, seen map[comparison]bool) bool {
	if !x.IsValid() || !y.IsValid() {
		return x.IsValid() == y.IsValid()
	}
	if x.Type() != y.Type() {
		return false
	}

	// cycle check
	if x.CanAddr() && y.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		yptr := unsafe.Pointer(y.UnsafeAddr())
		if xptr == yptr {
			return true // identical references
		}
		c := comparison{xptr, yptr, x.Type()}
		if seen[c] {
			return true // already seen
		}
		seen[c] = true
	}
	switch x.Kind() {
	case reflect.Bool:
		return x.Bool() == y.Bool()

	case reflect.String:
		return x.String() == y.String()

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64:
		return numbersEqualish(float64(x.Int()), float64(y.Int()))

	case reflect.Uintptr:
		return x.Uint() == y.Uint()

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return numbersEqualish(float64(x.Uint()), float64(y.Uint()))

	case reflect.Float32, reflect.Float64:
		return numbersEqualish(float64(x.Float()), float64(y.Float()))

	case reflect.Complex64, reflect.Complex128:
		realEqualish := numbersEqualish(float64(real(x.Complex())), float64(real(y.Complex())))
		imagEqualish := numbersEqualish(float64(imag(x.Complex())), float64(imag(y.Complex())))
		return realEqualish && imagEqualish
	case reflect.Chan, reflect.UnsafePointer, reflect.Func:
		return x.Pointer() == y.Pointer()

	case reflect.Ptr, reflect.Interface:
		return equalish(x.Elem(), y.Elem(), seen)

	case reflect.Array, reflect.Slice:
		if x.Len() != y.Len() {
			return false
		}
		for i := 0; i < x.Len(); i++ {
			if !equalish(x.Index(i), y.Index(i), seen) {
				return false
			}
		}
		return true

	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if !equalish(x.Field(i), y.Field(i), seen) {
				return false
			}
		}
		return true

	case reflect.Map:
		if x.Len() != y.Len() {
			return false
		}
		for _, k := range x.MapKeys() {
			if !equalish(x.MapIndex(k), y.MapIndex(k), seen) {
				return false
			}
		}
		return true
	}
	panic("unreachable")
}

// Equalish reports whether x and y are deeply equal, with numeric values
// differing by less than one part in a billion.
//
// Map keys are always compared with ==, not deeply.
// (This matters for keys containing pointers or interfaces.)
func Equalish(x, y interface{}) bool {
	seen := make(map[comparison]bool)
	return equalish(reflect.ValueOf(x), reflect.ValueOf(y), seen)
}

type comparison struct {
	x, y unsafe.Pointer
	t    reflect.Type
}
