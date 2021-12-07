package initialize

import (
	"log"
	"net"
)

func InitPort() (port int, err error) {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}
	li, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer func(li *net.TCPListener) {
		err := li.Close()
		if err != nil {
			log.Panic(err)
		}
	}(li)
	return li.Addr().(*net.TCPAddr).Port, nil
}
