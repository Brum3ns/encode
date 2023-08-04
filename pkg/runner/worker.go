package runner

import (
	"reflect"

	"github.com/Brum3ns/encode/pkg/encoder"
)

func (r *Runner) worker(jobs <-chan string, result chan<- string) {
	var (
		encode = reflect.ValueOf(&encoder.Encoders{})
		method reflect.Value
	)
	for item := range jobs {
		encodedValue := reflect.ValueOf(item)

		for _, e := range r.conf.Option.Encode {
			method = encode.MethodByName(e)
			encodedValue = method.Call([]reflect.Value{encodedValue})[0]
		}
		result <- encodedValue.Interface().(string)
	}
}
