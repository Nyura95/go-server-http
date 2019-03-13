package v1

import "go-server-http/models"

type payload struct {
	Users []*models.User
}
