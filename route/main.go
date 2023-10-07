package main

import (
	"log"
	"uts/database"
	"uts/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Membuat objek Fiber
	app := fiber.New()

	// Menghubungkan ke database
	database.Connect()

	// Mengaitkan route dengan endpoint
	app.Post("/insert", route.InsertData)
	app.Get("/getData", route.GetAllData)
	app.Get("/getDataUser/:id_user", route.GetUserByid)
	app.Get("/delete/:id_user", route.Delete)
	app.Put("/update/:id_user", route.Update)

	// Menjalankan aplikasi pada port 3000
	log.Fatal(app.Listen(":3000"))
}
