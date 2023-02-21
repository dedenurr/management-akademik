package repository

import (
	"database/sql"
	"log"

	"github.com/dedenurr/management-akademik/entity"
)

type DosenRepository interface {
	CreateDosen(dosen entity.Dosen) (entity.Dosen, error)
	ReadDosen() ([]entity.Dosen, error)
	UpdateDosen(dosen entity.Dosen) (entity.Dosen, error)
	DeleteDosen(dosen entity.Dosen) error
}

type dosenRepository struct {
	db *sql.DB
}

func NewDosenRepository(db *sql.DB) *dosenRepository {
	return &dosenRepository{db}
}

func (r *dosenRepository) CreateDosen(dosen entity.Dosen) (entity.Dosen, error)  {
	sqlStatement := `INSERT INTO dosen (nip, nama_dosen) VALUES ($1, $2) RETURNING *`
	err := r.db.QueryRow(sqlStatement,
		dosen.Nip,
		dosen.NamaDosen).Scan(
		&dosen.Nip,
		&dosen.NamaDosen)
	if err != nil {
		return dosen,err
	}
	return dosen, nil
}

func (r *dosenRepository) ReadDosen() ([]entity.Dosen, error)  {
	var result []entity.Dosen
	sqlStatement := "SELECT * FROM dosen"
	data, err := r.db.Query(sqlStatement)
		if err != nil {
		log.Fatal(err)
	}


	defer data.Close()

	for data.Next(){
		var dosen entity.Dosen

		err := data.Scan(&dosen.Nip, &dosen.NamaDosen)
			if err != nil {
		log.Fatal(err)
	}


		result = append(result, dosen)
	}
	return result, nil
}

func (r *dosenRepository) UpdateDosen(dosen entity.Dosen) (entity.Dosen, error) {
	sqlStatement := `
	UPDATE dosen
	SET nama_dosen = $1
	WHERE nip = $2 
	`
	err := r.db.QueryRow(
		sqlStatement,
		dosen.NamaDosen,
		dosen.Nip).Scan(
			&dosen.Nip,
			&dosen.NamaDosen)
	if err != nil {
		return dosen, err
	}
	return dosen, nil
}


func (r *dosenRepository) DeleteDosen(dosen entity.Dosen)error  {
	sqlStatement := "DELETE FROM dosen WHERE nip =$1"

	err := r.db.QueryRow(sqlStatement, dosen.Nip)

	return err.Err()
}