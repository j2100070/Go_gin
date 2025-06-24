package repositories

import (
	"Go_gin/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(newUser models.User) (*models.User, error)
	FindAll() (*[]models.User, error)
	FindByID(id uint) (*models.User, error)
	BulkCreate(users []models.User) error
	DeleteAll() error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(newUser models.User) (*models.User, error) {
	result := r.db.Create(&newUser)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newUser, nil
}

// FindAll retrieves all items from the database.

func (r *UserRepository) FindAll() (*[]models.User, error) {
	var users []models.User
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}

// id を指定してのアイテムを取得する
func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var item models.User
	result := r.db.First(&item, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &item, nil
}

func (r *UserRepository) BulkCreate(users []models.User) error {
	result := r.db.Create(&users)
	return result.Error
}

func (r *UserRepository) DeleteAll() error {
	result := r.db.Exec("DELETE FROM users")
	return result.Error
}
