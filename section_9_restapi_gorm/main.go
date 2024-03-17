package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
}

func main() {
	dsn := "root@tcp(127.0.0.1:3306)/eduwork?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal koneksi ke database")
	}

	app := fiber.New()

	// Baca products
	app.Get("/api/products/", func(c *fiber.Ctx) error {
		var products []Product
		if err := db.Find(&products).Error; err != nil {
			return err
		}

		return c.JSON(products)
	})

	// Buat product
	app.Post("/api/products", func(c *fiber.Ctx) error {
		var product Product

		// parse request body
		if err := c.BodyParser(&product); err != nil {
			return err
		}

		product.CreatedAt = time.Now()

		// buat product di database
		db.Create(&product)
		return c.JSON(product)
	})

	// Ambil data product dengan ID
	app.Get("/api/products/:id", func(c *fiber.Ctx) error {
		var product Product

		// ambil parameter id
		id := c.Params("id")
		if err := db.First(&product, id).Error; err != nil {
			return err
		}

		return c.JSON(product)
	})

	// Update data product
	app.Put("/api/products/:id", func(c *fiber.Ctx) error {
		var product Product

		// parameter id
		id := c.Params("id")
		if err := db.First(&product, id).Error; err != nil {
			return err
		}

		// parse body

		if err := c.BodyParser(&product); err != nil {
			return err
		}

		// simpan data product
		db.Save(&product)
		return c.JSON(product)
	})

	app.Delete("/api/products/id:", func(c *fiber.Ctx) error {
		//var product Product

		// parameter id
		id := c.Params("id")
		if err := db.Delete(&Product{}, id).Error; err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusNoContent)
	})

	port := "3000"

	app.Listen(":" + port)
}
