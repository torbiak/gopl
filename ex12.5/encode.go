// ex12.5 is a codec for json.
package json

import (
	"bytes"
	"fmt"
	"reflect"
)

// Marshal encodes a Go value to json.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v), 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// encode writes to buf a json representation of v.
func encode(buf *bytes.Buffer, v reflect.Value, indent int) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("null")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem(), indent)

	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteByte('[')
		indent += 1
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				fmt.Fprintf(buf, "\n%*s", indent, "")
			}
			if err := encode(buf, v.Index(i), indent); err != nil {
				return err
			}
			if i != v.Len()-1 {
				buf.WriteByte(',')
			}
		}
		buf.WriteByte(']')

	case reflect.Struct: // ((name value) ...)
		buf.WriteByte('{')
		indent += 1
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				fmt.Fprintf(buf, "\n%*s", indent, "")
			}
			start := buf.Len()
			fmt.Fprintf(buf, "%q: ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i), indent+buf.Len()-start); err != nil {
				return err
			}
			if i != v.NumField()-1 {
				buf.WriteByte(',')
			}
		}
		buf.WriteByte('}')

	case reflect.Map: // ((key value) ...)
		buf.WriteByte('{')
		indent += 1
		for i, key := range v.MapKeys() {
			if i > 0 {
				fmt.Fprintf(buf, "\n%*s", indent, "")
			}
			start := buf.Len()
			if err := encode(buf, key, 0); err != nil {
				return err
			}
			buf.WriteString(": ")
			if err := encode(buf, v.MapIndex(key), indent+buf.Len()-start); err != nil {
				return err
			}
			if i != len(v.MapKeys())-1 {
				buf.WriteByte(',')
			}
		}
		buf.WriteByte('}')

	case reflect.Bool: // true | false
		if v.Bool() {
			buf.WriteString("true")
		} else {
			buf.WriteString("false")
		}

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%g", v.Float())

	default: // chan, func, complex
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}
