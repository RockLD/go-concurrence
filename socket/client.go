package socket

import (
	"net"
	"fmt"
	"time"
)

func Client() {
	conn, err := net.DialTimeout("tcp","127.0.0.1:8501",5 * time.Second)
	if err != nil {
		fmt.Printf("Dial err is :%s\n",err)
		return
	}

	defer conn.Close()
	fmt.Printf("Connected to server.remote address:%s,local address:%s",conn.RemoteAddr(),conn.LocalAddr())
	time.Sleep(200 * time.Millisecond)

}
