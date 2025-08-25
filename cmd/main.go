package cmd

import (
	"go_product_catalog/configs"
	"go_product_catalog/db"
	"go_product_catalog/models"
	"log"
)

func main() {
	config := configs.LoadConfig()

	database := db.NewDb(config)

	err := database.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}
	log.Println("Миграция завершена. Таблица 'products' создана или обновлена.")
}
