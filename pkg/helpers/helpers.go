package helpers

import "reflect"

// 判断各类型值是否为空
func Empty(val interface{}) bool {
	if val == nil {
		return true
	}
	v := reflect.ValueOf(val)
	switch v.Kind() {
	// 字符串和数组，判断长度即可
	case reflect.String, reflect.Array:
		return v.Len() == 0
	// map和切片，判断长度并且判断是否开辟空间
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	// bool类型，取反即可
	case reflect.Bool:
		return !v.Bool()
	// int类型，大于0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	// 对象、指针类型， 判断是否开辟空间
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	// 其他类型，判断是否等于0值
	return reflect.DeepEqual(val, reflect.Zero(v.Type()).Interface())
}
