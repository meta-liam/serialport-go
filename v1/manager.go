package v1

import "errors"

type Manage struct {
	ports map[string]*Port //:= mack(map[string]*port)
}

func (m *Manage) GetDetailedPortsList() ([]*PortDetails, error) {
	return GetDetailedPortsList()
}

func (m *Manage) GetPorts() ([]string, error) {
	return GetPorts()
}

func (m *Manage) CreatePort(portName string, mode *Mode) {
	port := new(Port)
	port.Open(portName, mode)
	m.ports[portName] = port
}

func (m *Manage) Write(portName string, b []byte) (n int, err error) {
	if m.ports[portName] == nil || !m.ports[portName].IsOpen() {
		return 0, errors.New("no open")
	}
	return m.ports[portName].Write(b)
}

func (m *Manage) Read(portName string, b []byte) (n int, err error) {
	if m.ports[portName] == nil || !m.ports[portName].IsOpen() {
		return 0, errors.New("no open")
	}
	return m.ports[portName].Read(b)
}

func (m *Manage) Close(portName string) (err error) {
	if m.ports[portName] == nil || !m.ports[portName].IsOpen() {
		return errors.New("no open")
	}
	return m.ports[portName].Close()
}

func (m *Manage) CloseAll() {
	for k := range m.ports {
		if m.ports[k] != nil && m.ports[k].IsOpen() {
			m.ports[k].Close()
		}
	}
}
