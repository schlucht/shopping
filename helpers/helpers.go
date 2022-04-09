package helpers

const NAME = "Schmid"

type Addr struct {
	Street string
	Place string
}

func (m *Addr) ToString() string {
	return "Street: " + m.Street + "Place:" + m.Place
}