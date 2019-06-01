package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"microservices/pkg"
	"microservices/services/product/entity"
)

func migrate(db *pkg.MysqlRepository) {
	var product = &entity.Product{}
	if !db.HasTable(product) {
		_, err := db.AutoMigrate(product)
		if err != nil {
			panic("error migrating product table")
		}
		log.Printf("Migrating database...")
		for _, value := range readProductsFromJSONFile() {
			db.Save(&value)
		}
	}
}

func readProductsFromJSONFile() []entity.Product {
	jsonFile, err := ioutil.ReadFile("static/products.json")
	if err != nil {
		panic("Error reading products.json: " + err.Error())
	}

	var result []entity.Product
	err2 := json.Unmarshal(jsonFile, &result)
	if err2 != nil {
		panic("Error parsing products.json")
	}
	return result
}
