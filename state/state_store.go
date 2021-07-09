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

func (s *StateStore) InsertServiceInstance(id string, instance *models.ServiceInstanceResource) error {
	return s.insertServiceInstance(&serviceInstanceShim{
		id:       id,
		instance: instance,
	})
}

func (s *StateStore) insertServiceInstance(instance *serviceInstanceShim) error {
	txn := s.db.Txn(true)
	defer txn.Abort()

	existing, err := txn.First("service_instance", "id", instance.id)
	if err != nil {
		return err
	}

	if existing != nil {
		return fmt.Errorf("Service instance with ID %s already exists", instance.id)
	}

	if err := txn.Insert("service_instance", instance); err != nil {
		return err
	}

	txn.Commit()
	return nil
}

func (s *StateStore) GetServiceInstance(id string) (*models.ServiceInstanceResource, error) {
	shim, err := s.getServiceInstance(id)
	if err != nil {
		return nil, err
	}

	if shim != nil {
		return shim.instance, nil
	}

	return nil, nil
}

func (s *StateStore) getServiceInstance(id string) (*serviceInstanceShim, error) {
	txn := s.db.Txn(true)
	defer txn.Abort()

	obj, err := txn.First("service_instance", "id", id)
	if err != nil {
		return nil, err
	}

	if obj != nil {
		return obj.(*serviceInstanceShim), nil
	}

	return nil, nil
}

func (s *StateStore) InsertServiceBinding(id string, binding *models.ServiceBindingResource) error {
	return s.insertServiceBinding(&serviceBindingShim{
		id:      id,
		binding: binding,
	})
}

func (s *StateStore) insertServiceBinding(binding *serviceBindingShim) error {
	txn := s.db.Txn(true)
	defer txn.Abort()

	existing, err := txn.First("service_binding", "id", binding.id)
	if err != nil {
		return err
	}

	if existing != nil {
		return fmt.Errorf("Service binding with ID %s already exists", binding.id)
	}

	if err := txn.Insert("service_binding", binding); err != nil {
		return err
	}

	txn.Commit()
	return nil
}

func (s *StateStore) GetServiceBinding(id string) (*models.ServiceBindingResource, error) {
	shim, err := s.getServiceBinding(id)
	if err != nil {
		return nil, err
	}

	if shim != nil {
		return shim.binding, nil
	}

	return nil, nil
}

func (s *StateStore) getServiceBinding(id string) (*serviceBindingShim, error) {
	txn := s.db.Txn(true)
	defer txn.Abort()

	obj, err := txn.First("service_binding", "id", id)
	if err != nil {
		return nil, err
	}

	if obj != nil {
		return obj.(*serviceBindingShim), nil
	}

	return nil, nil
}
