package state

import "github.com/walkergriggs/hellosb/models"

type serviceInstanceShim struct {
	id       string
	instance *models.ServiceInstanceResource
}
