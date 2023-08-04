package options

import (
	"flag"
	"fmt"
)

func (opt *Options) customUsage() {
	space_width := "\t"
	flag.VisitAll(func(f *flag.Flag) {
		var defaultValue string
		if len(f.DefValue) > 0 {
			defaultValue = fmt.Sprintf("(Default: \033[33m%v\033[0m)", f.DefValue)
		}

		//Print the helpmenu:
		fmt.Println("Usage: encoder -e <encode_method> [OPTIONS] ...")
		fmt.Printf("  -%s%s %s\n", f.Name, (space_width + f.Usage), defaultValue)
	})
}
