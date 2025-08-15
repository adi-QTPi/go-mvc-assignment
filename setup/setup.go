package setup

import (
	"fmt"

	"github.com/adi-QTPi/go-mvc-assignment/config"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func MakeAdminUser() error {
	config.LoadAdminDetailsEnv()
	config.LoadJwtEnv()

	var admin models.User
	admin.UserName = config.ADMIN_USERNAME
	admin.Name = config.ADMIN_NAME
	adminPassword := config.ADMIN_PASSWORD
	admin.Role = "admin"

	alreadyExists, err, _ := models.GetUserByUsername(admin.UserName)
	if alreadyExists {
		return nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminPassword), 10)
	if err != nil {
		return fmt.Errorf("Failed to hash password, error from bcrypt in the middleware hash password : %v", err)

	}

	err = models.AddNewUser(admin, string(hashedPassword))
	if err != nil {
		return fmt.Errorf("error creating the admin user : %v", err)
	}

	return nil
}
