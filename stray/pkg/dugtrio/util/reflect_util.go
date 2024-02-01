package util

import (
	"reflect"
	"strconv"
	"strings"
)

// The sample of the `fullTag` String: "name,omitempty"
// It's the golang struct tags, for instance:
//
//	````
//	type testPerson struct {
//		Name        string `json:"name,omitempty"`
//	}
//
// ````
// In this example, the "name,omitempty" is the full tag the 'name' is the json key
func reflectParseFieldJsonKey(f *reflect.StructField) string {
	if f == nil {
		return ""
	}
	fullTag := f.Tag.Get("json")
	if fullTag == "" {
		return ""
	}
	tags := strings.Split(fullTag, ",")
	if len(tags) <= 0 {
		return fullTag
	}
	return tags[0]
}

func ReflectFindValueByName[T any](e T, n string) *reflect.Value {
	t := reflect.TypeOf(e)
	if t == nil {
		return nil
	}

	k := t.Kind()
	if reflect.Struct != k && k != reflect.Pointer {
		// The input source element is not an object or point
		return nil
	}

	v := reflect.ValueOf(e)
	if !v.IsValid() {
		return nil
	}

	ev := v
	if k == reflect.Pointer {
		ev = v.Elem()
		if !ev.IsValid() {
			return nil
		}
		t = ev.Type()
		if t == nil {
			return nil
		}
		if reflect.Struct != t.Kind() {
			// The input source element is a point but it did not point to an object
			return nil
		}
	}

	for i := 0; i < t.NumField(); i++ {
		// find by name
		f := t.Field(i)
		if f.Name == n {
			fv := ev.Field(i)
			if fv.IsValid() {
				return &fv
			} else {
				continue
			}
		}

		// find via json tag
		jsonKey := reflectParseFieldJsonKey(&f)
		if jsonKey == n {
			fv := ev.Field(i)
			if fv.IsValid() {
				return &fv
			} else {
				continue
			}
		}
	}
	return nil
}

func reflectElementCompare(v1, v2 *reflect.Value) int {
	if v1 == nil && v2 == nil {
		return 0
	}
	if v1 == nil && v2 != nil {
		return 1
	}
	if v1 != nil && v2 == nil {
		return -1
	}

	typeV1 := v1.Type()
	typeV2 := v2.Type()
	if typeV1 == nil || typeV2 == nil {
		return 0
	}
	if typeV1.Kind() != typeV2.Kind() {
		// The types of Value1 & Value2 are different.
		// It cannot be compared.
		return 0
	}
	switch typeV1.Kind() {
	case reflect.String:
		return strings.Compare(v1.String(), v2.String())
	case reflect.Bool:
		boolV1 := v1.Bool()
		boolV2 := v2.Bool()
		if boolV1 == boolV2 {
			return 0
		} else if boolV1 {
			return 1
		} else {
			return -1
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintV1 := v1.Uint()
		uintV2 := v2.Uint()
		if uintV1 == uintV2 {
			return 0
		} else if uintV1 > uintV2 {
			return 1
		} else {
			return -1
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intV1 := v1.Int()
		intV2 := v2.Int()
		if intV1 == intV2 {
			return 0
		} else if intV1 > intV2 {
			return 1
		} else {
			return -1
		}
	}
	return 0
}

func reflectElementMatch(v *reflect.Value, m string, ignoreCase bool, matchAll bool) bool {
	if v == nil || m == "" {
		return false
	}

	t := v.Type()
	if t == nil {
		return false
	}

	src := ""
	switch t.Kind() {
	case reflect.String:
		src = v.String()
	case reflect.Bool:
		src = strconv.FormatBool(v.Bool())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		src = strconv.FormatUint(v.Uint(), 10)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		src = strconv.FormatInt(v.Int(), 10)
	}

	if matchAll && ignoreCase && strings.EqualFold(src, m) {
		return true
	} else if matchAll && !ignoreCase && src == m {
		return true
	}
	// TODO !matchAll cases
	return false
}
