package repository

import (
	"database/sql"
	"log"

	"github.com/dedenurr/management-akademik/entity"
)

type MataKuliahRepository interface {
	CreateMataKuliah(mataKuliah entity.MataKuliah) (entity.MataKuliah, error)
	ReadMataKuliah() ([]entity.MataKuliah, error)
	UpdateMataKuliah(mataKuliah entity.MataKuliah) (entity.MataKuliah, error)
	DeleteMataKuliah(mataKuliah entity.MataKuliah) error
}

type mataKuliahRepository struct {
	db *sql.DB
}

func NewMataKuliahRepository(db *sql.DB) *mataKuliahRepository {
	return &mataKuliahRepository{db}
}

func (r *mataKuliahRepository) CreateMataKuliah(mataKuliah entity.MataKuliah) (entity.MataKuliah, error)  {
	sqlStatement := `INSERT INTO mataKuliah (kode_mata_kuliah, nama_mata_kuliah, sks) VALUES ($1, $2, $3) RETURNING *`
	err := r.db.QueryRow(sqlStatement,
		mataKuliah.KodeMatakuliah,
		mataKuliah.NamaMatakuliah,
		mataKuliah.Sks).Scan(
		&mataKuliah.KodeMatakuliah,
		&mataKuliah.NamaMatakuliah,
		&mataKuliah.Sks)
	if err != nil {
		return mataKuliah,err
	}
	return mataKuliah, nil
}

func (r *mataKuliahRepository) ReadMataKuliah() ([]entity.MataKuliah, error)  {
	var result []entity.MataKuliah
	sqlStatement := "SELECT * FROM matakuliah"
	data, err := r.db.Query(sqlStatement)
		if err != nil {
		log.Fatal(err)
	}

	defer data.Close()
	for data.Next(){
		var mataKuliah entity.MataKuliah

		err := data.Scan(&mataKuliah.KodeMatakuliah, &mataKuliah.NamaMatakuliah, &mataKuliah.Sks)
			if err != nil {
		log.Fatal(err)
	}
		result = append(result, mataKuliah)
	}
	return result, nil
}

func (r *mataKuliahRepository) UpdateMataKuliah(mataKuliah entity.MataKuliah) (entity.MataKuliah, error){
	sqlStatement := `
	UPDATE matakuliah
	SET  nama_mata_kuliah = $1,
	sks = $2
	WHERE kode_mata_kuliah = $3
	`
	err := r.db.QueryRow(
		sqlStatement,
		mataKuliah.NamaMatakuliah,
		mataKuliah.Sks,
		mataKuliah.KodeMatakuliah).Scan(
			&mataKuliah.NamaMatakuliah,
			&mataKuliah.Sks,
			&mataKuliah.KodeMatakuliah)
	if err != nil {
		return mataKuliah, err
	}
	return mataKuliah, nil
}

func (r *mataKuliahRepository) DeleteMataKuliah(mataKuliah entity.MataKuliah)error  {
	sqlStatement := "DELETE FROM matakuliah WHERE kode_mata_kuliah =$1"

	err := r.db.QueryRow(sqlStatement, mataKuliah.KodeMatakuliah)

	return err.Err()
}