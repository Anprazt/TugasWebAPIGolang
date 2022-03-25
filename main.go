package main

import (
	"net/http"

	"github.com/Anprazt/TugasWebAPIGolang/controllers"
	"github.com/Anprazt/TugasWebAPIGolang/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//versioning API
	v1 := r.Group("/v1")

	//Models
	db := models.SetUpModels()
	v1.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	v1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Tugas Web API golang",
		})
	})
	//MAHASISWA
	//Get all data
	v1.GET("/mahasiswa", controllers.ReadDataMhs)
	//Post data
	v1.POST("/mahasiswa", controllers.CreateDataMhs)
	//Update data
	v1.PUT("/mahasiswa/:nim", controllers.UpdateDataMhs)
	//Delete data
	v1.DELETE("/mahasiswa/:nim", controllers.DeleteDataMhs)

	//MATAKULIAH
	//Get all data
	v1.GET("/matakuliah", controllers.ReadDataMatkul)
	//Post data
	v1.POST("/matakuliah", controllers.CreateDataMatkul)
	//Update data
	v1.PUT("/matakuliah/:kode", controllers.UpdateDataMatkul)
	//Delete data
	v1.DELETE("/matakuliah/:kode", controllers.DeleteDataMatkul)

	r.Run()
}
