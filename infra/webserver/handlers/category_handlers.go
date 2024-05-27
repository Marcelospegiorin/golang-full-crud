package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Marcelospegiorin/crud-golang-example/infra/database"
	"github.com/Marcelospegiorin/crud-golang-example/infra/dto"
	"github.com/Marcelospegiorin/crud-golang-example/internal/entity"
)

type CategoryHandler struct {
	CategoryDB database.CategoryInterface
}

func NewCategoryHandler(db database.CategoryInterface) *CategoryHandler {
	return &CategoryHandler{
		CategoryDB: db,
	}
}

// Create Category
func (c *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category dto.CreateCategory
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := entity.NewCategory(category.Name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = c.CategoryDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Get Categories

func (c *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {

	sort := r.URL.Query().Get("sort")

	categories, err := c.CategoryDB.FindAll(sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
}
