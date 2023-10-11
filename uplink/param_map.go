package uplink

import "math"

type paramMap map[string][]int

func newParamMap(params []Parameter) paramMap {
	m := make(paramMap)
	for _, param := range params {
		m.insert(param.Title, param.RawValue)
	}
	return m
}

func (m *paramMap) insert(key string, value int) {
	(*m)[key] = append((*m)[key], value)
}

func (m *paramMap) int(key string, optionalIndex ...int) int {
	i := index(optionalIndex...)
	if i >= len((*m)[key]) {
		return 0
	}
	return (*m)[key][i]
}

func (m *paramMap) float32(key string, optionalIndex ...int) float32 {
	value := float32(m.int(key, optionalIndex...))
	if value == -32768 {
		return float32(math.NaN())
	}
	return value
}

func index(index ...int) int {
	if len(index) > 0 {
		return index[0]
	}
	return 0
}
