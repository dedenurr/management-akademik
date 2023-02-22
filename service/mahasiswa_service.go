package service

import (
	"github.com/dedenurr/management-akademik/entity"
	"github.com/dedenurr/management-akademik/repository"
)

type MahasiswaService interface {
	CreateMahasiswa(mahasiswa entity.Mahasiswa) (entity.Mahasiswa, error)
	ReadMahasiswa() ([]entity.Mahasiswa, error)
	UpdateMahasiswa(mahasiwa entity.Mahasiswa, nim string) (entity.Mahasiswa, error)
	DeleteMahasiswa(nim string) error
}



type mahasiswaService struct {
	mahasiswaRepo repository.MahasiswaRepository
}

func NewMahasiswaService(mahasiswaRepository repository.MahasiswaRepository) *mahasiswaService {
	return &mahasiswaService{mahasiswaRepository}
}

func (s *mahasiswaService) CreateMahasiswa(mahasiswa entity.Mahasiswa)(entity.Mahasiswa, error)  {
	var maha entity.Mahasiswa

	maha.Nim = mahasiswa.Nim
	maha.NamaMahasiswa = mahasiswa.NamaMahasiswa
	maha.TanggalLahir = mahasiswa.TanggalLahir
	maha.Alamat = mahasiswa.Alamat
	maha.JenisKelamin = mahasiswa.JenisKelamin

	newMaha, err := s.mahasiswaRepo.CreateMahasiswa(maha)
	if err != nil {
		return newMaha, err
	}

	return newMaha, nil
}

func (s *mahasiswaService) ReadMahasiswa() ([]entity.Mahasiswa, error) {
	maha, err := s.mahasiswaRepo.ReadMahasiswa()
	if err != nil {
		return maha, err
	}
	return maha, nil
}

func (s *mahasiswaService) UpdateMahasiswa(mahasiswa entity.Mahasiswa, nim string) (entity.Mahasiswa, error)  {
	var mah entity.Mahasiswa

	mah.Nim = nim
	mah.NamaMahasiswa = mahasiswa.NamaMahasiswa
	mah.TanggalLahir = mahasiswa.TanggalLahir
	mah.Alamat = mahasiswa.Alamat
	mah.JenisKelamin = mahasiswa.JenisKelamin

	newMah, err := s.mahasiswaRepo.UpdateMahasiswa(mah)
	if err != nil {
		return newMah, err
	}

	return newMah, err
}

func (s *mahasiswaService) DeleteMahasiswa(nim string)error {
	var maha entity.Mahasiswa

	maha.Nim = nim

	err := s.mahasiswaRepo.DeleteMahasiswa(maha)
	if err != nil {
		return err
	}

	return nil
}
