package village

import (
	"github.com/anousoneFS/go-workshop/internal/district"
	"gorm.io/gorm"
)

type Repository interface {
	GetByID(id uint) (Village, error)
	Create(p *Village) error
	GetAll() ([]Village, error)
	Update(p Village, id uint) error
	Delete(id uint) error
}

type Village struct {
	gorm.Model
	Name       string
	NameEn     string
	District   district.District
	DistrictID uint
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

type repository struct {
	db *gorm.DB
}

func (r repository) Create(p *Village) error {
	return r.db.Create(&p).Error
}

func (r repository) Delete(id uint) error {
	return r.db.Delete(&Village{}, id).Error
}

func (r repository) GetAll() ([]Village, error) {
	var villages []Village
	return villages, r.db.Find(&villages).Error
}

func (r repository) Update(p Village, id uint) error {
	return r.db.Model(&Village{}).Where("id = ?", id).Updates(p).Error
}

func (r repository) GetByID(id uint) (Village, error) {
	var village Village
	return village, r.db.First(&village, id).Error
}
