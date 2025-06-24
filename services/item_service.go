package services

import (
	"Go_gin/dto"
	"Go_gin/models"
	"Go_gin/repositories"
	"fmt"
)

type IItemService interface {
	Create(createItemInput dto.CreateItemInput) (*models.Item, error)
	FindAll() (*[]models.Item, error)
	FindByID(id uint) (*models.Item, error)
	BulkCreate(count int) error
	DeleteAll() error
}

type ItemService struct {
	repository repositories.IItemRepository
}

func NewItemService(repository repositories.IItemRepository) IItemService {
	return &ItemService{repository: repository}
}

func (s *ItemService) Create(createItemInput dto.CreateItemInput) (*models.Item, error) {
	newItem := models.Item{
		Name:        createItemInput.Name,
		Price:       createItemInput.Price,
		Description: createItemInput.Description,
		SoldOut:     false,
	}
	return s.repository.Create(newItem)
}

func (s *ItemService) FindAll() (*[]models.Item, error) {
	return s.repository.FindAll()
}

func (s *ItemService) FindByID(id uint) (*models.Item, error) {
	return s.repository.FindByID(id)
}

func (s *ItemService) BulkCreate(count int) error {
	if err := s.DeleteAll(); err != nil {
		return fmt.Errorf("failed to delete existing data: %w", err)
	}

	batchSize := 8000
	batches := count / batchSize
	for b := 0; b < batches; b++ {
		items := make([]models.Item, 0, batchSize)
		for i := 0; i < batchSize; i++ {
			index := b*batchSize + i + 1
			item := models.Item{
				Name:        fmt.Sprintf("Item %d", index),
				Price:       uint(index) * 10,
				Description: fmt.Sprintf("Description for item %d", index),
				SoldOut:     false,
			}
			items = append(items, item)
		}
		if err := s.repository.BulkCreate(items); err != nil {
			return err
		}
	}
	return nil
}

func (s *ItemService) DeleteAll() error {
	return s.repository.DeleteAll()
}
