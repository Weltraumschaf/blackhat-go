package proxy

import "github.com/urfave/cli/v2"

type options struct {
}

func newOptions(c *cli.Context) *options {
	return &options{}
}
