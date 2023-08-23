package options

import (
	"errors"
	"log"
	"reflect"
	"strconv"

	"github.com/Brum3ns/encode/pkg/encoder"
	"github.com/Brum3ns/encode/pkg/fail"
)

type Configure struct {
	Option  *Options
	Encoder *encoder.Encoders
	Decoder *encoder.Decoders
}

func configure(opt *Options) (*Configure, error) {
	conf := &Configure{
		Option:  opt,
		Encoder: encoder.NewEncoder(),
		Decoder: encoder.NewDecoder(),
	}

	//Check so that not both the encoder and decoder are set:
	if (len(opt.Encode) > 0 && len(opt.Decode) > 0) || len(opt.Decode) == 0 && len(opt.Encode) == 0 {
		log.Fatal("An encoder or decoder has to be set. Can't set both at once.")
	}

	v := reflect.ValueOf(conf.Option).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		item := t.Field(i)

		//Validation error detected for user input, return error to the user screen:
		if exist, ok := conf.MethodCall(item.Name); exist && !ok {
			if errcode, ok := strconv.Atoi(item.Tag.Get("errorcode")); ok == nil {
				return nil, errors.New(fail.ERRORCODES[errcode])
			} else {
				log.Panicf("Can't convert errorcode value \"%v\" for flag \"%s\".\n", errcode, item.Name)
			}
		}
	}

	return conf, nil
}

func (conf *Configure) MethodCall(name string) (bool, bool) {
	if v := reflect.ValueOf(conf).MethodByName(name); v.IsValid() {
		return true, v.Call(nil)[0].Interface() == true
	}
	return false, false
}

func (conf *Configure) Threads() bool {
	return conf.Option.Threads > 0 && conf.Option.Threads < 10000
}

func (conf *Configure) Stdin() bool {
	return len(conf.Option.Stdin) > 0
}

func (conf *Configure) Encode() bool {
	return len(conf.Option.Decode) > 0 || encoder.IsValidMethod(*conf.Encoder, conf.Option.Encode)
}

func (conf *Configure) Decode() bool {
	return len(conf.Option.Encode) > 0 || encoder.IsValidMethod(*conf.Decoder, conf.Option.Decode)
}
