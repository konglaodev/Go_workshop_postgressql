package main

import (
	"errors"
	"html"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Village struct {
	gorm.Model
	Name       string `json:"name"`
	NameEn     string `json:"name_en"`
	District   District
	DistrictID uint
}

type VillageRequest struct {
	ID         uint   `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	NameEn     string `json:"name_en,omitempty"`
	DistrictID uint   `json:"district_id,omitempty"`
}

type VillageResponse = VillageRequest

func (p *VillageRequest) EscapeWhiteSpace() {
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.NameEn = html.EscapeString(strings.TrimSpace(p.NameEn))
}

func (p VillageRequest) Validate() error {
	if p.Name == "" || p.NameEn == "" {
		return errors.New("validate failed")
	}
	return nil
}

func CreateVillage(c *fiber.Ctx) error {
	var body VillageRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid body"})
	}
	body.EscapeWhiteSpace()
	if err := body.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid body"})
	}
	village := Village{Name: body.Name, NameEn: body.NameEn, DistrictID: body.DistrictID}
	if err := db.Create(&village).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unexpect error"})
	}
	if err := db.Where("id = ?", village.ID).Find(&Village{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unexpect error"})
	}
	response := VillageResponse{ID: village.ID, Name: village.Name, NameEn: village.NameEn, DistrictID: village.DistrictID}
	return c.Status(fiber.StatusOK).JSON(response)
}

func GetAllVillage(c *fiber.Ctx) error {
	var village []Village
	if err := db.Find(&village).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unexpect error"})
	}
	var response []VillageResponse
	for _, i := range village {
		response = append(response, VillageResponse{
			ID:         i.ID,
			Name:       i.Name,
			NameEn:     i.NameEn,
			DistrictID: i.DistrictID,
		})
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func GetVillageByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var village Village
	if err := db.Find(&village, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unexpect error"})
	}
	response := VillageResponse{ID: village.ID, Name: village.Name, NameEn: village.NameEn, DistrictID: village.DistrictID}
	return c.Status(fiber.StatusOK).JSON(response)
}

func UpdateVillage(c *fiber.Ctx) error {
	var body VillageRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid body"})
	}
	body.EscapeWhiteSpace()
	// get
	var i Village
	if err := db.Find(&i, body.ID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unexpect error"})
	}
	if body.Name != "" {
		i.Name = body.Name
	}
	if body.NameEn != "" {
		i.NameEn = body.NameEn
	}
	if body.DistrictID != 0 {
		i.DistrictID = body.DistrictID
	}
	if err := db.Save(&i).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unexpect error"})
	}
	response := VillageResponse{ID: i.ID, Name: i.Name, NameEn: i.NameEn}
	return c.Status(fiber.StatusOK).JSON(response)
}

func DeleteVillage(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := db.Where("id = ?", id).Delete(&Village{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unexpect error"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "successfull."})
}
