package anet

import (
	"net"
)

type Protocol interface {
	Read(conn *net.TCPConn) (interface{}, error)
	Write(conn *net.TCPConn, data interface{}) error
}
