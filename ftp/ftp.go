package ftp

import (
	"bufio"
	"log"
	"net"
	"strings"
)

type params []string

const (
	EOL       = "\n"
	status200 = "200 ~ Command okay."
	status202 = "202 ~ User <%s> logged in successfully."
	status215 = "215 ~ OS : Linux Mint 20.3"
	status220 = "220 ~ Service ready for new user."
	status299 = "299 ~ Closing connection, adios."
	status400 = "400 ~ Bad Request."
	status500 = "500 ~ Internal Server Error."
	status501 = "501 ~ Service not implemented."
	status503 = "503 ~ Service unavailable."
)

func NewConn(conn net.Conn, rootDir string) *Conn {
	return &Conn{
		conn:    conn,
		rootDir: rootDir,
		workDir: "/",
		user:    "",
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
		log.Printf("<<< %s %s %v", c.user, command, args)

		switch command {
		case "SYST":
			c.syst()
		case "CWD":
			c.respond(status501)
		case "PWD":
			c.pwd()
		case "LIST":
			c.list(args)
		case "PORT":
			c.port(args)
		case "USER":
			c.setUser(args)
		case "QUIT":
			c.respond(status299)
			return
		case "RETR":
			c.respond(status501)
		case "TYPE":
			c.respond(status501)
		default:
			c.respond(status501)
		}
	}
	if s.Err() != nil {
		log.Print(s.Err())
	}
}
