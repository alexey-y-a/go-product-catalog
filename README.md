## Product Catalog API (Go + GORM + PostgreSQL + Docker)

Минималистичный бэкенд-сервис для управления каталогом продуктов, написанный на Go.  

## 🚀 Функционал
- Конфигурация через `.env`
- Подключение к PostgreSQL в Docker
- Модель `Product` с миграцией
- Поддержка массива изображений (`text[]` в PostgreSQL)

---

## 🛠 Технологии
- **Go 1.24**
- **GORM** — ORM
- **PostgreSQL 16.4** — в Docker
- **godotenv** — загрузка переменных окружения
- **Docker Compose** — оркестрация БД
