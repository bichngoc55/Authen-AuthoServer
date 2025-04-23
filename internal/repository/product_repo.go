package repository

import (
	cache "AUTHEN-AUTHOSERVER/internal/cache"
	db "AUTHEN-AUTHOSERVER/internal/db"
	models "AUTHEN-AUTHOSERVER/internal/model"
	"encoding/json"
	"fmt"
)
func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	val, err := cache.Client.Get(cache.Ctx, "products").Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &products); err == nil {
			return products, nil
		}
	}
	query := `SELECT id, name, description, price, created_by FROM products`
	err = db.DB.Select(&products, query)
	if err != nil {
		return nil, err
	}
	jsonData, _ := json.Marshal(products)
	cache.Client.Set(cache.Ctx, "products", jsonData, 0)
	return products, nil
}
func GetProductByID(id int )( *models.Product ,error){
	var product models.Product
	key :=fmt.Sprintf("product:%d",id)
	val,err := cache.Client.Get(cache.Ctx,key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &product); err == nil {
			return &product, nil
		}
	}
	query := `SELECT id, name, description, price, created_by FROM products WHERE id = $1`
	err = db.DB.Get(&product, query, id)
	if err != nil {
		return nil, err
	}
	jsonData, _ := json.Marshal(product)
	cache.Client.Set(cache.Ctx, key, jsonData, 0)
	return &product, nil
}
func CreateProduct(product *models.Product) error {
	query := `INSERT INTO products (name, description, price, created_by) VALUES ($1, $2, $3, $4)`
	// fmt.Print("Product created by user ID: %d", query)
	rows , err := db.DB.NamedQuery(query, product)
	if err != nil {
		return err
	}
	
	defer rows.Close()
	if rows.Next(){
		err = rows.Scan(&product.ID)
		if err != nil{
			return err
		}
	}
	cache.Client.Del(cache.Ctx, "products")
	return nil
}
func UpdateProduct(product *models.Product) error {
	query := `UPDATE products SET name = :name, description = :description, price = :price WHERE id = :id`
	_, err := db.DB.NamedExec(query, product)
	if err != nil {
		return err
	}
	cache.Client.Del(cache.Ctx, "products", fmt.Sprintf("product:%d", product.ID))
	return nil
}
	func DeleteProduct(id int) error {
		query := `DELETE FROM products WHERE id = $1`
		_, err := db.DB.Exec(query, id)
		if err != nil {
			return err
		}
		cache.Client.Del(cache.Ctx, "products", fmt.Sprintf("product:%d", id))
		return nil
	}