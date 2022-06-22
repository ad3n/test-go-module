package main

type Module struct {
	Value string
}

func (m *Module) SetValue(value string) {
	m.Value = value
}

func (m *Module) GetValue() string {
	return m.Value
}
