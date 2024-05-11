package utils

import "reflect"

// StructToMap
// @author: [Shansec](https://github.com/shansec)
// @function: StructToMap
// @description: 结构体转 Map
// @param: v interface{}
// @return: map[string]interface{}
func StructToMap(v interface{}) map[string]interface{} {
	vType := reflect.TypeOf(v)
	vValue := reflect.ValueOf(v)

	data := make(map[string]interface{})
	for i := 0; i < vType.NumField(); i++ {
		if vType.Field(i).Tag.Get("mapstructure") != "" {
			data[vType.Field(i).Tag.Get("mapstructure")] = vValue.Field(i).Interface()
		} else {
			data[vType.Field(i).Name] = vValue.Field(i).Interface()
		}
	}
	return data
}

func Pointer[T any](in T) (out *T) {
	return &in
}
