package repository

import (
	"fmt"
	"g-mysql/entity"
	"gorm.io/gorm"
)

type IProductRepository interface {
	SaveProduct(product entity.Product) (entity.Product, error)
	FindById(Id int) (entity.Product, error)
	UpdateProduct(product entity.Product) (entity.Product, error)
	DeleteProduct(Id int) (entity.Product, error)
	FindAllProduct(product []entity.Product) ([]entity.Product, error)
}

type Repository struct {
	DB *gorm.DB
}

func Newrepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) SaveProduct(product entity.Product) (entity.Product, error) {
	err := r.DB.Create(&product).Error
	if err != nil {
		fmt.Println(err)
	}
	return product, err
}

func (r *Repository) FindAllProduct(product []entity.Product) ([]entity.Product, error) {
	err := r.DB.Find(&product).Error
	if err != nil {
		panic(err.Error())
	}
	return product, err
}

func (r *Repository) FindById(Id int) (entity.Product, error) {
	var product entity.Product
	err := r.DB.Find(&product, Id).Error
	fmt.Println(product)
	return product, err
}

func (r *Repository) UpdateProduct(product entity.Product) (entity.Product, error) {
	r.DB.First(product)
	err := r.DB.Save(&product).Error
	return product, err
}

func (r *Repository) DeleteProduct(Id int) (entity.Product, error) {
	var product entity.Product
	err := r.DB.Delete(product, Id).Error
	return product, err
}
