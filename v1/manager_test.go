package v1_test

import (
	"fmt"
	"testing"

	v1 "github.com/meta-liam/serialport-go/v1"
)

func TestManager_GetPorts(t *testing.T) {
	m := new(v1.Manage)
	v, err := m.GetPorts()
	// fmt.Printf("%+v", v)
	// log.Println(err)
	if err != nil || v == nil {
		t.Errorf("GetPorts is nil")
	}
}

func TestManage_GetDetailedPortsList(t *testing.T) {
	v, err := v1.GetDetailedPortsList()
	fmt.Printf("%+v", v[0])
	// log.Println(err)
	if err != nil || v == nil {
		t.Errorf("GetDetailedPortsList is nil")
	}
}
