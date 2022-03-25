package models

type Mahasiswa struct {
	ID            int    `json:"id" binding:"required,uuid" gorm:"primary_key"`
	Nama          string `json:"nama" binding:"required"`
	Email         string `json:"email" binding:"required,email"`
	Prodi         string `json:"prodi" binding:"required"`
	Fakultas      string `json:"fakultas" binding:"required"`
	NIM           int    `json:"nim" binding:"required"`
	TahunAngkatan int    `json:"tahunangkatan" binding:"required"`
}
