package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Anprazt/TugasWebAPIGolang/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

type MatkulInput struct {
	ID             int    `json:"id" binding:"required"`
	KodeMataKuliah string `json:"kodematakuliah" binding:"required"`
	NamaMataKuliah string `json:"namamatakuliah" binding:"required,min=3"`
	JumlahSks      int    `json:"jumlahsks" binding:"required"`
	DosenPengampu  string `json:"dosenpengampu" binding:"required"`
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
	err := c.ShouldBindJSON(&dataInput)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			switch e.Tag() {
			case "required":
				report := fmt.Sprintf("%s is required", e.Field())
				errorMessages = append(errorMessages, report)
			case "min":
				report := fmt.Sprintf("%s must be more than 5 characters", e.Field())
				errorMessages = append(errorMessages, report)
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
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
	err := c.ShouldBindJSON(&dataInput)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			switch e.Tag() {
			case "required":
				report := fmt.Sprintf("%s is required", e.Field())
				errorMessages = append(errorMessages, report)
			case "min":
				report := fmt.Sprintf("%s must be more than 5 characters", e.Field())
				errorMessages = append(errorMessages, report)
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
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
		"Data":    true,
		"Message": "Berhasil hapus data",
	})
}
