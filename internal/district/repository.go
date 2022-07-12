package district

import (
	"github.com/anousoneFS/go-workshop/internal/province"
	"gorm.io/gorm"
)

type Repository interface {
	GetByID(id uint) (District, error)
	Create(p *District) error
	GetAll() ([]District, error)
	Update(p District, id uint) error
	Delete(id uint) error
}

type District struct {
	gorm.Model
	Name       string
	NameEn     string
	Province   province.Province
	ProvinceID uint
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

type repository struct {
	db *gorm.DB
}

func (r repository) Create(p *District) error {
	return r.db.Create(&p).Error
}

func (r repository) Delete(id uint) error {
	return r.db.Delete(&District{}, id).Error
}

func (r repository) GetAll() ([]District, error) {
	var districts []District
	return districts, r.db.Find(&districts).Error
}

func (r repository) Update(p District, id uint) error {
	return r.db.Model(&District{}).Where("id = ?", id).Updates(p).Error
}

func (r repository) GetByID(id uint) (District, error) {
	var province District
	return province, r.db.First(&province, id).Error
}
