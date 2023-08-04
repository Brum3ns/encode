package runner

func (r *Runner) job(jobs chan<- string) {
	for _, i := range r.conf.Option.Stdin {
		jobs <- i
	}
}
