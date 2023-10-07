package route

import (
	"uts/database"
	"uts/models"

	"github.com/gofiber/fiber/v2"
)

// Buatlah endpoint Insert Data sesuai dengan Collection Postman
func InsertData(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	database.DB.Create(&user)

	return c.JSON(fiber.Map{
		"Pesan": "Data telah berhasil di tambahkan",
	})
}

func GetAllData(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Find(&users)

	return c.JSON(fiber.Map{
		"data": users,
	})
}

func GetUserByid(c *fiber.Ctx) error {
	id_user := c.Params("id_user")
	var user models.User

	if err := database.DB.First(&user, id_user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Data user tidak ditemukan",
		})
	}

	return c.JSON(fiber.Map{
		"data": user,
	})
}

func Delete(c *fiber.Ctx) error {
	id_user := c.Params("id_user")

	if err := database.DB.Delete(&models.User{}, id_user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Data user tidak ditemukan",
		})
	}

	return c.JSON(fiber.Map{
		"Pesan": "Data telah dihapus",
	})
}

// Mengupdate data user berdasarkan id_user yang di tempatkan di parameter
func Update(c *fiber.Ctx) error {
	id_user := c.Params("id_user")

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	if err := database.DB.First(&user, id_user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Data user tidak ditemukan",
		})
	}

	update := models.User{
		Nama:     data["nama"],
		Email:    data["email"],
		Password: data["password"],
	}

	database.DB.Model(&user).Updates(update)

	return c.JSON(fiber.Map{
		"Pesan": "Data User telah di Update",
	})
}
