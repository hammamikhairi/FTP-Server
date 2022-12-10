package ftp

import (
	"fmt"
	"log"
	"net"
	"path/filepath"
	"strings"
)

type dataPort struct {
	h1, h2, h3, h4 int // host
	p1, p2         int // port
}

type Conn struct {
	conn     net.Conn
	dataPort *dataPort
	rootDir  string
	workDir  string
	user     string
}

func (c *Conn) respond(s string) {
	log.Print(">> ", s)
	_, err := fmt.Fprint(c.conn, s, EOL)
	if err != nil {
		log.Print(err)
	}
}

func (c *Conn) setUser(args params) {
	c.user = strings.Join(args, " ")

	c.respond(fmt.Sprintf(status202, c.user))
}

func (c *Conn) syst() {
	c.respond(status215)
}

func (c *Conn) pwd() {
	c.respond(c.workDir)
}

func (c *Conn) list(args params) {
	var directory string = filepath.Join(c.rootDir, c.workDir)

	if len(args) > 0 {
		directory = filepath.Join(directory, args[0])
	}

	c.respond(directory)

}

func (c *Conn) port(args params) {
	if len(args) != 1 {
		c.respond(status400)
		return
	}
	dataPort, err := func(hostPort string) (*dataPort, error) {
		var dp dataPort
		_, err := fmt.Sscanf(
			hostPort,
			"%d,%d,%d,%d,%d,%d",
			&dp.h1, &dp.h2, &dp.h3, &dp.h4, &dp.p1, &dp.p2,
		)
		if err != nil {
			return nil, err
		}
		return &dp, nil
	}(args[0])

	if err != nil {
		log.Print(err)
		c.respond(status500)
		return
	}
	c.dataPort = dataPort
	c.respond(status200)
}
