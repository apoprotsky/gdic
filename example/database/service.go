package database

// Service structure
type Service struct {
	Name string
}

// InitService initialize service
func (svc *Service) InitService() {
	svc.Name = "database"
}

// GetEmailByName return email by account name
func (svc *Service) GetEmailByName(name string) string {
	return name + "@example.com"
}
