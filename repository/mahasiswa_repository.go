package repository

import (
	"database/sql"
	"log"

	"github.com/dedenurr/management-akademik/entity"
)

type MahasiswaRepository interface {
	CreateMahasiswa(mahasiswa entity.Mahasiswa) (entity.Mahasiswa, error)
	ReadMahasiswa() ([]entity.Mahasiswa, error)
	UpdateMahasiswa(mahasiswa entity.Mahasiswa) (entity.Mahasiswa, error)
	DeleteMahasiswa(mahasiswa entity.Mahasiswa) error
}

type mahasiswaRepository struct {
	db *sql.DB
}

func NewMahasiswaRepository(db *sql.DB) *mahasiswaRepository {
	return &mahasiswaRepository{db}
}

func (r *mahasiswaRepository) CreateMahasiswa(mahasiswa entity.Mahasiswa) (entity.Mahasiswa, error)  {
	sqlStatement := `INSERT INTO mahasiswa (nim, nama_Mahasiswa, tanggal_lahir, alamat, jenis_kelamin) VALUES ($1, $2, $3, $4) RETURNING *`
	err := r.db.QueryRow(sqlStatement,
		mahasiswa.Nim,
		mahasiswa.NamaMahasiswa,
		mahasiswa.TanggalLahir,
		mahasiswa.Alamat,
		mahasiswa.JenisKelamin).Scan(
		&mahasiswa.Nim,
		&mahasiswa.NamaMahasiswa,
		&mahasiswa.TanggalLahir,
		&mahasiswa.Alamat,
		&mahasiswa.JenisKelamin)
	if err != nil {
		return mahasiswa,err
	}
	return mahasiswa, nil
}

func (r *mahasiswaRepository) ReadMahasiswa() ([]entity.Mahasiswa, error)  {
	var result []entity.Mahasiswa
	sqlStatement := "SELECT * FROM mahasiswa"
	data, err := r.db.Query(sqlStatement)
		if err != nil {
		log.Fatal(err)
	}


	defer data.Close()

	for data.Next(){
		var mahasiswa entity.Mahasiswa

		err := data.Scan(&mahasiswa.Nim,
			&mahasiswa.NamaMahasiswa,
			&mahasiswa.TanggalLahir,
			&mahasiswa.Alamat,
			&mahasiswa.JenisKelamin)
			if err != nil {
		log.Fatal(err)
	}


		result = append(result, mahasiswa)
	}
	return result, nil
}

func (r *mahasiswaRepository) UpdateMahasiswa(mahasiswa entity.Mahasiswa) (entity.Mahasiswa, error) {
	sqlStatement := `
	UPDATE mahasiswa
	SET nama_mahasiswa = $1, tanggal_lahir = $2, alamat = $3, jenis_kelamin = $4
	WHERE nim = $5 
	`
	err := r.db.QueryRow(
		sqlStatement,
		mahasiswa.NamaMahasiswa,
		mahasiswa.TanggalLahir,
		mahasiswa.Alamat,
		mahasiswa.JenisKelamin,
		mahasiswa.Nim).Scan(
		&mahasiswa.Nim,
		&mahasiswa.NamaMahasiswa,
		&mahasiswa.TanggalLahir,
		&mahasiswa.Alamat,
		&mahasiswa.JenisKelamin)
	if err != nil {
		return mahasiswa, err
	}
	return mahasiswa, nil
}


func (r *mahasiswaRepository) DeleteMahasiswa(mahasiswa entity.Mahasiswa)error  {
	sqlStatement := "DELETE FROM Mahasiswa WHERE nip =$1"

	err := r.db.QueryRow(sqlStatement, mahasiswa.Nim)

	return err.Err()
}