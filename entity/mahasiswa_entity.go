package entity

type Mahasiswa struct {
	Nim           string `json:"nim"`
	NamaMahasiswa string `json:"nama_mahasiswa"`
	TanggalLahir  string `json:"tanggal_lahir"`
	Alamat        string `json:"alamat"`
	JenisKelamin  string `json:"jenis_kelamin"`
}