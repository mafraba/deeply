package deeply

import (
	"log"
	"reflect"
)

// Merge will merge two maps by imposing values of 'b' over 'a'
func Merge(a, b map[string]interface{}) map[string]interface{} {
	c := Copy(b)
	for k, av := range a {
		bv, bok := b[k]
		if bok {
			at := reflect.TypeOf(av)
			bt := reflect.TypeOf(bv)
			if at == bt {
				switch av.(type) {
				case map[string]interface{}:
					c[k] = Merge(av.(map[string]interface{}), bv.(map[string]interface{}))
				default:
					c[k] = bv
				}
			} else {
				log.Printf("Non matching types '%v' and '%v' (%s)", at, bt, k)
			}
		} else {
			c[k] = av
		}
	}
	return c
}

// Copy makes a deep copy of a map[string]interface{} (no cloning of pointers, slices or arrays ...)
func Copy(m map[string]interface{}) map[string]interface{} {
	c := map[string]interface{}{}
	for k, v := range m {
		switch v.(type) {
		case map[string]interface{}:
			c[k] = Copy(v.(map[string]interface{}))
		default:
			c[k] = v
		}
	}
	return c
}
