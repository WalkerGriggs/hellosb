package state

import (
	"fmt"

	"github.com/walkergriggs/hellosb/models"
)

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

func (s *StateStore) DeleteServiceInstance(id string) error {
	txn := s.db.Txn(true)
	defer txn.Abort()

	existing, err := txn.First("service_instance", "id", id)
	if err != nil {
		return err
	}

	if existing == nil {
		return fmt.Errorf("Service instance %s not found", id)
	}

	if err := txn.Delete("service_instance", existing); err != nil {
		return fmt.Errorf("Failed to delete service instance: %v", err)
	}

	txn.Commit()
	return nil
}

func (s *StateStore) UpdateServiceInstance(id string, req *models.ServiceInstanceUpdateRequest) error {
	txn := s.db.Txn(true)
	defer txn.Abort()

	obj, err := txn.First("service_instance", "id", id)
	if err != nil {
		return err
	}

	if obj == nil {
		return fmt.Errorf("Service instance %s not found", id)
	}

	existing := obj.(*serviceInstanceShim).instance

	shim := &serviceInstanceShim{
		id: id,
		instance: &models.ServiceInstanceResource{
			DashboardURL:    existing.DashboardURL,
			MaintenanceInfo: existing.MaintenanceInfo,
			Metadata:        existing.Metadata,
			Parameters:      existing.Parameters,
			PlanID:          existing.PlanID,
			ServiceID:       existing.ServiceID,
		},
	}

	if req.MaintenanceInfo != nil {
		shim.instance.MaintenanceInfo = req.MaintenanceInfo
	}

	if req.PlanID != "" {
		shim.instance.PlanID = req.PlanID
	}

	if err := txn.Insert("service_instance", shim); err != nil {
		return err
	}

	txn.Commit()
	return nil
}
