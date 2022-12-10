// Implement a concurrent File Transfer Protocol (FTP) server. The server should interpret commands
// from each client such as cd to change directory, ls to list a directory, get to send the contents
// of a file, and close to close the connection.
//
// Thanks to Kdama for the solution that got me going https://github.com/kdama/gopl/blob/master/ch08/ex02/main.go
package main

import (
	"FTPServer/ftp"
	"fmt"
	"net"
	"path/filepath"
)

var port int = 6969
var rootDir string = "container"

func main() {
	server := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", server)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	absPath, err := filepath.Abs(rootDir)

	if err != nil {
		panic(err)
	}

	ftp.Serve(ftp.NewConn(c, absPath))
}
