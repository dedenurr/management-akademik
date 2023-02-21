package service

import (
	"github.com/dedenurr/management-akademik/entity"
	"github.com/dedenurr/management-akademik/repository"
)

type DosenService interface {
	CreateDosen(dosen entity.Dosen) (entity.Dosen, error)
	ReadDosen() ([]entity.Dosen, error)
	UpdateDosen(dosen entity.Dosen, nip string) (entity.Dosen, error)
	DeleteDosen(nip string) error
}



type dosenService struct {
	dosenRepo repository.DosenRepository
}

func NewDosenService(dosenRepository repository.DosenRepository) *dosenService {
	return &dosenService{dosenRepository}
}

func (s *dosenService) CreateDosen(dosen entity.Dosen)(entity.Dosen, error)  {
	var dos entity.Dosen

	dos.Nip = dosen.Nip
	dos.NamaDosen = dosen.NamaDosen

	newDos, err := s.dosenRepo.CreateDosen(dos)
	if err != nil {
		return newDos, err
	}

	return newDos, nil
}

func (s *dosenService) ReadDosen() ([]entity.Dosen, error) {
	dos, err := s.dosenRepo.ReadDosen()
	if err != nil {
		return dos, err
	}
	return dos, nil
}

func (s *dosenService) UpdateDosen(dosen entity.Dosen, nip string) (entity.Dosen, error)  {
	var dos entity.Dosen

	dos.Nip = nip
	dos.NamaDosen = dosen.NamaDosen

	newDos, err := s.dosenRepo.UpdateDosen(dos)
	if err != nil {
		return newDos, err
	}

	return newDos, err	
}

func (s *dosenService) DeleteDosen(nip string)error {
	var dos entity.Dosen

	dos.Nip = nip

	err := s.dosenRepo.DeleteDosen(dos)
	if err != nil {
		return err
	}

	return nil
}