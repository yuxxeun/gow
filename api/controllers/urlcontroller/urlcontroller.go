package urlcontroller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yuxxeun/gow/api/models"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {
	var urls []models.Url
	if err := models.DB.Find(&urls).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengambil data",
		})
	}
	return c.JSON(fiber.Map{
		"data": urls,
	})
}

func Show(c *fiber.Ctx) error {
	var url models.Url
	id := c.Params("id")

	if err := models.DB.First(&url, id).Error; err != nil {
		status := http.StatusInternalServerError
		msg := "Terjadi kesalahan"
		if err == gorm.ErrRecordNotFound {
			status = http.StatusNotFound
			msg = "Data tidak ditemukan"
		}
		return c.Status(status).JSON(fiber.Map{
			"message": msg,
		})
	}

	return c.JSON(fiber.Map{
		"data": url,
	})
}

func Create(c *fiber.Ctx) error {
	var url models.Url
	if err := c.BodyParser(&url); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Format data tidak valid",
		})
	}

	if err := models.DB.Create(&url).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menyimpan data",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": url,
	})
}

func Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var url models.Url

	if err := c.BodyParser(&url); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Format data tidak valid",
		})
	}

	result := models.DB.Model(&models.Url{}).Where("id = ?", id).Updates(url)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengupdate data",
		})
	}
	if result.RowsAffected == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Data tidak ditemukan",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil diupdate",
	})
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	result := models.DB.Delete(&models.Url{}, id)
	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus data",
		})
	}
	if result.RowsAffected == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Data tidak ditemukan",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil dihapus",
	})
}
