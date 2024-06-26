package utils

import "reflect"

// Função para converter uma struct para map[string]interface{}
func StructToMap(s interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(s)

	// Verifique se s é um ponteiro e, se for, obtenha o valor subjacente
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Verifique se s é uma struct
	if val.Kind() != reflect.Struct {
		panic("StructToMap: not a struct")
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		fieldName := typ.Field(i).Name
		fieldValue := val.Field(i).Interface()
		result[fieldName] = fieldValue
	}

	return result
}

// Função para converter um map[string]interface{} para uma struct
func MapToStruct(m map[string]interface{}, s interface{}) {
	v := reflect.ValueOf(s).Elem()
	for key, value := range m {
		field := v.FieldByName(key)
		if field.IsValid() && field.CanSet() {
			field.Set(reflect.ValueOf(value))
		}
	}
}
