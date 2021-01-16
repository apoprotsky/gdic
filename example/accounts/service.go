package accounts

import (
	"github.com/apoprotsky/gservices"
	"github.com/apoprotsky/gservices/example/database"
)

// Account structure
type Account struct {
	Name  string
	Email string
}

// Service structure
// Public fields of service structure can be initialized without Init method
type Service struct {
	databaseService *database.Service
}

// Init initializes service
// Method called by gservices when instance of service is created
func (svc *Service) Init(databaseService *database.Service) {
	svc.databaseService = databaseService
}

// GetByName returns Account structure by name
func (svc *Service) GetByName(name string) *Account {
	return &Account{
		Name:  name,
		Email: svc.databaseService.GetEmailByName(name),
	}
}

// GetService returns Service instance
func GetService() *Service {
	return gservices.Get((*Service)(nil)).(*Service)
}
