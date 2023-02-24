package entity

type Perkuliahan struct {
	Nim            string `json:"nim"`
	KodeMatakuliah string `json:"kode_mata_kuliah"`
	Nip            string `json:"nip"`
	Nilai          int    `json:"nilai"`
	Grade          string `json:"grade"`
	Id             int    `json:"id"`
}
