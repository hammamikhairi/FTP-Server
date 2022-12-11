package ftp

import "fmt"

type dataPort struct {
	h1, h2, h3, h4 int // host
	p1, p2         int // port
}

func (port *dataPort) toIPAddress() string {
	dataPort := port.p1<<8 + port.p2
	return fmt.Sprintf(
		"%d.%d.%d.%d:%d",
		port.h1, port.h2, port.h3, port.h4, dataPort,
	)
}
