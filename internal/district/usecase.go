package district

import (
	"errors"
	"fmt"
)

type usecase struct {
	repo Repository
}

type DistrictRequest struct {
	ID         uint   `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	NameEn     string `json:"name_en,omitempty"`
	ProvinceID uint   `json:"province_id,omitempty"`
}

type Usecase interface {
	Create(p DistrictRequest) error
	GetAll() ([]DistrictRequest, error)
	GetByID(id uint) (DistrictRequest, error)
	Update(p DistrictRequest, id uint) error
	Delete(id uint) error
}

func NewUsecase(repo Repository) Usecase {
	return &usecase{repo: repo}
}

func (u usecase) Create(body DistrictRequest) error {
	province := District{Name: body.Name, NameEn: body.NameEn, ProvinceID: body.ProvinceID}
	if err := u.repo.Create(&province); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (u usecase) GetAll() (i []DistrictRequest, err error) {
	districts, err := u.repo.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range districts {
		new := DistrictRequest{
			ID:         item.ID,
			Name:       item.Name,
			NameEn:     item.NameEn,
			ProvinceID: item.ProvinceID,
		}
		i = append(i, new)
	}
	return
}

func (u usecase) GetByID(id uint) (DistrictRequest, error) {
	district, err := u.repo.GetByID(id)
	if err != nil {
		fmt.Println(err)
		return DistrictRequest{}, err
	}
	response := DistrictRequest{
		ID:         district.ID,
		Name:       district.Name,
		NameEn:     district.NameEn,
		ProvinceID: district.ProvinceID,
	}
	return response, nil
}

func (u usecase) Update(p DistrictRequest, id uint) error {
	i := District{}
	if p.ID != 0 {
		i.ID = p.ID
		return errors.New("invalid id")
	}
	if p.Name != "" {
		i.Name = p.Name
	}
	if p.NameEn != "" {
		i.NameEn = p.NameEn
	}
	if err := u.repo.Update(i, id); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (u usecase) Delete(id uint) error {
	if err := u.repo.Delete(id); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
