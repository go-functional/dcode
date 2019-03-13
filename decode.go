package dcode

import (
	"fmt"
	"reflect"
)

// Decode decodes b into i using d
func Decode(b []byte, d Decoder, i interface{}) error {
	ret, err := d.call(b)
	if err != nil {
		return err
	}
	inputType := reflect.TypeOf(i)
	if inputType.Kind() != reflect.Ptr {
		return fmt.Errorf("Given %s type is not a pointer", inputType.Name())
	}
	inputTypeName := inputType.Elem().Name()
	returnedType := reflect.TypeOf(ret)
	if inputTypeName != returnedType.Name() {
		return fmt.Errorf(
			"Got type '%s', but expected '%s'",
			returnedType.Name(),
			inputType.Name(),
		)
	}
	// Need to call Elem to traverse the pointer
	inputVal := reflect.ValueOf(i).Elem()
	retVal := reflect.ValueOf(ret)
	if !inputVal.CanSet() {
		return fmt.Errorf("Couldn't set the decoded value")
	}
	inputVal.Set(retVal)
	return nil
}

// MapPair is a value to pass into the Map function. Create one
// of these with the Pair function, and see the documentation
// under the Map function for where this is used
type MapPair struct {
	decoder Decoder
	iface   interface{}
}

// Pair returns a new MapPair
func Pair(d Decoder, iface interface{}) MapPair {
	return MapPair{decoder: d, iface: iface}
}

// Map decodes separate JSON fields into separate values,
// using separate decoders, all from the same JSON.
//
// For example, if you have the following JSON:
//
//	json := `{"field1": 123, "field2": "456"}`
//
// And you have the following decoders:
//
//	dcode1 := Field("field1", Int())
//	dcode2 := Field("field2", String())
//
// You can do the following:
//
//	stucco := struct{
//		field1 int
//		field2 string
//	}
//
//	// check the error here!
//	Map(
//		[]byte(json),
//		Pair(dcode1, &stucco.field1),
//		Pair(dcode2, &stucco.field2),
//	)
//
// Note: you can decode JSON into structs using this function,
// but if you're trying to decode tons of fields from JSON
// to struct fields, you might be better off using
// encoding/json!
func Map(
	b []byte,
	pairs ...MapPair,
) error {
	for _, pair := range pairs {
		if err := Decode(b, pair.decoder, pair.iface); err != nil {
			return err
		}
	}
	return nil
}

// OneOf tries to decode b into i using every decoder in dList.
//
// It returns nil after the first decoder succeeds. If none
// succeeded, returns an error
func OneOf(b []byte, dList []Decoder, i interface{}) error {
	for _, d := range dList {
		if err := Decode(b, d, i); err == nil {
			return nil
		}
	}
	return fmt.Errorf("Tried %d decoders, none decoded", len(dList))
}
