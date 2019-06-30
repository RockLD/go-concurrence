package socket

import (
	"net"
	"fmt"
	"time"
	"io"
	"bytes"
)

const (
	SERVER_NETWORK = "tcp"
	SERVR_ADDRESS = "127.0.0.1:8501"
	DELIMITER = '\t'
)

func Serve() {
	listener, err := net.Listen("tcp","127.0.0.1:8501")
	if err != nil {
		fmt.Printf("serve listen err is :%s\n",err)
		return
	}

	defer listener.Close()
	fmt.Printf("Got listener for the server.(local address:%s)",listener.Addr())
	for  {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("serve accept err is :%s\n",err)
			return
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	for {
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		strReq,err := read(conn)
		if err != nil {
			if err == io.EOF{
				fmt.Printf("the connection is closed by another side")
			}else{
				fmt.Printf("read err:",err)
				return
			}
		}
		fmt.Printf("Received request %s",strReq)
	}
}

func read(conn net.Conn)(string,error){
	readBytes := make([]byte,1)
	var buffer bytes.Buffer
	for {
		_, err := conn.Read(readBytes)
		if err != nil {
			return "",err
		}
		readByte := readBytes[0]
		if readByte == DELIMITER {
			break
		}
		buffer.WriteByte(readByte)
	}
	return buffer.String(),nil
}
