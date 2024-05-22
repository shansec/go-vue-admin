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

func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// MaheHump 将字符串转换为驼峰命名
func MaheHump(s string) string {
	words := strings.Split(s, "-")

	for i := 1; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}

	return strings.Join(words, "")
}

// 随机字符串
func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[RandomInt(0, len(letters))]
	}
	return string(b)
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}