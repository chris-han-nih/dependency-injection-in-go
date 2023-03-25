package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Saver interface {
	Save(data []byte) (int, error)
}

func SavePerson(person *Person, saver Saver) (int, error) {
	err := person.validate()
	if err != nil {
		return 0, err
	}
	bytes, err := person.encode()
	if err != nil {
		return 0, err
	}
	return saver.Save(bytes)
}

type Person struct {
	Name  string
	Phone string
}

func (p *Person) validate() error {
	if p.Name == "" {
		return errors.New("name is empty")
	}
	if p.Phone == "" {
		return errors.New("phone is empty")
	}
	return nil
}

func (p *Person) encode() ([]byte, error) {
	return json.Marshal(p)
}

func LoadPerson(ID int, decodePerson func(data []byte) *Person) (*Person, error) {
	if ID <= 0 {
		return nil, fmt.Errorf("invalid ID '%d' supplied", ID)
	}
	bytes, err := loadPerson(ID)
	if err != nil {
		return nil, err
	}
	return decodePerson(bytes), nil
}

func loadPerson(ID int) ([]byte, error) {
	// TODO: implement
	return nil, errors.New("not implemented")
}
