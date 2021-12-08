package scanner

import "github.com/urfave/cli/v2"

type options struct {
	targetHost      string
	startPort       int
	endPort         int
	showClosedPorts bool
}

func newOptions(c *cli.Context) *options {
	return &options{
		targetHost:      c.String("target"),
		startPort:       c.Int("start"),
		endPort:         c.Int("end"),
		showClosedPorts: c.Bool("closed")}
}

func (o *options) getTargetHost() string {
	return o.targetHost
}

func (o *options) getStartPort() int {
	return o.startPort
}

func (o *options) getEndPort() int {
	return o.endPort
}

func (o *options) isShowClosedPorts() bool {
	return o.showClosedPorts
}

func (o *options) isHideClosedPorts() bool {
	return !o.isShowClosedPorts()
}
