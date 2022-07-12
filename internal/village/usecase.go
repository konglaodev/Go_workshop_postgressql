package village

import (
	"errors"
	"fmt"
)

type usecase struct {
	repo Repository
}

type VillageRequest struct {
	ID         uint   `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	NameEn     string `json:"name_en,omitempty"`
	DistrictID uint   `json:"district_id,omitempty"`
}

type Usecase interface {
	Create(p VillageRequest) error
	GetAll() ([]VillageRequest, error)
	GetByID(id uint) (VillageRequest, error)
	Update(p VillageRequest, id uint) error
	Delete(id uint) error
}

func NewUsecase(repo Repository) Usecase {
	return &usecase{repo: repo}
}

func (u usecase) Create(body VillageRequest) error {
	village := Village{Name: body.Name, NameEn: body.NameEn, DistrictID: body.DistrictID}
	if err := u.repo.Create(&village); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (u usecase) GetAll() (i []VillageRequest, err error) {
	villages, err := u.repo.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range villages {
		new := VillageRequest{
			ID:         item.ID,
			Name:       item.Name,
			NameEn:     item.NameEn,
			DistrictID: item.DistrictID,
		}
		i = append(i, new)
	}
	return
}

func (u usecase) GetByID(id uint) (VillageRequest, error) {
	village, err := u.repo.GetByID(id)
	if err != nil {
		fmt.Println(err)
		return VillageRequest{}, err
	}
	response := VillageRequest{
		ID:         village.ID,
		Name:       village.Name,
		NameEn:     village.NameEn,
		DistrictID: village.DistrictID,
	}
	return response, nil
}

func (u usecase) Update(p VillageRequest, id uint) error {
	i := Village{}
	if p.ID != 0 {
		i.ID = p.ID
		return errors.New("invalid id")
	}
	if p.Name != "" {
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
