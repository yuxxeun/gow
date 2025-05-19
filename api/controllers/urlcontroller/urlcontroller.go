package urlcontroller

import (
	"net/http"

	"github.com/yuxxeun/gow/server/models"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	var url []models.Url
	models.DB.Find(&url)

	return c.JSON(fiber.Map{
		"data": url,
	})
}

func Show(c *fiber.Ctx) error {

	id := c.Params("id")
	var url models.Url
	if err := models.DB.First(&url, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Data tidak ditemukan",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Data tidak ditemukan",
		})
	}

	return c.JSON(url)
}

func Create(c *fiber.Ctx) error {

	var book models.Url
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := models.DB.Create(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(book)
}

func Update(c *fiber.Ctx) error {

	id := c.Params("id")

	var url models.Url
	if err := c.BodyParser(&url); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if models.DB.Where("id = ?", id).Updates(&url).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Tidak dapat mengupdate data",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil diupdate",
	})
}

func Delete(c *fiber.Ctx) error {

	id := c.Params("id")

	var book models.Url
	if models.DB.Delete(&book, id).RowsAffected == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Tidak dapat menghapus data",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil dihapus",
	})
}
