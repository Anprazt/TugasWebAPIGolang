package controllers

import (
	"net/http"
	"time"

	"github.com/Anprazt/TugasWebAPIGolang/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type MatkulInput struct {
	ID             int    `json:"id"`
	KodeMataKuliah string `json:"kodematakuliah"`
	NamaMataKuliah string `json:"namamatakuliah"`
	JumlahSks      int    `json:"jumlahsks"`
	DosenPengampu  string `json:"dosenpengampu"`
}

//GET DATA
func ReadDataMatkul(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var matkul []models.MataKuliah
	db.Find(&matkul)
	c.JSON(http.StatusOK, gin.H{
		"Data": matkul,
		"Time": time.Now(),
	})
}

//POST DATA
func CreateDataMatkul(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	//validasi
	var dataInput MatkulInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	//input data
	matkul := models.MataKuliah{
		ID:             dataInput.ID,
		KodeMataKuliah: dataInput.KodeMataKuliah,
		NamaMataKuliah: dataInput.NamaMataKuliah,
		JumlahSks:      dataInput.JumlahSks,
		DosenPengampu:  dataInput.DosenPengampu,
	}

	db.Create(&matkul)
	//menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Message": "Berhasil input data",
		"Data":    matkul,
		"Time":    time.Now(),
	})
}

//UPDATE DATA
func UpdateDataMatkul(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	//cek data
	var matkul models.MataKuliah
	if err := db.Where("kodematakuliah = ?", c.Param("kodematakuliah")).First(&matkul).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Data matakuliah tidak ditemukan",
		})
		return
	}
	//validasi
	var dataInput MatkulInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	//proses ubah data
	db.Model(&matkul).Update(dataInput)

	//menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Message": "Berhasil ubah data",
		"Data":    matkul,
		"Time":    time.Now(),
	})
}

//DELETE DATA
func DeleteDataMatkul(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek data
	var matkul models.MataKuliah
	if err := db.Where("kodematakuliah = ?", c.Param("kodematakuliah")).First(&matkul).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Data matakuliah tidak ditemukan",
		})
		return
	}
	//proses hapus data
	db.Delete(&matkul)

	//menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Data": true,
	})
}
