// The runner is the primary process of the tool
package runner

import (
	"fmt"
	"os"
	"time"

	"github.com/Brum3ns/encode/pkg/config"
)

type Runner struct {
	conf *config.Configure
}

func NewRunner(conf *config.Configure) error {
	runner := &Runner{
		conf: conf,
	}

	total := len(runner.conf.Option.Stdin)

	jobs := make(chan string)
	results := make(chan string, total)

	for t := 0; t <= runner.conf.Option.Threads; t++ {
		go runner.worker(jobs, results)
	}

	//Give job to workers
	runner.job(jobs)

	//Super cool loading bar ^-^
	if total > 10000 {
		go loadingbar()
	}

	//Listener
	for i := 0; i < len(runner.conf.Option.Stdin); i++ {
		fmt.Println(<-results)
	}
	close(jobs)
	close(results)

	return nil
}
func loadingbar() {
	l := []string{"⠁", "⠃", "⠇", "⠧", "⠷", "⠿", "⠾", "⠼", "⠸", "⠘", "⠈", "⠉"}
	for i := 0; true; i++ {
		if i >= len(l) {
			i = 0
		}
		fmt.Fprint(os.Stderr, (l[i] + "\r"))
		time.Sleep(42 * time.Millisecond)
	}
}
