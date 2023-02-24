package repository

import (
	"database/sql"
	"log"

	"github.com/dedenurr/management-akademik/entity"
	"github.com/dedenurr/management-akademik/helper"
)

type PerkuliahanRepository interface {
	CreatePerkuliahan(perkuliahan entity.Perkuliahan) (entity.Perkuliahan, error)
	ReadPerkuliahan() ([]entity.Perkuliahan, error)
	UpdatePerkuliahan(perkuliahan entity.Perkuliahan) (entity.Perkuliahan, error)
	DeletePerkuliahan(perkuliahan entity.Perkuliahan) error
}

type perkuliahanRepository struct {
	db *sql.DB
}

func NewPerkuliahanRepository(db *sql.DB) *perkuliahanRepository {
	return &perkuliahanRepository{db}
}

func (r *perkuliahanRepository) CreatePerkuliahan(perkuliahan entity.Perkuliahan) (entity.Perkuliahan, error)  {
	sqlStatement := `INSERT INTO perkuliahan (nim, kode_mata_kuliah, nip, nilai, grade) VALUES ($1, $2, $3, $4, $5) RETURNING *`
	err := r.db.QueryRow(sqlStatement,
		perkuliahan.Nim,
		perkuliahan.KodeMatakuliah,
		perkuliahan.Nip,
		perkuliahan.Nilai,
		perkuliahan.Grade).Scan(
			&perkuliahan.Nim,
			&perkuliahan.KodeMatakuliah,
			&perkuliahan.Nip,
			&perkuliahan.Nilai,
			&perkuliahan.Grade,
			&perkuliahan.Id)
/* 	if err != nil {
		return perkuliahan, err
	} */
	helper.PanicIfError(err)
	return perkuliahan, nil
}

func (r *perkuliahanRepository) ReadPerkuliahan() ([]entity.Perkuliahan, error)  {
	var result []entity.Perkuliahan
	sqlStatement := "SELECT * FROM perkuliahan"
	data, err := r.db.Query(sqlStatement)
		if err != nil {
		log.Fatal(err)
	}

	defer data.Close()

	for data.Next(){
		var perkuliahan entity.Perkuliahan

		err := data.Scan(&perkuliahan.Nim,
			&perkuliahan.KodeMatakuliah,
			&perkuliahan.Nip,
			&perkuliahan.Nilai,
			&perkuliahan.Grade,
			&perkuliahan.Id)
			if err != nil {
				log.Fatal(err)
			}


		result = append(result, perkuliahan)
	}
	return result, nil
}

func (r *perkuliahanRepository) UpdatePerkuliahan(perkuliahan entity.Perkuliahan) (entity.Perkuliahan, error) {
	sqlStatement := `
	UPDATE perkuliahan
	SET nim = $1,
    kode_mata_kuliah = $2,
    nip = $3,
    nilai = $4,
    grade = $5
    WHERE id = $6 
	`
	err := r.db.QueryRow(
		sqlStatement,
		perkuliahan.Nim,
		perkuliahan.KodeMatakuliah,
		perkuliahan.Nip,
		perkuliahan.Nilai,
		perkuliahan.Grade,
		perkuliahan.Id).Scan(
			&perkuliahan.Nim,
			&perkuliahan.KodeMatakuliah,
			&perkuliahan.Nip,
			&perkuliahan.Nilai,
			&perkuliahan.Grade,
			&perkuliahan.Id)
	if err != nil {
		return perkuliahan, err
	}
	return perkuliahan, nil
}

func (r *perkuliahanRepository) DeletePerkuliahan(perkuliahan entity.Perkuliahan)error  {
	sqlStatement := "DELETE FROM perkuliahan WHERE id =$1"

	err := r.db.QueryRow(sqlStatement,perkuliahan.Id)

	return err.Err()
}