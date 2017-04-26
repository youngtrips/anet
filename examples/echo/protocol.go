package echo

import (
	"bufio"
	"bytes"
	"encoding/binary"
	//	"fmt"
	"io"
	"net"
)

type EchoProtocol struct {
	header []byte
	writer *bufio.Writer
}

func NewEchoProtocol() *EchoProtocol {
	return &EchoProtocol{
		header: make([]byte, 2),
		writer: nil,
	}
}

func (p *EchoProtocol) Read(conn *net.TCPConn) (interface{}, error) {
	if _, err := io.ReadFull(conn, p.header); err != nil {
		return nil, err
	}
	size := binary.BigEndian.Uint16(p.header)
	data := make([]byte, size)
	if _, err := io.ReadFull(conn, data); err != nil {
		return nil, err
	}
	return string(data), nil
}

func encode(data []byte) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, uint16(len(data))); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func rawSend(w *bufio.Writer, data []byte) error {
	/*
		for _, b := range data {
			fmt.Printf("%02x ", b)
		}
		fmt.Printf("\n")
	*/

	if _, err := w.Write(data); err != nil {
		return err
	}
	if err := w.Flush(); err != nil {
		return err
	}
	return nil
}

func (p *EchoProtocol) Write(conn *net.TCPConn, data interface{}) error {
	if p.writer == nil {
		p.writer = bufio.NewWriter(conn)
	}

	buf, err := encode([]byte(data.(string)))
	if err != nil {
		return err
	}
	return rawSend(p.writer, buf)
}
