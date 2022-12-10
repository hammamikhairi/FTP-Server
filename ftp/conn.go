package ftp

import (
	"fmt"
	"log"
	"net"
)

type Conn struct {
	conn    net.Conn
	rootDir string
	workDir string
}

func (c *Conn) respond(s string) {
	log.Print(">> ", s)
	_, err := fmt.Fprint(c.conn, s, EOL)
	if err != nil {
		log.Print(err)
	}
}
