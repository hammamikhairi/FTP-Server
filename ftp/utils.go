package ftp

import "log"

func fatal(msg interface{}) {
	log.Println("<!> ", msg)
}
