package main

import (
	"fmt"

	v1 "github.com/meta-liam/serialport-go/v1"
)

func main() {
	path := "/dev/ttyS0"

	m := new(v1.Manage)
	v, _ := m.GetPorts()
	path = v[0]
	fmt.Printf("%+v\n", v)
	fmt.Printf("%+v\n", path)
	mode := &v1.Mode{
		BaudRate: 115200,
	}
	m.CreatePort(path, mode)

	buf := make([]byte, 256)
	for {
		n, _ := m.Read(path, buf)
		fmt.Printf("read(%d): %s\n", n, string(buf[:n]))
		m.Write(path, buf[:n])
	}
	// m.Close(path) //关闭
}
