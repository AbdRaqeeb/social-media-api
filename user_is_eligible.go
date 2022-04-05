package main

import (
	"errors"
	"fmt"
)

func userIsEligible(email, password string, age int) error {
	if email == "" {
		return errors.New("email can't be empty")
	}

	if password == "" {
		return errors.New("password can't be empty")
	}

	const minAge = 18

	if age < minAge {
		return fmt.Errorf("age must be at least %d years old", minAge)
	}

	return nil
}
