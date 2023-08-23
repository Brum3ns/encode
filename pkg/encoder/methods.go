package encoder

import (
	"fmt"
	"reflect"
)

type Structure interface {
	Encoders | Decoders
}

func IsValid[T Structure](t T, l []string) bool {

	encode := reflect.TypeOf(&t)
	for _, name := range l {
		for i := 0; i < encode.NumMethod(); i++ {
			if name == encode.Method(i).Name {
				break
			}
			if i == encode.NumMethod()-1 {
				return false
			}
		}
	}
	return true
}

func Show[T Structure](t T) string {
	var v = "  ---"
	encode := reflect.TypeOf(&t)
	for i := 0; i < encode.NumMethod(); i++ {
		name := encode.Method(i).Name
		v += fmt.Sprintf("\n  %d. %s", (i + 1), name)
	}
	v += "\n  ---\n"
	return v
}

func SetNames[T Structure](t T) []string {
	var lst []string
	encode := reflect.TypeOf(&t)
	for i := 0; i < encode.NumMethod(); i++ {
		method := encode.Method(i)
		lst = append(lst, method.Name)
	}
	return lst
}
