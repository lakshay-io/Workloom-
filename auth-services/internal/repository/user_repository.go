package repositories

import (
	"github.com/workloom/shared/models"
	"gorm.io/gorm"
)

type UserRepository interface {
    CreateUser(user *models.User) error
    FindUserByEmail(email string) (*models.User, error)
	SaveUser(user *models.User) error
}

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db}
}

func (r *userRepository) CreateUser(user *models.User) error {
    return r.db.Create(user).Error
}

func (r *userRepository) FindUserByEmail(email string) (*models.User, error) {
    var user models.User
    err := r.db.Where("email = ?", email).First(&user).Error
    return &user, err
}

func (r *userRepository) SaveUser(user *models.User) error {
	return r.db.Save(user).Error
}
