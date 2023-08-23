package runner

import (
	"reflect"

	"github.com/Brum3ns/encode/pkg/encoder"
)

func (r *Runner) worker(jobs <-chan string, result chan<- string) {
	var (
		lst    []string
		typ    reflect.Value
		method reflect.Value
	)
	switch {
	case len(r.conf.Option.Encode) > 0:
		lst = r.conf.Option.Encode
		typ = reflect.ValueOf(&encoder.Encoders{})

	case len(r.conf.Option.Decode) > 0:
		lst = r.conf.Option.Decode
		typ = reflect.ValueOf(&encoder.Decoders{})
	}
	for item := range jobs {
		value := reflect.ValueOf(item)

		for _, e := range lst {
			method = typ.MethodByName(e)
			value = method.Call([]reflect.Value{value})[0]
		}
		result <- value.Interface().(string)
	}
}
