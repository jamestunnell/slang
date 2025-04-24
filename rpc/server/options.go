package server

import "time"

type Options struct {
	Port        int
	StartupWait time.Duration
}

type OptionsMod func(*Options)

const (
	DefaultPort = 0
	DefaultPath = "/slang/rpc"
	DefaultWait = 50 * time.Millisecond
)

func DefaultOptions() *Options {
	return &Options{
		Port:        0,
		StartupWait: DefaultWait,
	}
}

func WithPort(port int) OptionsMod {
	return func(opts *Options) {
		opts.Port = port
	}
}

func WithStartupWait(wait time.Duration) OptionsMod {
	return func(opts *Options) {
		opts.StartupWait = wait
	}
}
