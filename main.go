package main

import (
	"multec-api-sop/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1/datasop")
	{
		Datasop := new(controllers.DataSopRsController)
		//Photo := new(controllers.PhotoController)
		//MEMBUAT DATA
		v1.POST("/create_sop", Datasop.Create)
		//MEMBUAT DATA
		v1.POST("/create_detail/:id", Datasop.TambahDetail)
		//MEMBUAT DATA
		v1.PUT("/update_detail/:id/:id_det", Datasop.UpdateDetail)
		//DATA ALL
		v1.GET("/view_sop", Datasop.FindCoba)
		//DATA BY ID
		v1.GET("/view_sop/:id", Datasop.FindId)
		//UPDAT SOP
		v1.PUT("/update_sop/:id", Datasop.Update)
		//UPDAT GAMBAR
		v1.PUT("/update_image/:id/:id_det", Datasop.UpdateGambar)
		//DELETE
		v1.GET("/delete_sop/:id", Datasop.Delete)

		v1.GET("/delete_detail/:id/:id_det", Datasop.DeleteDetail)
		//UPLOAD GAMBAR + FORM INPUTAN
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"status": "Not Found"})
	})

	router.Run()
}
