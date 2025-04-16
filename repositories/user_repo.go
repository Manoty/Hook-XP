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

func (repo *UserRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
if err := repo.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}