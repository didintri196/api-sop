package models

import (
	"multec-api-sop/db"
	"multec-api-sop/forms"
	"strconv"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

var DbName = "DBsimmultek"

type DataSopRs struct {
	ID             bson.ObjectId   `json:"_id" bson:"_id,omitempty"`
	NamaPerusahaan string          `json:"namaperusahaan" bson:"namaperusahaan"`
	JudulSop       string          `json:"judulsop" bson:"judulsop"`
	NoDokumen      *NoDokumen      `json:"nodokumen" bson:"nodokumen"`
	NoRevisi       string          `json:"norevisi" bson:"norevisi"`
	Halaman        string          `json:"halaman" bson:"halaman"`
	ProsedurTetap  string          `json:"prosedurtetap" bson:"prosedurtetap" `
	TanggalTerbit  string          `json:"tanggalterbit" bson:"tanggalterbit"`
	DetailDokumen  []DetailDokumen `json:"detaildokumen"  binding:"required"`
	Divisi         *Divisi         `json:"divisi" bson:"divisi"`
	Published      bool            `json:"published" bson:"published"`
}
type DetailDokumen struct {
	Id       string `json:"_id" bson:"_id,omitempty"`
	Urutan   int    `json:"urutan" binding:"required"`
	Nama     string `json:"nama" binding:"required"`
	Prosedur string `json:"prosedur" binding:"required"`
	Gambar   string `json:"gambar" binding:"required"`
}
type NoDokumen struct {
	KodePerusahaan string `json:"kodeperusahaan" binding:"required"`
	KodeSpo        string `json:"kodespo" binding:"required"`
	KodeUnit       string `json:"kodeunit" binding:"required"`
	NoUrutSpo      string `json:"nourutspo" binding:"required"`
}

type Divisi struct {
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type DataSopRsModel struct{}

//CreateModel
func (m *DataSopRsModel) Create(data forms.DataSopRsCommand, nodokumen forms.NoDokumen, divisi forms.Divisi) (err error) {
	db := db.DBConnection{}
	db.NewConnection()
	defer db.Close()
	collection := db.Session.DB(DbName).C("data_sop")
	err = collection.Insert(bson.M{
		"namaperusahaan": data.NamaPerusahaan,
		"judulsop":       data.JudulSop,
		"nodokumen":      nodokumen,
		"norevisi":       data.NoRevisi,
		"halaman":        data.Halaman,
		"prosedurtetap":  data.ProsedurTetap,
		"tanggalterbit":  data.TanggalTerbit,
		"detaildokumen":  data.DetailDokumen,
		"divisi":         divisi,
		"published":      data.Published,
	})
	return err
}

//Menampilkan Seluruh Data
func (m *DataSopRsModel) Find() (sop []DataSopRs, err error) {
	db := db.DBConnection{}
	db.NewConnection()
	defer db.Close()
	collection := db.Session.DB(DbName).C("data_sop")
	err = collection.Find(bson.M{}).All(&sop)
	return sop, err
}

//Menampilkan By ID
func (m *DataSopRsModel) FindId(id string) (sop DataSopRs, err error) {
	db := db.DBConnection{}
	db.NewConnection()
	defer db.Close()
	collection := db.Session.DB(DbName).C("data_sop")
	err = collection.FindId(bson.ObjectIdHex(id)).One(&sop)
	return sop, err
}

//Menampilkan By Judul REGEX
func (m *DataSopRsModel) FindJudulSop(judul string, kategori string) (sop []DataSopRs, err error) {
	db := db.DBConnection{}
	db.NewConnection()
	defer db.Close()
	collection := db.Session.DB(DbName).C("data_sop")
	regex_judul := bson.M{"$regex": bson.RegEx{Pattern: judul, Options: "i"}}
	if kategori != "" {
		err = collection.Find(bson.M{
			"$or": []interface{}{
				bson.M{"judulsop": regex_judul},
			},
			"idkategori": kategori,
		}).All(&sop)
	} else {
		err = collection.Find(bson.M{
			"$or": []interface{}{
				bson.M{"judulsop": regex_judul},
			},
		}).All(&sop)
	}

	return sop, err
}

//Update Data
func (m *DataSopRsModel) Update(id string, data forms.DataSopRsCommand, nodokumen forms.NoDokumen, divisi forms.Divisi) (err error) {
	db := db.DBConnection{}
	db.NewConnection()
	defer db.Close()
	collection := db.Session.DB(DbName).C("data_sop")
	err = collection.UpdateId(bson.ObjectIdHex(id), bson.M{
		"namaperusahaan": data.NamaPerusahaan,
		"judulsop":       data.JudulSop,
		"nodokumen":      nodokumen,
		"norevisi":       data.NoRevisi,
		"halaman":        data.Halaman,
		"prosedurtetap":  data.ProsedurTetap,
		"tanggalterbit":  data.TanggalTerbit,
		"detaildokumen":  data.DetailDokumen,
		"divisi":         divisi,
		"published":      data.Published,
	})

	return err
}

//Update Data
func (m *DataSopRsModel) UpdateGambar(id string, id_det string, gambar string) (err error) {
	db := db.DBConnection{}
	db.NewConnection()
	defer db.Close()
	collection := db.Session.DB(DbName).C("data_sop")
	query := bson.M{
		"_id":               bson.ObjectIdHex(id),
		"detaildokumen._id": id_det,
	}
	err = collection.Update(query, bson.M{
		"$set": bson.M{
			"detaildokumen.$.gambar": gambar,
		},
	})

	return err
}

func (m *DataSopRsModel) Delete(id string) (err error) {
	db := db.DBConnection{}
	db.NewConnection()
	defer db.Close()
	collection := db.Session.DB(DbName).C("data_sop")
	err = collection.RemoveId(bson.ObjectIdHex(id))
	return err
}

func (m *DataSopRsModel) FindCoba(filter string, sort string, pageNo string, perPage string, show string, divisi string) (data []DataSopRs, err error, count int, er error) {
	sorting := sort
	if strings.Contains(sort, "asc") {
		sorting = strings.Replace(sort, "|asc", "", -1)
	} else if strings.Contains(sort, "desc") {
		sorting = strings.Replace(sort, "|desc", "", -1)
		sorting = "-" + sorting
	}

	db := db.DBConnection{}
	db.NewConnection()
	defer db.Close()
	collection := db.Session.DB(DbName).C("data_sop")
	pn, _ := strconv.Atoi(pageNo)
	pp, _ := strconv.Atoi(perPage)
	regex := bson.M{"$regex": bson.RegEx{Pattern: filter, Options: "i"}}
	if divisi != "" {
		if show == "all" {
			query := collection.Find(bson.M{
				"divisi": bson.M{
					"_id": divisi,
				},
				"$or": []interface{}{
					bson.M{"judulsop": regex},
				},
			})
			err = query.All(&data)
			count, er = query.Count()
		} else {
			query := collection.Find(bson.M{
				"divisi": bson.M{
					"_id": divisi,
				},
				"$or": []interface{}{
					bson.M{"judulsop": regex},
				},
			})
			count, er = query.Count()
			err = query.Sort(sorting).Skip((pn - 1) * pp).Limit(pp).All(&data)
		}
	} else {
		if show == "all" {
			query := collection.Find(bson.M{
				"published": true,
				"$or": []interface{}{
					bson.M{"judulsop": regex},
				},
			})
			err = query.All(&data)
			count, er = query.Count()
		} else {
			query := collection.Find(bson.M{
				"published": true,
				"$or": []interface{}{
					bson.M{"judulsop": regex},
				},
			})
			count, er = query.Count()
			err = query.Sort(sorting).Skip((pn - 1) * pp).Limit(pp).All(&data)

		}

	}

	return data, err, count, er
}

func (m *DataSopRsModel) TambahDetail(id string, data forms.DetailDokumen) (err error) {
	db := db.DBConnection{}
	db.NewConnection()
	defer db.Close()
	collection := db.Session.DB(DbName).C("data_sop")
	err = collection.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{
		"$addToSet": bson.M{
			"detaildokumen": bson.M{
				"_id":      data.Id,
				"urutan":   data.Urutan,
				"nama":     data.Nama,
				"prosedur": data.Prosedur,
				"gambar":   data.Gambar,
			},
		},
	})
	return err
}

func (m *DataSopRsModel) UpdateDetail(id string, id_det string, data forms.DetailDokumen) (err error) {
	db := db.DBConnection{}
	db.NewConnection()
	defer db.Close()
	collection := db.Session.DB(DbName).C("data_sop")
	query := bson.M{
		"_id":               bson.ObjectIdHex(id),
		"detaildokumen._id": id_det,
	}
	err = collection.Update(query, bson.M{
		"$set": bson.M{
			"detaildokumen.$.urutan":   data.Urutan,
			"detaildokumen.$.nama":     data.Nama,
			"detaildokumen.$.prosedur": data.Prosedur,
		},
	})

	return err
}

func (m *DataSopRsModel) DeleteDetail(id string, id_det string) (err error) {
	db := db.DBConnection{}
	db.NewConnection()
	defer db.Close()
	collection := db.Session.DB(DbName).C("data_sop")
	query := bson.M{
		"_id":               bson.ObjectIdHex(id),
		"detaildokumen._id": id_det,
	}
	update := bson.M{
		"$pull": bson.M{
			"detaildokumen": bson.M{
				"_id": id_det},
		},
	}
	err = collection.Update(query, update)

	return err
}
