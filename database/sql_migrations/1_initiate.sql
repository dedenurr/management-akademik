-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE dosen (
    nip VARCHAR(12) NOT NULL,
    nama_dosen VARCHAR(25) NOT NULL,
    PRIMARY KEY (nip)
);

CREATE TABLE mahasiswa (
    nim VARCHAR(9) NOT NULL,
    nama_mahasiswa VARCHAR(25) NOT NULL,
    tanggal_lahir date NOT NULL,
    alamat VARCHAR(50) NOT NULL,
    jenis_kelamin VARCHAR(25) NOT NULL,
    PRIMARY KEY (nim)
);

CREATE TABLE matakuliah (
    kode_mata_kuliah VARCHAR(7) NOT NULL,
    nama_mata_kuliah VARCHAR(20) NOT NULL,
    sks INT NOT NULL,
    PRIMARY KEY (kode_mata_kuliah)
);

CREATE TABLE perkuliahan (
    nim VARCHAR(9) DEFAULT NULL,
    kode_mata_kuliah VARCHAR(7) DEFAULT NULL,
    nip VARCHAR(12) DEFAULT NULL,
    nilai INT NOT NULL,
    grade VARCHAR(12),
    CONSTRAINT fk_dosen_nip FOREIGN KEY (nip) REFERENCES dosen(nip) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_mahasiswa_nim FOREIGN KEY (nim) REFERENCES mahasiswa(nim) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_matakuliah_kode_mata_kuliah FOREIGN KEY (kode_mata_kuliah) REFERENCES matakuliah(kode_mata_kuliah) ON DELETE CASCADE ON UPDATE CASCADE
);

-- +migrate StatementEnd