package database

import (
	"testing"

	"github.com/Marcelospegiorin/crud-golang-example/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateCategory(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Category{})
	category, _ := entity.NewCategory("Celulares")
	assert.NoError(t, err)
	categoryDB := NewCategory(db)

	err = categoryDB.Create(category)
	assert.Nil(t, err)
	assert.NotEmpty(t, category.ID)
	assert.NotEmpty(t, category.Name)
	assert.NotEmpty(t, category.CreatedAt)

	var categoryFound *entity.Category
	err = db.First(&categoryFound, "id = ?", category.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, category.ID, categoryFound.ID)
	assert.Equal(t, category.Name, categoryFound.Name)
}

func TestFindCategory(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Category{})
	category, err := entity.NewCategory("Celulares")
	assert.NoError(t, err)
	db.Create(category)

	categoryDB := NewCategory(db)
	categoryFind, err := categoryDB.FindByID(category.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, category.ID, categoryFind.ID)
	assert.Equal(t, "Celulares", categoryFind.Name)
}
