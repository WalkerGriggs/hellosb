package state

import (
	memdb "github.com/hashicorp/go-memdb"
)

type StateStore struct {
	db *memdb.MemDB
}

func NewStateStore() (*StateStore, error) {
	s := &StateStore{}

	var err error
	if s.db, err = s.setupDB(); err != nil {
		return nil, err
	}

	return s, nil
}
