package encoder

import (
	"fmt"
	"reflect"
)

type Structure interface {
	Encoders | Decoders
}

func Show() {
	var (
		lst_encoders = GetMethods(Encoders{})
		lst_decoders = GetMethods(Decoders{})
		unique       = getUnique(lst_encoders, lst_decoders)
	)
	methods := lst_encoders
	for _, i := range lst_decoders {
		if !inSlice(i, methods) {
			methods = append(methods, i)
			continue
		}
	}
	fmt.Println("---")
	for id, method := range methods {
		supported := ""
		if inSlice(method, unique["encoder"]) {
			supported = "- (\033[33mencoder\033[0m)"
		} else if inSlice(method, unique["decoder"]) {
			supported = "- (\033[36mdecoder\033[0m)"
		}
		fmt.Printf("%d. %s %s\n", id, method, supported)
	}
	fmt.Println("---")
}

// Show all the encoders and decoders given in a map form (key = encoder/decoder, list = supported methods)
func getUnique(encoders, decoders []string) map[string][]string {
	var (
		uniuqe = make(map[string][]string)
		m      = map[string][]string{
			"encoder": encoders,
			"decoder": decoders,
		}
	)
	l := []string{"encoder", "decoder"}
	j := 1
	for i := 0; i < len(l); i++ {
		if i == 1 {
			j--
		}
		for _, method := range m[l[i]] {
			if !inSlice(method, m[l[j]]) {
				uniuqe[l[i]] = append(uniuqe[l[i]], method)
			}
		}
	}
	return uniuqe
}

// Check if a methods is shared within the encoder and the decoder
func inSlice(s string, l []string) bool {
	for _, i := range l {
		if i == s {
			return true
		}
	}
	return false
}

// Check if a method is valid within the given [struct]ure
func IsValidMethod[T Structure](t T, l []string) bool {
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

// Get all the supported encoder/decoder methods
func GetMethods[T Structure](t T) []string {
	var l []string
	encode := reflect.TypeOf(&t)
	for i := 0; i < encode.NumMethod(); i++ {
		l = append(l, encode.Method(i).Name)
	}
	return l
}

// Get all the method names
func GetMethodNames[T Structure](t T) []string {
	var lst []string
	encode := reflect.TypeOf(&t)
	for i := 0; i < encode.NumMethod(); i++ {
		method := encode.Method(i)
		lst = append(lst, method.Name)
	}
	return lst
}
