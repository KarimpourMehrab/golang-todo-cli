package memoryStore

import "todo/entity"

type Category struct {
	categories []entity.Category
}

func NewCategoryRepo() *Category {
	return &Category{
		categories: make([]entity.Category, 0),
	}
}

func (c *Category) CreateCategory(category entity.Category) (entity.Category, error) {
	category.ID = len(c.categories) + 1
	c.categories = append(c.categories, category)

	return category, nil
}

func (c *Category) ListCategory(userId int) ([]entity.Category, error) {
	var currentUserCategories []entity.Category
	for _, category := range c.categories {
		if category.UserId == userId {
			currentUserCategories = append(currentUserCategories, category)
		}
	}
	return currentUserCategories, nil
}
