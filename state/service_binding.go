package state

import (
	"fmt"

	"github.com/walkergriggs/hellosb/models"
)

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

func (s *StateStore) DeleteServiceBinding(id string) error {
	txn := s.db.Txn(true)
	defer txn.Abort()

	existing, err := txn.First("service_binding", "id", id)
	if err != nil {
		return err
	}

	if existing == nil {
		return fmt.Errorf("Service binding %s not found", id)
	}

	if err := txn.Delete("service_binding", existing); err != nil {
		return fmt.Errorf("Failed to delete service binding: %v", err)
	}

	txn.Commit()
	return nil
}
