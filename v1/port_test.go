package v1_test

import (
	"fmt"
	"testing"

	v1 "github.com/meta-liam/serialport-go/v1"
)

func Test_GetPorts(t *testing.T) {
	v, err := v1.GetPorts()
	// log.Println("v:", v)
	// log.Println(err)
	if err != nil || v == nil {
		t.Errorf("GetPorts is nil")
	}
}

func Test_GetDetailedPortsList(t *testing.T) {
	v, err := v1.GetDetailedPortsList()
	// fmt.Printf("%+v", v[0])
	// log.Println(err)
	if err != nil || v == nil {
		t.Errorf("GetDetailedPortsList is nil")
	}
}

func TestNewPort(t *testing.T) {
	_port := new(v1.Port)
	mode := &v1.Mode{
		BaudRate: 115200,
	}
	_port.Open("/dev/cu.BLTH ", mode)
	fmt.Printf("p:%+v", _port)
}
