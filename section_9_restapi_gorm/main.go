package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID           int       `json:"id" gorm:"primarykey;autoIncrement"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password"`
	CreatedAt    time.Time `json:"created_at"`
}

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

	// Register user
	app.Post("/api/users/register", func(c *fiber.Ctx) error {
		var user User
		// parse request body
		// declare variable err dan inisialisasi dengan nilai dari bodyparse,
		// kemudian cek err apakah bernilai nil
		if err := c.BodyParser(&user); err != nil {
			return err
		}

		log.Println("Body Parser: ", &user)

		// pengecekan username atau email yang sudah ada
		var existingUser User
		if db.Where("username = ? OR email = ?", user.Username, user.Email).First(&existingUser).Error == nil {
			// User dengan username atau email sudah terdaftar
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": "Username atau email sudah terdaftar",
			})
		}

		fmt.Println(user.PasswordHash)
		// hash password sebelum simpan ke database
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		// assign value hashedPassword ke struct user field PasswordHash
		user.PasswordHash = string(hashedPassword)
		log.Println("Hashed Password: ", user.PasswordHash)

		// simpan ke database
		db.Create(&user)

		// mengembalikan response
		return c.JSON(user)
	})

	app.Post("/api/users/login", func(c *fiber.Ctx) error {
		// declare requestLogin dengan struct email dan password
		// kemudian body parse requestLogin
		var requestLogin struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.BodyParser(&requestLogin); err != nil {
			return err
		}

		// declare user untuk ambil data dari db
		// kemudian bandingkan nilainya dengan requestLogin
		var user User
		if err := db.Where("email = ? ", requestLogin.Email).Find(&user).Error; err != nil {
			return fiber.ErrNotFound
		}

		// compare password
		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(requestLogin.Password)); err != nil {
			// kembalikan response error jika password tidak cocok
			return fiber.ErrUnauthorized
		}

		return c.JSON(user)
	})

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
			log.Println("Body parser: ", product)
			log.Println("Error parsing request body:", err)
			return err
		}

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

	app.Delete("/api/products/:id", func(c *fiber.Ctx) error {

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
