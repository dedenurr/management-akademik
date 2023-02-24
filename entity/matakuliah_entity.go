package entity

type MataKuliah struct {
	KodeMatakuliah string `json:"kode_mata_kuliah"`
	NamaMatakuliah string `json:"nama_mata_kuliah"`
	Sks            int    `json:"sks"`
}