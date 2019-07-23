package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"multec-api-sop/forms"
	"multec-api-sop/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var DataSopRsModel = new(models.DataSopRsModel)

//var MailLibrary = new(library.MailLibrary)

type DataSopRsController struct{}

//FUNGSI CREATE
func (heading *DataSopRsController) Create(c *gin.Context) {
	nodokumen := forms.NoDokumen{
		KodePerusahaan: c.PostForm("kodeperusahaan"),
		KodeSpo:        c.PostForm("kodespo"),
		KodeUnit:       c.PostForm("kodeunit"),
		NoUrutSpo:      c.PostForm("nourutspo"),
	}
	divisi := forms.Divisi{
		Id:   c.PostForm("id_divisi"),
		Name: c.PostForm("nama_divisi"),
	}

	var jsonText = []byte(`[]`)
	var idents []forms.DetailDokumen
	if err := json.Unmarshal([]byte(jsonText), &idents); err != nil {
		log.Println(err)
	}
	//KODE PARSING ARRAY
	Nama := c.PostFormArray("nama[]")
	Prosedur := c.PostFormArray("prosedur[]")
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	//fmt.Println(Nama)
	var i = 0
	if len(Nama) == len(Prosedur) && len(Nama) == len(files) {
		for _, element := range Nama {
			codeEvnt := "DOK" + strconv.Itoa(int(time.Now().UnixNano())+rand.Intn(9999))

			//fmt.Println("Uploading Gambar")
			path := "./assets/foto/prosedur/"
			time := time.Now().UnixNano() / 1000000
			filename := strconv.Itoa(i) + strconv.Itoa(int(time)) + ".png"
			if err := c.SaveUploadedFile(files[i], path+"/"+filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
				return
			} else {
				idents = append(idents, forms.DetailDokumen{Id: codeEvnt, Urutan: i + 1, Nama: element, Prosedur: Prosedur[i], Gambar: filename})
				fmt.Println("Urutan : ", i+1, " Nama : ", element, " Prosedur : ", Prosedur[i], " Gambar: ", filename)
			}
			i++
		}
	} else {
		c.JSON(406, gin.H{"message": "Data Or Image Cannot Be Null"})
		c.Abort()
		return
	}
	datasop := forms.DataSopRsCommand{
		NamaPerusahaan: c.PostForm("namaperusahaan"),
		JudulSop:       c.PostForm("judulsop"),
		NoRevisi:       c.PostForm("norevisi"),
		Halaman:        c.PostForm("halaman"),
		ProsedurTetap:  c.PostForm("prosedurtetap"),
		TanggalTerbit:  c.PostForm("tanggalterbit"),
		Published:      true,
		DetailDokumen:  idents,
	}

	err := DataSopRsModel.Create(datasop, nodokumen, divisi)
	if err != nil {
		c.JSON(406, gin.H{"message": "Heading could not be Create", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"message": "Heading Created"})

}

//FUNGSI MENAMPILKAN DATA
func (heading *DataSopRsController) Find(c *gin.Context) {
	list, err := DataSopRsModel.Find()
	if err != nil {
		c.JSON(404, gin.H{"message": "Find Error", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": list})
	}
}

//FUNGSI MENAMPILKAN DATA BY ID
func (heading *DataSopRsController) FindId(c *gin.Context) {
	id := c.Param("id")
	list, err := DataSopRsModel.FindId(id)
	if err != nil {
		c.JSON(404, gin.H{"message": "Find ID Error", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": list})
	}
}

//FUNGSI CARI DATA BY JUDUL
func (heading *DataSopRsController) FindJudulSop(c *gin.Context) {
	judul := c.Query("judul")
	kategori := c.Query("kategori")
	list, err := DataSopRsModel.FindJudulSop(judul, kategori)
	if err != nil {
		c.JSON(404, gin.H{"message": "Find SOP Error", "error": err.Error()})
		c.Abort()
	} else {
		query := gin.H{"judul": judul, "kategori": kategori}
		c.JSON(200, gin.H{"data": list, "query": query})
	}
}

//FUNGSI UPDATE DATA
//Update
func (heading *DataSopRsController) Update(c *gin.Context) {
	id := c.Param("id")
	nodokumen := forms.NoDokumen{
		KodePerusahaan: c.PostForm("kodeperusahaan"),
		KodeSpo:        c.PostForm("kodespo"),
		KodeUnit:       c.PostForm("kodeunit"),
		NoUrutSpo:      c.PostForm("nourutspo"),
	}
	divisi := forms.Divisi{
		Id:   c.PostForm("id_divisi"),
		Name: c.PostForm("nama_divisi"),
	}

	var jsonText = []byte(`[]`)
	var idents []forms.DetailDokumen
	if err := json.Unmarshal([]byte(jsonText), &idents); err != nil {
		log.Println(err)
	}
	//KODE PARSING ARRAY
	Nama := c.PostFormArray("nama[]")
	Prosedur := c.PostFormArray("prosedur[]")
	Files := c.PostFormArray("upload[]")
	Id := c.PostFormArray("_id[]")
	//fmt.Println(Nama)
	var i = 0
	codeEvnt := ""
	if len(Nama) == len(Prosedur) && len(Nama) == len(Files) {
		for _, element := range Nama {
			if Id[i] == "null" {
				codeEvnt = "DOK" + strconv.Itoa(int(time.Now().UnixNano())+rand.Intn(9999))
			} else {
				codeEvnt = Id[i]
			}
			//codeEvnt := "DOK" + strconv.Itoa(int(time.Now().UnixNano())+rand.Intn(9999))

			idents = append(idents, forms.DetailDokumen{Id: codeEvnt, Urutan: i + 1, Nama: element, Prosedur: Prosedur[i], Gambar: Files[i]})
			//fmt.Println("Urutan : ", i+1, " Nama : ", element, " Prosedur : ", Prosedur[i], " Gambar: ", files[i].Filename)

			i++
		}
	} else {
		c.JSON(406, gin.H{"message": "Data Cannot Be Null"})
		c.Abort()
		return
	}
	datasop := forms.DataSopRsCommand{
		NamaPerusahaan: c.PostForm("namaperusahaan"),
		JudulSop:       c.PostForm("judulsop"),
		NoRevisi:       c.PostForm("norevisi"),
		Halaman:        c.PostForm("halaman"),
		ProsedurTetap:  c.PostForm("prosedurtetap"),
		TanggalTerbit:  c.PostForm("tanggalterbit"),
		Published:      true,
		DetailDokumen:  idents,
	}
	err := DataSopRsModel.Update(id, datasop, nodokumen, divisi)
	if err != nil {
		c.JSON(406, gin.H{"message": "Heading Could Not Be Updated", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"message": "Heading Updated"})
}

func (heading *DataSopRsController) UpdateGambar(c *gin.Context) {
	id := c.Param("id")
	id_det := c.Param("id_det")
	file, _ := c.FormFile("file")
	path := "./assets/foto/prosedur/"
	time := time.Now().UnixNano() / 1000000
	filename := strconv.Itoa(int(time)) + ".png"

	if err := c.SaveUploadedFile(file, path+"/"+filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	erro := DataSopRsModel.UpdateGambar(id, id_det, filename)
	if erro != nil {
		c.JSON(406, gin.H{"message": "Image Could Not Be Updated", "error": erro.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"message": "Image Sucessful Updated"})
}
func (heading *DataSopRsController) TambahDetail(c *gin.Context) {
	id := c.Param("id")
	urutan := c.PostForm("urutan")
	nama := c.PostForm("nama")
	prosedur := c.PostForm("prosedur")
	file, errimg := c.FormFile("file")
	codeEvnt := "DOK" + strconv.Itoa(int(time.Now().UnixNano())+rand.Intn(9999))
	if errimg != nil {
		c.JSON(404, gin.H{"message": "terjadi kesalahan", "error": errimg.Error()})
		c.Abort()
		return
	}
	path := "./assets/foto/prosedur/"
	time := time.Now().UnixNano() / 1000000
	filename := strconv.Itoa(int(time)) + ".png"

	if err := c.SaveUploadedFile(file, path+"/"+filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	x, _ := strconv.Atoi(urutan)
	data := forms.DetailDokumen{
		Id:       codeEvnt,
		Urutan:   x,
		Nama:     nama,
		Prosedur: prosedur,
		Gambar:   filename,
	}
	erro := DataSopRsModel.TambahDetail(id, data)
	if erro != nil {
		c.JSON(406, gin.H{"message": "Detail Could Not Be Added", "error": erro.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"message": "Detail Sucessful Added"})
}

func (heading *DataSopRsController) UpdateDetail(c *gin.Context) {
	id := c.Param("id")
	id_det := c.Param("id_det")
	urutan := c.PostForm("urutan")
	nama := c.PostForm("nama")
	prosedur := c.PostForm("prosedur")
	x, _ := strconv.Atoi(urutan)
	data := forms.DetailDokumen{
		Urutan:   x,
		Nama:     nama,
		Prosedur: prosedur,
	}
	erro := DataSopRsModel.UpdateDetail(id, id_det, data)
	if erro != nil {
		c.JSON(406, gin.H{"message": "Detail Could Not Be Update", "error": erro.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"message": "Detail Sucessful Update"})
}

//Delete
func (heading *DataSopRsController) DeleteDetail(c *gin.Context) {
	id := c.Param("id")
	id_set := c.Param("id_det")
	err := DataSopRsModel.DeleteDetail(id, id_set)
	if err != nil {
		c.JSON(406, gin.H{"message": "Detail Could not be Deleted", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"message": "Detail Deleted"})
}

//Delete
func (heading *DataSopRsController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := DataSopRsModel.Delete(id)
	if err != nil {
		c.JSON(406, gin.H{"message": "Heading Could not be Deleted"})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"message": "Heading Deleted"})
}

//FUNGSI MENAMPILKAN DATA COBA
func (heading *DataSopRsController) FindCoba(c *gin.Context) {
	filter := c.Query("filter")
	sort := c.Query("sort")
	pageNo := c.Query("page")
	perPage := c.Query("per_page")
	show := c.Query("show")
	divisi := c.Query("divisi")
	prev_page := 0
	next_page := 0
	//set default value

	if sort == "" {
		sort = "id"
	}
	if pageNo == "" {
		pageNo = "1"
	}
	if perPage == "" {
		perPage = "5"
	}
	pp, _ := strconv.Atoi(perPage)
	pn, _ := strconv.Atoi(pageNo)

	data, err, count, er := DataSopRsModel.FindCoba(filter, sort, pageNo, perPage, show, divisi)
	lastPage := float64(count) / float64(pp)
	if pp != 0 {
		if count%pp == 0 {
			lastPage = lastPage
		} else {
			lastPage = lastPage + 1
		}
	} else {
		lastPage = float64(count) / float64(5)
	}
	if pageNo == "1" {
		prev_page = 1
		next_page = 2
	} else {
		prev_page = pn - 1
		if pn == int(lastPage) {
			next_page = pn
		} else {
			next_page = pn + 1
		}
	}
	if err != nil && er != nil {
		c.JSON(404, gin.H{
			"message": "terjadi kesalahan",
			"error":   err.Error(),
		})
		c.Abort()
	} else {
		c.JSON(200, gin.H{
			"total":        count,
			"per_page":     pp,
			"current_page": pn,
			"last_page":    int(lastPage),
			"next_page":    next_page,
			"prev_page":    prev_page,
			"from":         ((pn * pp) - pp) + 1,
			"to":           pn * pp,
			"data":         data,
			"status":       "ok",
		})
		c.Abort()
	}
}
