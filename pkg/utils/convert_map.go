/*
Copyright © 2025 Abinand P <abinand0911@gmail.com>
*/
package utils

import "fmt"

func ConvertInterfaceMap(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m := make(map[string]interface{})
		for k, v := range x {
			m[fmt.Sprint(k)] = ConvertInterfaceMap(v)
		}
		return m
	case []interface{}:
		for i, v := range x {
			x[i] = ConvertInterfaceMap(v)
		}
		return x
	default:
		return x
	}
}
