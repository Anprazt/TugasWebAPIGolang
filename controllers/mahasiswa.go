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

type MahasiswaInput struct {
	ID            int    `json:"id" binding:"required,uuid" gorm:"primary_key"`
	Nama          string `json:"nama" binding:"required"`
	Email         string `json:"email" binding:"required,email"`
	Prodi         string `json:"prodi" binding:"required"`
	Fakultas      string `json:"fakultas" binding:"required"`
	NIM           int    `json:"nim" binding:"required"`
	TahunAngkatan int    `json:"tahunangkatan" binding:"required"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

//GET DATA
func ReadDataMhs(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var mahasiswa []models.Mahasiswa
	db.Find(&mahasiswa)
	c.JSON(http.StatusOK, gin.H{
		"Data": mahasiswa,
		"Time": time.Now(),
	})
}

//POST DATA
func CreateDataMhs(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	//validasi
	var dataInput MahasiswaInput
	err := c.ShouldBindJSON(&dataInput)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			switch e.Tag() {
			case "required":
				report := fmt.Sprintf("%s is required", e.Field())
				errorMessages = append(errorMessages, report)
			case "email":
				report := fmt.Sprintf("%s is not valid email", e.Field())
				errorMessages = append(errorMessages, report)
			case "minsize":
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
	mahasiswa := models.Mahasiswa{
		ID:            dataInput.ID,
		Nama:          dataInput.Nama,
		Email:         dataInput.Email,
		Prodi:         dataInput.Prodi,
		Fakultas:      dataInput.Fakultas,
		NIM:           dataInput.NIM,
		TahunAngkatan: dataInput.TahunAngkatan,
	}

	db.Create(&mahasiswa)
	//menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Message": "Berhasil input data",
		"Data":    mahasiswa,
		"Time":    time.Now(),
	})
}

//UPDATE DATA
func UpdateDataMhs(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	//cek data
	var mahasiswa models.Mahasiswa
	if err := db.Where("nim = ?", c.Param("nim")).First(&mahasiswa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Data mahasiswa tidak ditemukan",
		})
		return
	}
	//validasi
	var dataInput MahasiswaInput
	err := c.ShouldBindJSON(&dataInput)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			switch e.Tag() {
			case "required":
				report := fmt.Sprintf("%s is required", e.Field())
				errorMessages = append(errorMessages, report)
			case "email":
				report := fmt.Sprintf("%s is not valid email", e.Field())
				errorMessages = append(errorMessages, report)
			case "minsize":
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
	db.Model(&mahasiswa).Update(dataInput)

	//menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Message": "Berhasil ubah data",
		"Data":    mahasiswa,
		"Time":    time.Now(),
	})
}

//DELETE DATA
func DeleteDataMhs(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek data
	var mahasiswa models.Mahasiswa
	if err := db.Where("nim = ?", c.Param("nim")).First(&mahasiswa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Data mahasiswa tidak ditemukan",
		})
		return
	}
	//proses hapus data
	db.Delete(&mahasiswa)

	//menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Data": true,
	})
}
