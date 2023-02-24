package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primarykey"`
	Name     string
	Products []Product
}
type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	//SerialNumber SerialNumber
	gorm.Model //cria um uppdate, delete e um cretat
}

/*
	type SerialNumber struct {
		ID        int `gorm:"primaryKey"`
		Number    string
		ProductID int
	}
*/
func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local" // acrecentar quando for usar gorm.model ?charset=utf8mb4&parseTime=True&loc=local
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	//create category
	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	//create product
	db.Create(&Product{
		Name:       "Armario",
		Price:      258.66,
		CategoryID: category.ID,
	})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			println("-", product.Name, category.Name)
		}
	}
	/*
		db.Create(&SerialNumber{
			Number:    "123456",
			ProductID: 1,
		})

		var products []Product
		db.Preload("Category").Preload("SerialNumber").Find(&products)
		for _, product := range products {
			fmt.Println(product.Name, product.Category.Name, product.SerialNumber)
		}



	*/
	/*
			//salva 1 produto
					db.Create(&Product{
						Name:  "Notebook",
						Price: 1000.00,
					})

		//salva varios produtos
				products := []Product{
					{Name: "tablete samsung", Price: 850.00},
					{Name: "Iphone", Price: 5850.00},
					{Name: "Monitor 50 pelegadas", Price: 7850.00},
				}
				db.Create(&products)
	*/

	//selecionar um produto
	/*
		var product Product
		db.First(&product, 1)
		fmt.Println(product)
	*/

	//select all
	/*
		var products []Product
		db.Find(&products)
		for _, product := range products {
			fmt.Println(product)
		}
	*/

	//paginação
	/*
		var products []Product
		db.Limit(2).Offset(2).Find(&products)
		for _, product := range products {
			fmt.Println(product)
		}
	*/

	//where
	/*
		var products []Product
		db.Where("price > ?", 90).Find(&products)
		for _, product := range products {
			fmt.Println(product)
		}
	*/

	//LIKE
	/*
		var products []Product
		db.Where("name LIKE ?", "%book%").Find(&products)
		for _, product := range products {
			fmt.Println(product)
		}
	*/
	//renomear
	/*
			var p Product
			db.First(&p, 7)
			p.Name = "Novo nome"
			db.Save(&p)
		//deletar
			var p2 Product
			db.First(&p2, 7)
			db.Delete(&p2)
	*/

}
