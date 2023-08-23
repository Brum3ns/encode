package options

import (
	"flag"
	"fmt"

	"github.com/Brum3ns/encode/pkg/encoder"
)

func (opt *Options) customUsage() {
	space_width := "\t"
	fmt.Printf("Usage: encoder -e/-d <type> [OPTIONS] ...\n%s\n", encoder.Show(encoder.Decoders{}))
	flag.VisitAll(func(f *flag.Flag) {
		var defaultValue string
		if len(f.DefValue) > 0 {
			defaultValue = fmt.Sprintf("(Default: \033[33m%v\033[0m)", f.DefValue)
		}

		//Print the helpmenu:
		fmt.Printf("  -%s%s %s\n", f.Name, (space_width + f.Usage), defaultValue)
	})
}
