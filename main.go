package main

import (
	"net/http"

	"github.com/Anprazt/TugasWebAPIGolang/controllers"
	"github.com/Anprazt/TugasWebAPIGolang/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//Models
	db := models.SetUpModels()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Tugas Web API golang",
		})
	})
	//MAHASISWA
	//Get all data
	r.GET("/mahasiswa", controllers.ReadDataMhs)
	//Post data
	r.POST("/mahasiswa", controllers.CreateDataMhs)
	//Update data
	r.PUT("/mahasiswa/:nim", controllers.UpdateDataMhs)
	//Delete data
	r.DELETE("/mahasiswa/:nim", controllers.DeleteDataMhs)

	//MATAKULIAH
	//Get all data
	r.GET("/matakuliah", controllers.ReadDataMatkul)
	//Post data
	r.POST("/matakuliah", controllers.CreateDataMatkul)
	//Update data
	r.PUT("/matakuliah/:kode", controllers.UpdateDataMatkul)
	//Delete data
	r.DELETE("/matakuliah/:kode", controllers.DeleteDataMatkul)

	r.Run()
}
