package main

import (
	
	"khub/database"
	"khub/route"
	"log"
	"github.com/joho/godotenv"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	database.Connect()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	//gruping router
	member := app.Group("/member") // /api
	product := app.Group("/product") // /
	//router
	member.Post("/insert_member", route.Insert_member)
	member.Put("/update_member/:id_member", route.Update_member)
	member.Get("/getall_member", route.Getall_member)
	member.Get("/get_data_member/:id_member", route.Getdatamember)
	member.Delete("/delete_member/:id_member", route.Delete_member)
	product.Get("/reviews/:id_product", route.Review_produk)
	product.Post("/insert_review/:id_member/:id_product", route.Insert_reviews_produk)
	product.Get("/getlike/:id_product/:id_member", route.Cek_like_product)
	product.Put("/like/:id_member/:id_product", route.Like_unlike_product)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(os.Getenv("PORT")))
}
