package repositories

import (
	"Go_gin/models"

	"gorm.io/gorm"
)

type IItemRepository interface {
	Create(newItem models.Item) (*models.Item, error)
	FindAll() (*[]models.Item, error)
	FindByID(id uint) (*models.Item, error)
	BulkCreate(items []models.Item) error
	DeleteAll() error
}

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) IItemRepository {
	return &ItemRepository{db: db}
}

func (r *ItemRepository) Create(newItem models.Item) (*models.Item, error) {
	result := r.db.Create(&newItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newItem, nil
}

// FindAll retrieves all items from the database.

func (r *ItemRepository) FindAll() (*[]models.Item, error) {
	var items []models.Item
	result := r.db.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return &items, nil
}

// id を指定してのアイテムを取得する
func (r *ItemRepository) FindByID(id uint) (*models.Item, error) {
	var item models.Item
	result := r.db.First(&item, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &item, nil
}

func (r *ItemRepository) BulkCreate(items []models.Item) error {
	result := r.db.Create(&items)
	return result.Error
}

func (r *ItemRepository) DeleteAll() error {
	result := r.db.Exec("DELETE FROM items")
	return result.Error
}
