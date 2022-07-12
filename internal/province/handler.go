package province

import (
	"errors"
	"fmt"
	"html"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

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

type handler struct {
	uc Usecase
}

func NewHandler(uc Usecase, app *fiber.App) {
	h := &handler{uc: uc}
	api := app.Group("/api/v1")
	api.Get("/provinces", h.GetAll)
	api.Post("/provinces", h.Create)
	api.Get("/provinces/:id", h.GetByID)
	api.Patch("/provinces/:id", h.Update)
	api.Delete("/provinces/:id", h.Delete)
}

func (h handler) Create(c *fiber.Ctx) error {
	var body ProvinceRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid body"})
	}
	body.EscapeWhiteSpace()
	if err := body.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid body"})
	}
	if err := h.uc.Create(body); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unexpected error"})
	}
	// todo: get province
	return c.Status(fiber.StatusOK).JSON(body)
}

func (h handler) GetAll(c *fiber.Ctx) error {
	i, err := h.uc.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unexpected error"})
	}
	return c.Status(fiber.StatusOK).JSON(i)
}

func (h handler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	u64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	wd := uint(u64)
	i, err := h.uc.GetByID(wd)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unexpected error"})
	}
	return c.Status(fiber.StatusOK).JSON(i)
}

func (h handler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	u64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	wd := uint(u64)
	var body ProvinceRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid body"})
	}
	body.EscapeWhiteSpace()
	if err := body.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid body"})
	}
	err = h.uc.Update(body, wd)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unexpected error"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "successfully."})
}

func (h handler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	u64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	wd := uint(u64)
	if err := h.uc.Delete(wd); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "unexpected error"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "successfully."})
}
