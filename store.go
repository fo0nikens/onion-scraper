package main

import (
	"encoding/csv"
	"os"
)

// Store is our storage
type Store struct {
}

// NewStore creates new storage layer
func NewStore() *Store {
	return &Store{}
}

// CSV returns csv data storage
func (s *Store) CSV(filename string) (*csv.Writer, error) {
	file, err := os.Open(filename)
	if err != nil {
		file, err = os.Create(filename)
		if err != nil {
			return nil, err
		}
	}

	return csv.NewWriter(file), nil
}
