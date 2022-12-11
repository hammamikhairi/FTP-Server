// Implementation of the File Transfer Protocol (FTP) server.
//
// Thanks to Angus Morrison for the solution that got me going https://github.com/AngusGMorrison/the_go_programming_language/tree/master/ch8/ex2/ftpserver
package main

import (
	"FTPServer/ftp"
	"fmt"
	"log"
	"net"
	"path/filepath"
)

var port int = 6969
var rootDir string = "container"

func init() {
	log.Println("Firing up FTP server...")
}

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
