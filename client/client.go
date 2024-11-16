package main

import (
	"log"
	"net"
	"time"
)

func main() {
	ipAddress := "192.168.80.118:80"
	conn, err := net.Dial("tcp", ipAddress)
	if err != nil {
		log.Fatalln(err)
	}
	c := conn.(*net.TCPConn)
	// if err := c.SetLinger(1); err != nil {
	// 	log.Fatalln(err)
	// }

	if _, err := c.Write([]byte("hello")); err != nil {
		log.Fatalln(err)
	}

	time.Sleep(2 * time.Second)

	if err := c.Close(); err != nil {
		log.Fatalln(err)
	}
}