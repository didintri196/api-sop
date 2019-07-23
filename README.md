# Dokumentasi API SOP MULTEC
### 0. TABEL COLLECTION SOP
| Collection    | Tipe Data | Description  |
| ------------- |:-------------:| -------------|
| _id   | Unique	  	| `id` Unik Heading|
| namaperusahaan         | String      |   Nama Perusahaan yang membuat SOP |
| judulsop | String | Judul SOP yang dibuat|
| nodokumen | String | data `*nodokumen`|
|norevisi | String | nomor revisi|
|halaman | String | jumlah halaman|
|prosedurtetap | String | Prosedur Tetap|
|tanggalterbit|String| tanggal terbit SOP| 
|divisi|String| data `*divisi`|
|published|bool| Status True or False|
|detaildokumen | String | data `[]detaildoumen` |

### Collection `*Nodokumen`
| Collection    | Tipe Data | Description  |
| ------------- |:-------------:| -------------|
| kodeperusahaan | String | Kode 2 Digit Nama Perusahaan|
|kodespo | String | 2 Digit nama SOP |
|kodeunit | String | 2 Digit kode unit|
|nourutspo| String | nomor urut dengan nomor SOP yang sebelumnya|

### Collection `*divisi`
| Collection    | Tipe Data | Description  |
| ------------- |:-------------:| -------------|
| id | String | `*Id Divisi`|
|name | String | Nama Divisi |


### Collection `[]detaildoumen`
| Collection    | Tipe Data | Description  |
| ------------- |:-------------:| -------------|
| _id | String | `*Id Divisi`|
|urutan | String | Urutan detail dokumen |
|nama | String | Nama detail dokumen |
|prosedur | String | Prosedur detail dokumen |
|gambar | String | Nama gambar detail dokumen |


### 1.  CREATE HEADING DATA

#### HTTP Request
```json
POST localhost:8080/v1/datasop/create_sop
```

#### Post Form Data

| Name    |               | Description  |
| ------------- |:-------------:| -------------|
| namaperusahaan         | required      |   Diisi dengan nama perusahaan yang membuat SOP |
| judulsop | required | Diisi dengan Judul SOP yang dibuat|
| nodokumen | required | isi sesuai kategori|
| kodeperusahaan | required | Diisi dengan Kode 2 Digit Nama Perusahaan|
|kodespo | required | Diisi dengan kode 2 Digit nama SOP |
|kodeunit | required | Diisi dengan 2 Digit kode unit|
|nourutspo| required | Diisi dengan nomor urut dengan nomor SOP yang sebelumnya|
|norevisi | required | Diisi dengan nomor revisi|
|halaman | required | Diisi dengan jumlah halaman|
|prosedurtetap | required | Diisi dengan Prosedur Tetap|
|tanggalterbit|Date| Diisi dengan tanggal terbit SOP| 
|id_divisi|required|Diisi dengan id divisi SOP|
|nama_divisi|required|Diisi dengan nama divisi SOP|
|nama[]|required|Diisi dengan nama detail prosedur SOP|
|prosedur[]|required|Diisi dengan isi detail prosedur SOP|
|upload[]|required|Diisi dengan gambar detail prosedur SOP|

`note : tanda array berarti bisa ditambahkan form tak terbatas`
#### Result :
```json
{
    "message": "Heading Created"
}
```
|  HTTP  | Message    |  Description  |
| :------------- | ------------- |:--------------|
|200|Heading Created| Data Berhasil Di buat di database|
|406| Invalid form |Data ada yang kosong|
|406 | Heading could not be Create | Data Gagal Di buat di database


### 2.  GET HEADING DATA

#### - GET ALL DATA :
#### HTTP Request
```json
GET localhost:8080/v1/datasop/view_sop?show=all
```

#### Result :
```json
{
    "current_page": 1,
    "data": [
        {
            "_id": "5c124a89b56c03e00862c577",
            "namaperusahaan": "RS MAJU JAYA2",
            "judulsop": "CARA MERAWAT PASIEN",
            "nodokumen": {
                "kodeperusahaan": "RS",
                "kodespo": "SP",
                "kodeunit": "UN",
                "nourutspo": "01"
            },
            "norevisi": "1",
            "halaman": "1",
            "prosedurtetap": "iya",
            "tanggalterbit": "05/12/2016",
            "detaildokumen": [
                {
                    "_id": "DOK1544704166744649065",
                    "urutan": 3,
                    "nama": "asdasdasd",
                    "prosedur": "asdasdasd",
                    "gambar": "1544704166744.png"
                }
            ],
            "divisi": {
                "id": "00001",
                "name": "IT"
            },
            "published": true
        }
    ],
    "from": 1,
    "last_page": 1,
    "next_page": 2,
    "per_page": 5,
    "prev_page": 1,
    "status": "ok",
    "to": 5,
    "total": 1
}
```
#### PARAMETER YANG TERSEDIA
| Param | Value   |  Description  |
| :------------- | ------------- |:--------------|
|filter |`string`| untuk pencarian berdasarkan judul |
|sort |`_id`| Menentukan acuan urutan data |
|page |`int`| Di isi dengan urutan halaman |
|per_page |`int`| Menentukan jumlah data yang ditampilkan |
|show | `all`| Menentukan Apakaah data ingin di tampilkan semua  |
|divisi | `id_divisi`| Menampilkan data dengan acuan divisi  |


#### - GET BY ID DATA :
#### HTTP Request
```json
GET localhost:8080/v1/datasop/view_sop/:id
```

#### Result :
```json
{
    "data": {
        "_id": "5c124a89b56c03e00862c577",
        "namaperusahaan": "RS MAJU JAYA2",
        "judulsop": "CARA MERAWAT PASIEN",
        "nodokumen": {
            "kodeperusahaan": "RS",
            "kodespo": "SP",
            "kodeunit": "UN",
            "nourutspo": "01"
        },
        "norevisi": "1",
        "halaman": "1",
        "prosedurtetap": "iya",
        "tanggalterbit": "05/12/2016",
        "detaildokumen": [
            {
                "_id": "DOK1544704166744649065",
                "urutan": 3,
                "nama": "asdasdasd",
                "prosedur": "asdasdasd",
                "gambar": "1544704166744.png"
            }
        ],
        "divisi": {
            "id": "00001",
            "name": "IT"
        },
        "published": true
    }
}
```

|  HTTP  | Message    |  Description  |
| :------------- | ------------- |:--------------|
|200|| Data Berhasil Di buat di database|
|404| Find Error |Mengambil Data Database Gagal / Data tidak ada|

### 3.  UPDATE HEADING DATA

#### HTTP Request
```json
PUT localhost:8080/v1/datasop/update_sop/:id
```
| Parameter    |               | Description  |
| ------------- |:-------------:| -------------|
| :id         | required      | Diisi `id` unik dari DB

#### Post Form Data

| Name    |               | Description  |
| ------------- |:-------------:| -------------|
| namaperusahaan         | required      |   Diisi dengan nama perusahaan yang membuat SOP |
| judulsop | required | Diisi dengan Judul SOP yang dibuat|
| nodokumen | required | isi sesuai kategori|
| kodeperusahaan | required | Diisi dengan Kode 2 Digit Nama Perusahaan|
|kodespo | required | Diisi dengan kode 2 Digit nama SOP |
|kodeunit | required | Diisi dengan 2 Digit kode unit|
|nourutspo| required | Diisi dengan nomor urut dengan nomor SOP yang sebelumnya|
|norevisi | required | Diisi dengan nomor revisi|
|halaman | required | Diisi dengan jumlah halaman|
|prosedurtetap | required | Diisi dengan Prosedur Tetap|
|tanggalterbit|Date| Diisi dengan tanggal terbit SOP| 
|id_divisi|required|Diisi dengan id divisi SOP|
|nama_divisi|required|Diisi dengan nama divisi SOP|
|_id[]|required|Diisi dengan id detail prosedur SOP|
|nama[]|required|Diisi dengan nama detail prosedur SOP|
|prosedur[]|required|Diisi dengan isi detail prosedur SOP|
|upload[]|required|Diisi dengan gambar detail prosedur SOP|

`note : tanda array berarti bisa ditambahkan form tak terbatas`

#### Result :
```json
{
    "message": "Heading Updated"
}
```
|  HTTP  | Message    |  Description  |
| :------------- | ------------- |:--------------|
|200|Heading Updated| Data Berhasil Di Update|
|406| Invalid form |Data ada yang kosong|
|406 | Heading could not be Update | Data Gagal Di Update di database


### 4.  DELETE HEADING DATA

#### HTTP Request
```json
GET localhost:8080/v1/datasop/delete_sop/:id
```

#### Result :
```json
{
    "message": "Heading Deleted"
}
```

|  HTTP  | Message    |  Description  |
| :------------- | ------------- |:--------------|
|200|Heading Deleted| Data Berhasil Di Hapus|
|406| Heading Could not be Deleted | Data Gagal Di Hapus di database

### 5.  CREATE DETAIL PROSEDUR SOP DATA

#### HTTP Request
```json
POST localhost:8080/v1/datasop/create_detail/:id_sop
```

#### Post Form Data

| Name    |               | Description  |
| ------------- |:-------------:| -------------|
|urutan |required|Diisi dengan urutan detail prosedur SOP|
|nama|required|Diisi dengan nama detail prosedur SOP|
|prosedur|required|Diisi dengan isi detail prosedur SOP|
|upload|required|Diisi dengan gambar detail prosedur SOP|

#### Result :
```json
{
    "message": "Detail Sucessful Added"
}
```
|  HTTP  | Message    |  Description  |
| :------------- | ------------- |:--------------|
|200|Detail Sucessful Added| Data Berhasil Di buat di database|
|406| Invalid form |Data ada yang kosong|
|406 | Detail Could Not Be Added | Data Gagal Di buat di database


### 6.  UPDATE DETAIL SOP DATA NON IMAGE

#### HTTP Request
```json
POST localhost:8080/v1/datasop/update_detail/:id/:id_det
```

#### Post Form Data

| Name    |               | Description  |
| ------------- |:-------------:| -------------|
|urutan |required|Diisi dengan urutan detail prosedur SOP|
|nama|required|Diisi dengan nama detail prosedur SOP|
|prosedur|required|Diisi dengan isi detail prosedur SOP|
|upload|required|Diisi dengan gambar detail prosedur SOP|

#### Result :
```json
{
    "message": "Detail Sucessful Update"
}
```
|  HTTP  | Message    |  Description  |
| :------------- | ------------- |:--------------|
|200|Detail Sucessful Update| Data Berhasil Di buat di database|
|406| Invalid form |Data ada yang kosong|
|406 | Detail Could Not Be Update | Data Gagal Di buat di database

### 7.  UPDATE IMAGE DETAIL SOP DATA NON IMAGE

#### HTTP Request
```json
POST localhost:8080/v1/datasop/update_image/:id/:id_det
```

#### Post Form Data

| Name    |               | Description  |
| ------------- |:-------------:| -------------|
|upload|required|Diisi dengan gambar detail prosedur SOP|

#### Result :
```json
{
    "message": "Image Sucessful Update"
}
```
|  HTTP  | Message    |  Description  |
| :------------- | ------------- |:--------------|
|200|Image Sucessful Update| Data Berhasil Di buat di database|
|406| Invalid form |Data ada yang kosong|
|406 | Image Could Not Be Update | Data Gagal Di buat di database

### 8.  DELETE HEADING DATA

#### HTTP Request
```json
GET localhost:8080/v1/datasop/delete_detail/:id/:id_det
```

#### Result :
```json
{
    "message": "Detail Deleted""
}
```

|  HTTP  | Message    |  Description  |
| :------------- | ------------- |:--------------|
|200|Detail Deleted| Data Berhasil Di Hapus|
|406| Detail Could not be Deleted | Data Gagal Di Hapus di database
