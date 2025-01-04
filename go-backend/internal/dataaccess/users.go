package users

import (
	"github.com/zzibo/cvwo-web-forum/internal/database"
	"github.com/zzibo/cvwo-web-forum/internal/models"
)

func List(db *database.Database) ([]models.User, error) {
	users := []models.User{
		{
			ID:   1,
			Name: "CVWO",
		},
	}
	return users, nil
}
