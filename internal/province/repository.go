package province

import "gorm.io/gorm"

type Repository interface {
	GetByID(id uint) (Province, error)
	Create(p *Province) error
	GetAll() ([]Province, error)
	Update(p Province, id uint) error
	Delete(id uint) error
}

type Province struct {
	gorm.Model
	Name   string
	NameEn string
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

type repository struct {
	db *gorm.DB
}

func (r repository) Create(p *Province) error {
	return r.db.Create(&p).Error
}

func (r repository) Delete(id uint) error {
	return r.db.Delete(&Province{}, id).Error
}

func (r repository) GetAll() ([]Province, error) {
	var provinces []Province
	return provinces, r.db.Find(&provinces).Error
}

func (r repository) Update(p Province, id uint) error {
	return r.db.Model(&Province{}).Where("id = ?", id).Updates(p).Error
}

func (r repository) GetByID(id uint) (Province, error) {
	var province Province
	return province, r.db.First(&province, id).Error
}
