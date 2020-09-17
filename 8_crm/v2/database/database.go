package database

type Database struct {}

func (Database) GetUserById(userId int) (map[string]interface{}, error) {
	return nil, nil
}

func (Database) GetCompany() (map[string]interface{}, error) {
	return nil, nil
}

func (Database) SaveCompany(newNumber int) error {
	return nil
}

func (Database) SaveUser(user interface{}) error {
	return nil
}