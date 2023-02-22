package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})
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
	var p Product
	db.First(&p, 7)
	p.Name = "Novo nome"
	db.Save(&p)

	var p2 Product
	db.First(&p2, 7)
	db.Delete(&p2)

}
