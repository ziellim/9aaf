package main

import "errors"

type Landmark struct {
	Id		string		`json:"id""`
	Value	string		`json:"value""`
}

func (tm *Landmark) getAll() ([]Landmark, error) {
	return nil, errors.New("Not implemented")
}
