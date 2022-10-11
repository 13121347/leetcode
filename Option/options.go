package Option

type Option func(opts *Options)

type Options struct {
	a int
	b int
	c int
}

func loadOptions(options ...Option) *Options {
	opts := new(Options)

	for _, optionfunc := range options {
		optionfunc(opts)
	}
	return opts
}
