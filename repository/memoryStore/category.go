package memoryStore

import "todo/entity"

type Category struct {
	categories []entity.Category
}

func (c *Category) CreateCategory(category entity.Category) (entity.Category, error) {
	category.ID = len(c.categories) + 1
	c.categories = append(c.categories, category)
	return category, nil
}

func (c *Category) UserHasCategory(authenticatedUserId, categoryId int) bool {
	for _, category := range c.categories {
		if category.UserId == authenticatedUserId && category.ID == categoryId {
			return true
		}
	}
	return false
}
