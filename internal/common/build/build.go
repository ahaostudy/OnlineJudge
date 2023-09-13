package build

import (
	"errors"
	"fmt"
	"reflect"
)

type Builder struct {
	err error
}

// Build
// 将in对象的所有字段解析到out对象，out必须为指针类型
// 其中字段类型要求为“相近”，如int和float都是数字类型，为相近类型
// 遇到不可转换、多余、缺失字段时会跳过而不会返回错误
// 不支持切片、数组等复杂类型的字段转换，会跳过
// Build完会返回Builder对象，支持链式调用
func (b *Builder) Build(in, out interface{}) *Builder {
	if b.err != nil {
		return b
	}

	inValue := getElem(reflect.ValueOf(in))
	outValue := getPtr(reflect.ValueOf(out))

	if inValue.Kind() != reflect.Struct || outValue.Kind() != reflect.Ptr || outValue.Elem().Kind() != reflect.Struct {
		fmt.Println(inValue.Kind() != reflect.Struct, outValue.Kind() != reflect.Ptr, outValue.Elem().Kind() != reflect.Struct)
		b.err = errors.New("in must be a structure and out must be a structure pointer")
		return b
	}

	inType := inValue.Type()
	outType := outValue.Elem().Type()

	for i := 0; i < inType.NumField(); i++ {
		inField := inType.Field(i)
		inFieldValue := inValue.Field(i)

		// 复杂结构直接跳过
		if isComplexStructures(inField.Type.Kind()) {
			continue
		}

		// 判断是否存在名称相同的字段，不存在直接跳过
		if _, ok := outType.FieldByName(inField.Name); !ok {
			continue
		}

		// 获取对应字段
		outFieldValue := outValue.Elem().FieldByName(inField.Name)

		// 判断是否可转换
		if isAssignable(inFieldValue.Type(), outFieldValue.Type()) {
			outFieldValue.Set(inFieldValue.Convert(outFieldValue.Type()))
		}
	}

	return b
}

func (b *Builder) Error() error {
	return b.err
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

func isComplexStructures(kind reflect.Kind) bool {
	switch kind {
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func, reflect.Interface:
		return true
	}
	return false
}

// 获取值类型
func getElem(val reflect.Value) reflect.Value {
	if val.Kind() != reflect.Ptr {
		return val
	}
	return getElem(val.Elem())
}

// 获取最后一层指针
func getPtr(val reflect.Value) reflect.Value {
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Ptr {
		return val
	}
	return getPtr(val.Elem())
}
