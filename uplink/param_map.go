package uplink

import (
	"math"
	"reflect"
	"strconv"
)

type paramMap map[string]int

func newParamMap(params []Parameter) paramMap {
	m := make(paramMap)
	for _, param := range params {
		m[param.Name] = param.RawValue
	}
	return m
}

func parseParams(params []Parameter, obj interface{}) {
	m := newParamMap(params)

	val := reflect.ValueOf(obj).Elem()
	typ := reflect.TypeOf(obj).Elem()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		rawValue := m[field.Tag.Get("name")]
		scale, err := strconv.ParseFloat(field.Tag.Get("scale"), 64)
		if err != nil {
			scale = 1.0
		}
		if field.Type.String() == "int" {
			val.Field(i).SetInt(int64(math.Round((float64(rawValue) * scale))))
		} else {
			if rawValue == -32768 {
				val.Field(i).SetFloat(math.NaN())
			} else {
				val.Field(i).SetFloat(float64(rawValue) * scale)
			}
		}
	}
}
