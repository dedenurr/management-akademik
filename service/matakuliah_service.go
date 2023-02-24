package service

import (
	"github.com/dedenurr/management-akademik/entity"
	"github.com/dedenurr/management-akademik/repository"
)

type MataKuliahService interface {
	CreateMataKuliah(mataKuliah entity.MataKuliah) (entity.MataKuliah, error)
	ReadMataKuliah() ([]entity.MataKuliah, error)
	UpdateMataKuliah(mataKuliah entity.MataKuliah, id string) (entity.MataKuliah, error)
	DeleteMataKuliah(id string) error
}

type mataKuliahService struct {
	mataKuliahRepo repository.MataKuliahRepository
}

func NewMataKuliahService(mataKuliahRepository repository.MataKuliahRepository) *mataKuliahService {
	return &mataKuliahService{mataKuliahRepository}
}

func (s *mataKuliahService) CreateMataKuliah(mataKuliah entity.MataKuliah) (entity.MataKuliah, error) {
	var mKul entity.MataKuliah

	mKul.KodeMatakuliah = mataKuliah.KodeMatakuliah
	mKul.NamaMatakuliah = mataKuliah.NamaMatakuliah
	mKul.Sks = mataKuliah.Sks

	newMKul, err := s.mataKuliahRepo.CreateMataKuliah(mKul)
	if err != nil {
		return newMKul, err
	}
	return newMKul, nil
	
}

func (s *mataKuliahService) ReadMataKuliah() ([]entity.MataKuliah, error) {
	mKul, err := s.mataKuliahRepo.ReadMataKuliah()
	if err != nil {
		return mKul, err
	}
	return mKul, nil
}

func (s *mataKuliahService) UpdateMataKuliah(mataKuliah entity.MataKuliah, id string) (entity.MataKuliah, error)  {
	var mKul entity.MataKuliah

	mKul.KodeMatakuliah = id
	mKul.NamaMatakuliah = mataKuliah.NamaMatakuliah
	mKul.Sks = mataKuliah.Sks

	newMKul, err := s.mataKuliahRepo.UpdateMataKuliah(mKul)
	if err != nil {
		return newMKul, err
	}

	return newMKul, err	
}

func (s *mataKuliahService) DeleteMataKuliah(id string) error {
	var mKul entity.MataKuliah

	mKul.KodeMatakuliah = id

	err := s.mataKuliahRepo.DeleteMataKuliah(mKul)
	if err != nil {
		return err
	}

	return nil
}
