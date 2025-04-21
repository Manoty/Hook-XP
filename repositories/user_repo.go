package repositories


import (
	"StreefySherehes/databases"
	"StreefySherehes/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() UserRepository {
	return UserRepository{
		DB: databases.DB,
	}

}

func (repo *UserRepository) CreateUser(user *models.User) error {
	
		return databases.DB.Create(user).Error
	}

func (repo *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
if err := databases.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func(repo *UserRepository) UpdatePassword(user *models.User) error {
	return repo.DB.Model(user).Update("password", user.Password).Error
}