package storage

type User struct {
	Id   string
	Name string
}

type MapStorage struct {
	IdByUser map[string]*User
}

func (m *MapStorage) GetUserById(str string) (*User, error) {
	return m.IdByUser[str], nil
}

func (m *MapStorage) AddUser(a string, name string) error {

}
