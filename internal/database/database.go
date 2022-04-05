package database

import (
	"encoding/json"
	"errors"
	"os"
)

type Client struct {
	path string
}

type databaseSchema struct {
	Users map[string]User `json:"users"`
	Posts map[string]Post `json:"posts"`
}

func NewClient(path string) Client {
	return Client{
		path: path,
	}
}

func (c Client) createDB() error {
	data, err := json.MarshalIndent(databaseSchema{
		Users: make(map[string]User),
		Posts: make(map[string]Post),
	}, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(c.path, data, 0666)

	if err != nil {
		return err
	}

	return nil
}

func (c Client) EnsureDB() error {
	_, err := os.ReadFile(c.path)

	if errors.Is(err, os.ErrNotExist) {
		return c.createDB()
	}

	return err
}

func (c Client) updateDB(db databaseSchema) error {
	data, err := json.MarshalIndent(db, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(c.path, data, 0666)
	if err != nil {
		return err
	}

	return nil
}

func (c Client) readDB() (databaseSchema, error) {
	data, err := os.ReadFile(c.path)

	if err != nil {
		return databaseSchema{}, err
	}

	data = []byte(data)

	dbData := databaseSchema{}

	err = json.Unmarshal(data, &dbData)
	if err != nil {
		return databaseSchema{}, err
	}

	return dbData, nil
}
