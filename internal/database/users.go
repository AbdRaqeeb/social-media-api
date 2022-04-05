package database

import (
	"errors"
	"time"
)

// User
type User struct {
	CreatedAt time.Time `json:"createdAt"`
	Email string `json:"email"`
	Password string `json:"password"`
	Name string `json:"name"`
	Age int `json:"age"`
}


func (c Client) CreateUser(email, password, name string, age int) (User, error) {
	data, err := c.readDB()
	if err != nil {
		return User{}, err
	}

	_, ok := data.Users[email]
	if ok {
		return User{}, errors.New("user already exists")
	}

	user := User{
		CreatedAt: time.Now().UTC(),
		Email: email,
		Name: name,
		Password: password,
		Age: age,
	}

	data.Users[email] = user

	err = c.updateDB(data)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (c Client) UpdateUser(email, password, name string, age int) (User, error) {
	data, err := c.readDB()
	if err != nil {
		return User{}, err
	}

	_, ok := data.Users[email]
	if !ok {
		return User{}, errors.New("user doesn't exist")
	}

	user := User{
		CreatedAt: time.Now().UTC(),
		Email: email,
		Name: name,
		Password: password,
		Age: age,
	}

	data.Users[email] = user

	err = c.updateDB(data)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (c Client) GetUser(email string) (User, error) {
	data, err := c.readDB()
	if err != nil {
		return User{}, err
	}

	_, ok := data.Users[email]
	if !ok {
		return User{}, errors.New("user doesn't exist")
	}

	return data.Users[email], nil
}

func (c Client) DeleteUser(email string) error {
	data, err := c.readDB()
	if err != nil {
		return err
	}

	_, ok := data.Users[email]
	if !ok {
		return errors.New("user not found")
	}

	delete(data.Users, email)

	err = c.updateDB(data)
	if err != nil {
		return err
	}

	return nil
}
