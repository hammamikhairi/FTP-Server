package ftp

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

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

func (c *Conn) establishConnection() (conn net.Conn, err error) {
	conn, err = net.Dial("tcp", c.dataPort.toIPAddress())
	if err != nil {
		fatal(err)
		return
	}
	return
}

func (c *Conn) list(args params) {
	var directory string = filepath.Join(c.rootDir, c.workDir)
	var relative string = "~" + c.workDir

	if len(args) > 0 {
		relative = relative + args[0]
		directory = filepath.Join(directory, args[0])
	}

	files, err := os.ReadDir(directory)
	if err != nil {
		fatal(err)
		c.respond(status450)
		return
	}

	c.respond(status150)
	socket, err := c.establishConnection()
	if err != nil {
		fatal(err)
		c.respond(status500)
		return
	}

	defer socket.Close()

	fmt.Fprint(socket, "Directory : ", relative, EOL)
	for _, file := range files {
		if file.IsDir() {
			_, err = fmt.Fprint(socket, "\t", "<D>  ", file.Name(), EOL)
		} else {
			_, err = fmt.Fprint(socket, "\t", "<F>  ", file.Name(), EOL)
		}

		if err != nil {
			fatal(err)
			c.respond(status500)
		}
	}
	c.respond(status226)
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
