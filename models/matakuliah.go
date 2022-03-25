package models

type MataKuliah struct {
	ID             int    `json:"id" binding:"required"`
	KodeMataKuliah string `json:"kodematakuliah" binding:"required"`
	NamaMataKuliah string `json:"namamatakuliah" binding:"required,min=3"`
	JumlahSks      int    `json:"jumlahsks" binding:"required"`
	DosenPengampu  string `json:"dosenpengampu" binding:"required"`
}
