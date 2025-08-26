package main

import (
	"go_product_catalog/configs"
	"go_product_catalog/db"
	"go_product_catalog/models"
	"go_product_catalog/repository"
	"log"
)

func main() {
	config := configs.LoadConfig()

	database := db.NewDb(config)

	productRepo := repository.NewProductRepository(database.DB)

	log.Println("1. Создаём новый продукт...")
	product := &models.Product{
		Name:        "iPhone 15",
		Description: "Latest Apple smartphone with advanced camera",
		Images:      []string{"https://example.com/iphone15-1.jpg", "https://example.com/iphone15-2.jpg"},
	}

	if err := productRepo.Create(product); err != nil {
		log.Fatalf("Ошибка создания продукта: %v", err)
	}
	log.Printf("Продукт создан с ID: %d\n", product.ID)

	log.Println("2. Получаем продукт по ID...")
	found, err := productRepo.GetByID(product.ID)
	if err != nil {
		log.Fatalf("Ошибка получения: %v", err)
	}
	if found == nil {
		log.Fatal("Продукт не найден")
	}
	log.Printf("Найден: %s — %s\n", found.Name, found.Description)

	log.Println("3. Обновляем продукт...")
	found.Name = "iPhone 15 Pro"
	found.Description = "Even better camera and titanium frame"
	found.Images = []string{"https://example.com/iphone15-pro-1.jpg"}

	if err := productRepo.Update(found.ID, found); err != nil {
		log.Fatalf("Ошибка обновления: %v", err)
	}
	log.Println("Продукт обновлён")

	log.Println("4. Удаляем продукт...")
	if err := productRepo.Delete(product.ID); err != nil {
		log.Fatalf("Ошибка удаления: %v", err)
	}
	log.Println("Продукт удалён")

	afterDelete, _ := productRepo.GetByID(product.ID)
	if afterDelete == nil {
		log.Println("Подтверждено: продукт больше не существует")
	}
}
