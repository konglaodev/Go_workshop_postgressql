package main

import (
	"errors"
	"html"
	"strings"

	"github.com/anousoneFS/go-workshop/internal/province"
	"github.com/gofiber/fiber/v2"
)

// province

type ProvinceResponse struct {
	ID     uint
	Name   string
	NameEn string
}

type ProvinceRequest struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	NameEn string `json:"name_en"`
}

func (p *ProvinceRequest) EscapeWhiteSpace() {
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.NameEn = html.EscapeString(strings.TrimSpace(p.NameEn))
}

func (p ProvinceRequest) Validate() error {
	if p.Name == "" || p.NameEn == "" {
		return errors.New("validate failed")
	}
	return nil
}

func CreateProvince(c *fiber.Ctx) error {
	var body ProvinceRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid body"})
	}
	body.EscapeWhiteSpace()
	if err := body.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid body"})
	}

	province := province.Province{Name: body.Name, NameEn: body.NameEn}
	if err := db.Create(&province).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unexpect error"})
	}
	// todo: get province
	return c.Status(fiber.StatusCreated).JSON(province)
}

func GetProvinceByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var province province.Province
	if err := db.Find(&province, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unexpect error"})
	}
	response := ProvinceResponse{ID: province.ID, Name: province.Name, NameEn: province.NameEn}
	return c.Status(fiber.StatusOK).JSON(response)
}

func GetAllProvince(c *fiber.Ctx) error {
	var provinces []province.Province
	if err := db.Find(&provinces).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unexpect error"})
	}
	var newProvinces []ProvinceResponse
	for _, i := range provinces {
		newProvinces = append(newProvinces, ProvinceResponse{
			ID:     i.ID,
			Name:   i.Name,
			NameEn: i.NameEn,
		})
	}
	return c.Status(fiber.StatusOK).JSON(newProvinces)
}

func UpdateProvince(c *fiber.Ctx) error {
	var body ProvinceRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid body"})
	}
	body.EscapeWhiteSpace()
	// get
	var i province.Province
	if err := db.Find(&i, body.ID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unexpect error"})
	}
	if body.Name != "" {
		i.Name = body.Name
	}
	if body.NameEn != "" {
		i.NameEn = body.NameEn
	}
	if err := db.Save(&i).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unexpect error"})
	}
	response := ProvinceResponse{ID: i.ID, Name: i.Name, NameEn: i.NameEn}
	return c.Status(fiber.StatusOK).JSON(response)
}

func DeleteProvince(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := db.Where("id = ?", id).Delete(&province.Province{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unexpect error"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "successfull."})
}
