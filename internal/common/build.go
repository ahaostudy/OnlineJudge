package common

import (
	"errors"
	"reflect"
)

// Build
// 将in对象的所有字段解析到out对象，out必须为指针类型
// 其中字段类型要求为“相近”，如int和float都是数字类型，为相近类型
// 遇到不可转换、多余、缺失字段时会跳过而不会返回错误
func Build(in, out any) error {
	inValue := reflect.ValueOf(in)
	outValue := reflect.ValueOf(out)

	if inValue.Kind() == reflect.Ptr {
		inValue = inValue.Elem()
	}
	if inValue.Kind() != reflect.Struct || outValue.Kind() != reflect.Ptr || outValue.Elem().Kind() != reflect.Struct {
		return errors.New("in must be a structure and out must be a structure pointer")
	}

	inType := inValue.Type()
	outType := outValue.Elem().Type()

	for i := 0; i < inType.NumField(); i++ {
		inField := inType.Field(i)
		inFieldValue := inValue.Field(i)

		for j := 0; j < outType.NumField(); j++ {
			outField := outType.Field(j)
			outFieldValue := outValue.Elem().Field(j)

			if inField.Name == outField.Name && isAssignable(inFieldValue.Type(), outFieldValue.Type()) {
				outFieldValue.Set(inFieldValue.Convert(outFieldValue.Type()))
				break
			}
		}
	}

	return nil
}

// 判断两个类型是否可以转换
func isAssignable(from, to reflect.Type) bool {
	fromKind := from.Kind()
	toKind := to.Kind()

	if fromKind == toKind {
		return true
	}

	// 检查数字类型的互相转换
	if isNumericKind(fromKind) && isNumericKind(toKind) {
		return true
	}

	return false
}

// 判断是否为数字类型
func isNumericKind(kind reflect.Kind) bool {
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return true
	}
	return false
}
