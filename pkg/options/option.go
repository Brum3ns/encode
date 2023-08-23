package options

import (
	"bufio"
	"flag"
	"os"
	"strings"
)

type Options struct {
	Stdin   []string `errorcode:"1001"`
	Encode  []string `errorcode:"1002"`
	Decode  []string `errorcode:"1003"`
	Threads int      `errorcode:"1004"`
}

func Option() (*Configure, error) {
	opt := &Options{}

	flag.Func("e", "The encoder to be used for each input given by stdin", opt.AddEncoder)
	flag.Func("d", "The decoder to be used for each input given by stdin", opt.AddDecoder)
	flag.IntVar(&opt.Threads, "t", 42, "Threads to use")
	flag.Usage = opt.customUsage
	flag.Parse()

	if err := opt.stdin(); err != nil {
		return nil, err
	}

	return configure(opt)
}

func (opt *Options) stdin() error {
	if data, err := os.Stdin.Stat(); err != nil {
		return err

	} else if data.Mode()&os.ModeNamedPipe > 0 {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			opt.Stdin = append(opt.Stdin, s.Text())
		}
	}
	return nil
}

func (opt *Options) AddEncoder(s string) error {
	opt.Encode = SetUppercaseSlice(opt.Encode, s)
	return nil
}

func (opt *Options) AddDecoder(s string) error {
	opt.Decode = SetUppercaseSlice(opt.Decode, s)
	return nil
}

func SetUppercaseSlice(l []string, s string) []string {
	for _, i := range strings.Split(s, ",") {
		l = append(l, strings.Title(strings.ToLower(i)))
	}
	return l
}
