package models

type MataKuliah struct {
	ID             int    `json:"id" binding:"required,uuid"`
	KodeMataKuliah string `json:"kodematakuliah" binding:"required,uuid"`
	NamaMataKuliah string `json:"namamatakuliah" binding:"required,gte=4"`
	JumlahSks      int    `json:"jumlahsks" binding:"required"`
	DosenPengampu  string `json:"dosenpengampu" binding:"required"`
}
