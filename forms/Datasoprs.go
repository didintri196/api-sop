package forms

type DataSopRsCommand struct {
	NamaPerusahaan string          `json:"namaperusahaan"  binding:"required"`
	JudulSop       string          `json:"judulsop"  binding:"required"`
	NoDokumen      *NoDokumen      `json:"nodokumen"  binding:"required"`
	NoRevisi       string          `json:"norevisi"  binding:"required"`
	Halaman        string          `json:"halaman"  binding:"required"`
	ProsedurTetap  string          `json:"prosedurtetap"  binding:"required"`
	TanggalTerbit  string          `json:"tanggalterbit"  binding:"required"`
	DetailDokumen  []DetailDokumen `json:"detaildokumen"  binding:"required"`
	Divisi         *Divisi         `json:"divisi"  binding:"required"`
	Published      bool            `json:"published" binding:"required"`
}

type NoDokumen struct {
	KodePerusahaan string `json:"kodeperusahaan" binding:"required"`
	KodeSpo        string `json:"kodespo" binding:"required"`
	KodeUnit       string `json:"kodeunit" binding:"required"`
	NoUrutSpo      string `json:"nourutspo" binding:"required"`
}

type DetailDokumen struct {
	Id       string `json:"_id" bson:"_id,omitempty"`
	Urutan   int    `json:"urutan" binding:"required"`
	Nama     string `json:"nama" binding:"required"`
	Prosedur string `json:"prosedur" binding:"required"`
	Gambar   string `json:"gambar" binding:"required"`
}
type UpdateDetailDokumen struct {
	Id       string `json:"_id" bson:"_id,omitempty"`
	Urutan   int    `json:"urutan" binding:"required"`
	Nama     string `json:"nama" binding:"required"`
	Prosedur string `json:"prosedur" binding:"required"`
	Gambar   string `json:"gambar"`
}

type UpdateDataSopRsCommand struct {
	NamaPerusahaan string                `json:"namaperusahaan"  binding:"required"`
	JudulSop       string                `json:"judulsop"  binding:"required"`
	NoDokumen      *NoDokumen            `json:"nodokumen"  binding:"required"`
	NoRevisi       string                `json:"norevisi"  binding:"required"`
	Halaman        string                `json:"halaman"  binding:"required"`
	ProsedurTetap  string                `json:"prosedurtetap"  binding:"required"`
	TanggalTerbit  string                `json:"tanggalterbit"  binding:"required"`
	DetailDokumen  []UpdateDetailDokumen `json:"detaildokumen"  binding:"required"`
	Divisi         *Divisi               `json:"divisi"  binding:"required"`
	Published      bool                  `json:"published" binding:"required"`
}

type Divisi struct {
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
