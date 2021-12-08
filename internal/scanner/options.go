package scanner

import "github.com/urfave/cli/v2"

type options struct {
	target string
	start  int
	end    int
}

func newOptions(c *cli.Context) *options {
	return &options{
		target: c.String("target"),
		start:  c.Int("start"),
		end:    c.Int("end")}
}

func (o *options) getTarget() string {
	return o.target
}

func (o *options) getStart() int {
	return o.start
}
func (o *options) getEnd() int {
	return o.end
}
