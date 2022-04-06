package user

type User struct {
	FirstName string
	LastName  string
	Age       string
}

func (m *User) ToString() string {
	return m.FirstName + " " + m.LastName
}