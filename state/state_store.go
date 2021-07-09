package state

import (
	"fmt"

	memdb "github.com/hashicorp/go-memdb"
	"github.com/walkergriggs/hellosb/models"
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

func (s *StateStore) InsertServiceInstance(instance *models.ServiceInstanceResource) error {
	txn := s.db.Txn(true)
	defer txn.Abort()

	existing, err := txn.First("service_instance", "id", instance.ServiceID)
	if err != nil {
		return err
	}

	if existing != nil {
		return fmt.Errorf("Service instance with ID %s already exists", instance.ServiceID)
	}

	if err := txn.Insert("service_instance", instance); err != nil {
		return err
	}

	txn.Commit()

	return nil
}

func (s *StateStore) GetServiceInstance(id string) (*models.ServiceInstanceResource, error) {
	txn := s.db.Txn(true)
	defer txn.Abort()

	obj, err := txn.First("service_instance", "id", id)
	if err != nil {
		return nil, err
	}

	if obj != nil {
		return obj.(*models.ServiceInstanceResource), nil
	}

	return nil, nil
}
