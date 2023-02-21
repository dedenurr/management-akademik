-- +migrate Up
-- +migrate StatementBegin
CREATE DATABASE `university`

CREATE TABLE `dosen` (
  `nip` varchar(12) NOT NULL,
  `nama_dosen` varchar(25) NOT NULL,
  PRIMARY KEY (`nip`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

CREATE TABLE `mahasiswa` (
  `nim` varchar(9) NOT NULL,
  `nama_mahasiswa` varchar(25) NOT NULL,
  `tanggal_Lahir` date NOT NULL,
  `alamat` varchar(50) NOT NULL,
  `jenis_Kelamin` varchar(30) NOT NULL,
  PRIMARY KEY (`nim`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

CREATE TABLE `matakuliah` (
  `kode_mata_kuliah` varchar(7) NOT NULL,
  `nama_mata_kuliah` varchar(20) NOT NULL,
  `sks` int(2) NOT NULL,
  PRIMARY KEY (`kode_mata_kuliah`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

CREATE TABLE `perkuliahan` (
  `nim` varchar(9) DEFAULT NULL,
  `kode_mata_kuliah` varchar(7) DEFAULT NULL,
  `nip` varchar(12) DEFAULT NULL,
  `nilai` int NOT NULL,
  `grade` char(1) NOT NULL,
  KEY `nip` (`nip`),
  KEY `nim` (`nim`),
  KEY `kode_mata_kuliah` (`kode_mata_kuliah`),
  CONSTRAINT `fk_dosen_nip` FOREIGN KEY (`nip`) REFERENCES `dosen` (`nip`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_mahasiswa_nim` FOREIGN KEY (`nim`) REFERENCES `mahasiswa` (`nim`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_matakuliah_kode_mata_kuliah` FOREIGN KEY (`kode_mata_kuliah`) REFERENCES `matakuliah` (`kode_mata_kuliah`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8

-- +migrate StatementEnd