package service

import (
	models "AUTHEN-AUTHOSERVER/internal/model"
	repository "AUTHEN-AUTHOSERVER/internal/repository"
)
func GetAllProducts() ([]models.Product, error) {
	return repository.GetAllProducts()
}

func GetProductByID(id int) (*models.Product, error) {
	return repository.GetProductByID(id)
}

func CreateProduct(product *models.Product) error {
	return repository.CreateProduct(product)
}

func UpdateProduct(product *models.Product) error {
	return repository.UpdateProduct(product)
}

func DeleteProduct(id int) error {
	return repository.DeleteProduct(id)
}