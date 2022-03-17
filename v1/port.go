package v1

import (
	"fmt"
	"log"

	"go.bug.st/serial"
	"go.bug.st/serial/enumerator"
)

type PortDetails struct {
	Name         string
	IsUSB        bool
	VID          string
	PID          string
	SerialNumber string
}

func GetPorts() ([]string, error) {
	return serial.GetPortsList()

}

func GetDetailedPortsList() ([]*PortDetails, error) {
	list := make([]*PortDetails, 0)
	v, err := enumerator.GetDetailedPortsList()
	log.Println("v:::", v)
	for i := 0; i < len(v); i++ {
		item := new(PortDetails)
		item.Name = v[i].Name
		item.IsUSB = v[i].IsUSB
		item.IsUSB = v[i].IsUSB
		item.PID = v[i].PID
		item.SerialNumber = v[i].SerialNumber
		item.VID = v[i].VID
		list = append(list, item)
	}
	return list, err
}

type Port struct {
	port serial.Port
}

type Mode struct {
	BaudRate int             // The serial port bitrate (aka Baudrate) 115200
	DataBits int             // Size of the character (must be 5, 6, 7 or 8)
	Parity   serial.Parity   // Parity (see Parity type for more info)
	StopBits serial.StopBits // Stop bits (see StopBits type for more info)
}

func (p *Port) Open(portName string, mode *Mode) error {
	_mode := &serial.Mode{
		BaudRate: mode.BaudRate,
	}
	port, err := serial.Open(portName, _mode)
	p.port = port
	fmt.Printf("::%+v", p.port)
	return err
}

func (p *Port) Write(b []byte) (n int, err error) {
	return p.port.Write(b) //[]byte("10,20,30\n\r"))
}

func (p *Port) Read(b []byte) (n int, err error) {
	return p.port.Read(b)
}

func (p *Port) IsOpen() bool {
	return true
}

func (p *Port) Close() error {
	return p.port.Close()
}
