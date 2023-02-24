package service

import (
	"github.com/dedenurr/management-akademik/entity"
	"github.com/dedenurr/management-akademik/helper"
	"github.com/dedenurr/management-akademik/repository"
)

type PerkuliahanService interface {
	CreatePerkuliahan(inputPerkuliahan entity.InputPerkuliahan) (entity.Perkuliahan, error)
	ReadPerkuliahan() ([]entity.Perkuliahan, error)
	UpdatePerkuliahan(inputPerkuliahan entity.InputPerkuliahan, id int) (entity.Perkuliahan, error)
	DeletePerkuliahan(id int) error
}



type perkuliahanService struct {
	perkuliahanRepo repository.PerkuliahanRepository
}

func NewPerkuliahanService(perkuliahanRepository repository.PerkuliahanRepository) *perkuliahanService {
	return &perkuliahanService{perkuliahanRepository}
}

func (s *perkuliahanService) CreatePerkuliahan(inputPerkuliahan entity.InputPerkuliahan) (entity.Perkuliahan, error)  {
	var pkh entity.Perkuliahan

	pkh.Nim = inputPerkuliahan.Nim
	pkh.KodeMatakuliah = inputPerkuliahan.KodeMatakuliah
	pkh.Nip = inputPerkuliahan.Nip
	pkh.Nilai = inputPerkuliahan.Nilai
	pkh.Grade = helper.FormatNilai(pkh.Nilai)
	
	newPkh, err := s.perkuliahanRepo.CreatePerkuliahan(pkh)
	if err != nil {
		return newPkh, err
	}

	return newPkh, nil
}

func (s *perkuliahanService) ReadPerkuliahan() ([]entity.Perkuliahan, error) {
	pkh, err := s.perkuliahanRepo.ReadPerkuliahan()
	if err != nil {
		return pkh, err
	}
	return pkh, nil
}

func (s *perkuliahanService) UpdatePerkuliahan(inputPerkuliahan entity.InputPerkuliahan, id int) (entity.Perkuliahan, error) {
	var pkh entity.Perkuliahan

	pkh.Nim = inputPerkuliahan.Nim
	pkh.KodeMatakuliah = inputPerkuliahan.KodeMatakuliah
	pkh.Nip = inputPerkuliahan.Nip
	pkh.Nilai = inputPerkuliahan.Nilai
	pkh.Grade = helper.FormatNilai(pkh.Nilai)
	pkh.Id = id

	newPkh, err := s.perkuliahanRepo.UpdatePerkuliahan(pkh)
	if err != nil {
		return newPkh, err
	}

	return newPkh, err	
}

func (s *perkuliahanService) DeletePerkuliahan(id int)error {
	var pkh entity.Perkuliahan

	pkh.Id = id

	err := s.perkuliahanRepo.DeletePerkuliahan(pkh)
	if err != nil {
		return err
	}

	return nil
}