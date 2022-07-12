package main

import (
	"errors"
	"html"
	"strings"

	"github.com/anousoneFS/go-workshop/internal/province"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type District struct {
	gorm.Model
	Name       string `json:"name"`
	NameEn     string `json:"name_en"`
	Province   province.Province
	ProvinceID uint
}

type DistrictRequest struct {
	ID         uint   `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	NameEn     string `json:"name_en,omitempty"`
	ProvinceID uint   `json:"province_id,omitempty"`
}

type DistrictResponse = DistrictRequest

func (p *DistrictRequest) EscapeWhiteSpace() {
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.NameEn = html.EscapeString(strings.TrimSpace(p.NameEn))
}

func (p DistrictRequest) Validate() error {
	if p.Name == "" || p.NameEn == "" {
		return errors.New("validate failed")
	}
	return nil
}

func CreateDistrict(c *fiber.Ctx) error {
	var body DistrictRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid body"})
	}
	body.EscapeWhiteSpace()
	if err := body.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid body"})
	}
	district := District{Name: body.Name, NameEn: body.NameEn, ProvinceID: body.ProvinceID}
	if err := db.Create(&district).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unexpect error"})
	}
	if err := db.Where("id = ?", district.ID).Find(&District{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unexpect error"})
	}
	response := DistrictResponse{ID: district.ID, Name: district.Name, NameEn: district.NameEn, ProvinceID: district.ProvinceID}
	return c.Status(fiber.StatusOK).JSON(response)
}

func GetAllDistrict(c *fiber.Ctx) error {
	var district []District
	if err := db.Find(&district).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unexpect error"})
	}
	var response []DistrictResponse
	for _, i := range district {
		response = append(response, DistrictResponse{
			ID:         i.ID,
			Name:       i.Name,
			NameEn:     i.NameEn,
			ProvinceID: i.ProvinceID,
		})
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func GetDistrictByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var district District
	if err := db.Find(&district, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unexpect error"})
	}
	response := DistrictResponse{ID: district.ID, Name: district.Name, NameEn: district.NameEn, ProvinceID: district.ProvinceID}
	return c.Status(fiber.StatusOK).JSON(response)
}

func UpdateDistrictByID(c *fiber.Ctx) error {
	var body DistrictRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid body"})
	}
	body.EscapeWhiteSpace()
	// get
	var i District
	if err := db.Find(&i, body.ID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unexpect error"})
	}
	if body.Name != "" {
		i.Name = body.Name
	}
	if body.NameEn != "" {
		i.NameEn = body.NameEn
	}
	if body.ProvinceID != 0 {
		i.ProvinceID = body.ProvinceID
	}
	if err := db.Save(&i).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unexpect error"})
	}
	response := DistrictResponse{ID: i.ID, Name: i.Name, NameEn: i.NameEn}
	return c.Status(fiber.StatusOK).JSON(response)
}

func DeleteDistric(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := db.Where("id = ?", id).Delete(&District{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unexpect error"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "successfull."})
}
