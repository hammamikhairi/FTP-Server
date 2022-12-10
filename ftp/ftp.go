package ftp

import (
	"bufio"
	"log"
	"net"
	"strings"
)

const (
	EOL       = "\n"
	status220 = "[220] Service ready for new user <%s>."
	status501 = "[501] Service not implemented."
)

func NewConn(conn net.Conn, rootDir string) *Conn {
	return &Conn{
		conn:    conn,
		rootDir: rootDir,
		workDir: "/",
	}
}

func Serve(c *Conn) {
	c.respond(status220)

	s := bufio.NewScanner(c.conn)
	for s.Scan() {
		input := strings.Fields(s.Text())
		if len(input) == 0 {
			continue
		}
		command, args := input[0], input[1:]
		log.Printf("<<< %s %v", command, args)

		switch command {
		case "SYST":
			c.respond("args")
		case "CWD":
			c.respond("cwd")
		case "LIST":
			c.respond("list")
		case "PORT":
			c.respond("port")
		case "USER":
			c.respond("user")
		case "QUIT":
			c.respond("quit")
			return
		case "RETR":
			c.respond("retr")
		case "TYPE":
			c.respond("type")
		default:
			c.respond(status501)
		}
	}
	if s.Err() != nil {
		log.Print(s.Err())
	}
}
