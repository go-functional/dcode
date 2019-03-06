package dcode

import (
	"fmt"
	"reflect"
)

type Decoder func([]byte) (interface{}, error)

func Decode(b []byte, d Decoder, i interface{}) error {
	ret, err := d(b)
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
