package util

import (
	"errors"
	"fmt"
	"io"
	"reflect"
)

type stringValue struct {
	key       string
	value     reflect.Value
	omitEmpty bool
}

type stringValueArray []stringValue

func writeStruct(w io.Writer, val reflect.Value) (err error) {
	_, err = fmt.Fprint(w, "d")
	if err != nil {
		return
	}

	typ := val.Type()
	numFields := val.NumField()
	svList := make(stringValueArray, numFields)

	for i := 0; i < numFields; i++ {
		field := typ.Field(i)
		bencodeKey(field, &svList[i])
		// The tag `bencode:"-"` should mean that this field must be ignored
		// See https://golang.org/pkg/encoding/json/#Marshal or https://golang.org/pkg/encoding/xml/#Marshal
		// We set a zero value so that it is ignored by the writeSVList() function
		if svList[i].key == "-" {
			svList[i].value = reflect.Value{}
		} else {
			svList[i].value = val.Field(i)
		}
	}

	err = writeSVList(w, svList)
	if err != nil {
		return
	}

	_, err = fmt.Fprint(w, "e")
	if err != nil {
		return
	}
	return
}

func writeValue(w io.Writer, val reflect.Value) (err error) {
	if !val.IsValid() {
		err = errors.New("can't write null value")
		return
	}

	fmt.Print("test 11111\n")
	switch v := val; v.Kind() {
	case reflect.String:
		s := v.String()
		fmt.Printf("str1\n")
		_, err = fmt.Fprintf(w, "%d:%s", len(s), s)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Printf("int1\n")
		_, err = fmt.Fprintf(w, "i%de", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		_, err = fmt.Fprintf(w, "i%de", v.Uint())
	case reflect.Interface:
		err = writeValue(w, v.Elem())
	case reflect.Struct:
		err = writeStruct(w, v)
	}
	return
}

func Marshal(w io.Writer, val interface{}) error {
	return writeValue(w, reflect.ValueOf(val))
}
