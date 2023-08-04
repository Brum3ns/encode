package options

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Brum3ns/encode/pkg/encoder"
)

type Options struct {
	Stdin   []string `errorcode:"1001"`
	Encode  []string `errorcode:"1002"`
	Threads int      `errorcode:"1003"`
}

func Option() (*Configure, error) {
	opt := &Options{}

	flag.Func("e", fmt.Sprintf("The encoder to be used for each input given by stdin\n%s", encoder.ListEncoders()), opt.StringToSlice)
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

func (opt *Options) StringToSlice(s string) error {
	for _, i := range strings.Split(s, ",") {
		opt.Encode = append(opt.Encode, strings.Title(strings.ToLower(i)))
	}
	return nil
}
