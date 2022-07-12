package province

import (
	"errors"
	"fmt"
)

type usecase struct {
	repo Repository
}

type Usecase interface {
	Create(p ProvinceRequest) error
	GetAll() ([]ProvinceRequest, error)
	GetByID(id uint) (ProvinceRequest, error)
	Update(p ProvinceRequest, id uint) error
	Delete(id uint) error
}

func NewUsecase(repo Repository) Usecase {
	return &usecase{repo: repo}
}

func (u usecase) Create(body ProvinceRequest) error {
	province := Province{Name: body.Name, NameEn: body.NameEn}
	if err := u.repo.Create(&province); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (u usecase) GetAll() ([]ProvinceRequest, error) {
	i, err := u.repo.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	var responses []ProvinceRequest
	for _, item := range i {
		response := ProvinceRequest{
			ID:     item.ID,
			Name:   item.Name,
			NameEn: item.NameEn,
		}
		responses = append(responses, response)
	}
	return responses, nil
}

func (u usecase) GetByID(id uint) (ProvinceRequest, error) {
	i, err := u.repo.GetByID(id)
	if err != nil {
		fmt.Println(err)
		return ProvinceRequest{}, err
	}
	response := ProvinceRequest{
		ID:     i.ID,
		Name:   i.Name,
		NameEn: i.NameEn,
	}
	return response, nil
}

func (u usecase) Update(p ProvinceRequest, id uint) error {
	i := Province{}
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
