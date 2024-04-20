package dugtrio

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockJsonStruct struct {
	V1 string `json:"value1,omitempty"`
	V2 string `json:"value2,omitempty"`
	V3 int    `json:"value3,omitempty"`
}

type mockNoneJsonStruct struct {
	V1 string
	V2 string
}

func TestReflectParseFiledJsonKey(t *testing.T) {
	f, _ := reflect.TypeOf(mockJsonStruct{}).FieldByName("V1")
	k := reflectParseFieldJsonKey(&f)
	assert.Equal(t, "value1", k, "Invalid Json key")
	f, _ = reflect.TypeOf(mockJsonStruct{}).FieldByName("V2")
	k = reflectParseFieldJsonKey(&f)
	assert.Equal(t, "value2", k, "Invalid Json key")
	f, _ = reflect.TypeOf(mockNoneJsonStruct{}).FieldByName("V1")
	k = reflectParseFieldJsonKey(&f)
	assert.Equal(t, "", k, "Invalid Json key")
	f, _ = reflect.TypeOf(mockNoneJsonStruct{}).FieldByName("V2")
	k = reflectParseFieldJsonKey(&f)
	assert.Equal(t, "", k, "Invalid Json key")
}

func TestReflectFindFiled(t *testing.T) {
	v := &mockJsonStruct{V1: "a", V2: "b", V3: 2}
	result := ReflectFindValueByName(v, "value1")
	assert.NotNil(t, result, "value1 is a valid field name.")
	assert.Equal(t, reflect.String, result.Type().Kind(), "value1 is a string")
	assert.Equal(t, "a", result.String(), "value1 is a string 'a'")
	result = ReflectFindValueByName(v, "value2")
	assert.NotNil(t, result, "value2 is a valid field name.")
	assert.Equal(t, reflect.String, result.Type().Kind(), "value2 is a string")
	assert.Equal(t, "b", result.String(), "value2 is a string 'b'")
	result = ReflectFindValueByName(v, "value3")
	assert.NotNil(t, result, "value3 is a valid field name.")
	assert.Equal(t, reflect.Int, result.Type().Kind(), "value3 is int")
	assert.Equal(t, int(2), int(result.Int()), "value3 is an INT number '2'")

	result = ReflectFindValueByName(v, "valueInvalid")
	assert.Nil(t, result, "valueInvalid is an invalid field name.")
}

func TestReflectCompare(t *testing.T) {
	v1 := &mockJsonStruct{V1: "a", V2: "b", V3: 2}
	v2 := &mockJsonStruct{V1: "a", V2: "c", V3: 3}
	r := reflectElementCompare(ReflectFindValueByName(v1, "V3"), ReflectFindValueByName(v2, "V3"))
	assert.Greater(t, 0, r, "v1.V3 < v2.V3")
	r = reflectElementCompare(ReflectFindValueByName(v1, "V1"), ReflectFindValueByName(v2, "V1"))
	assert.Equal(t, 0, r, "v1.V1 == v2.V1")
	r = reflectElementCompare(ReflectFindValueByName(v1, "V2"), ReflectFindValueByName(v2, "V2"))
	assert.Greater(t, 0, r, "v1.V2 < v2.V2")
}

type refVal struct {
	val1 string
}

func TestReflectUtil(t *testing.T) {
	v1 := refVal{val1: "abc"}
	v2 := refVal{val1: "ABC"}
	fmt.Printf("v1 = %s, v2 = %s\n", v1, v2)

	vstr1 := ReflectFindValueByName(v1, "val1")
	fmt.Printf("val = %v\n", vstr1)
	if vstr1.Comparable() {
		fmt.Printf("Comparable\n")
	}

	vstr2 := ReflectFindValueByName(v2, "val1")
	fmt.Printf("val = %v\n", vstr2)

	if vstr1.Equal(*vstr2) {
		fmt.Printf("equal")
	} else {
		fmt.Printf("!equal")
	}
}
