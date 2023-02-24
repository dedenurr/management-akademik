# final-project-management-akademik
API untuk Manajemen Akademik yang berfungsi untuk memonitoring nilai akademik mahasiswa  dalam perkuliahan pada sebuah universitas. Untuk membangun API ini teridiri dari beberapa entitas diantaranya:
    1. mahasiswa
    2. dosen
    3. matakuliah
    4. perkuliahan

## Resource
* Railway : https://management-akademik-production.up.railway.app/
* Postman : https://documenter.getpostman.com/view/25656509/2s93CNMtFD
* Slides  : https://drive.google.com/drive/u/4/folders/1Obl_Gt3MrK60SvzUdeWp-oTIdRUyfQXP

### Build With
    1. Go versi 1.19
    2. Framework Go Gin
    3. Postgres

### Get started
Download terlebih dahulu library-library yang akan digunakan:
    
    go get -u "github.com/gin-gonic/gin"
    go get -u "github.com/lib/pq"
    go get -u "github.com/rubenv/sql-migrate"
    go get -u "github.com/gobuffalo/packr/v2"
    go get -u "github.com/joho/godotenv"
    go get -u "github.com/appleboy/gin-jwt/"

### Authorization
    
    JWT Token


### Ada 4 jenis API  :
| Path               | Deskripsi                                                                               |
|--------------------|-----------------------------------------------------------------------------------------|
| {{URL}} / bo / ... | API untuk back office, hanya boleh diakses oleh user dengan role admin (isAdmin = true) | 
| {{URL}} / ...      | API yang berhubungan langsung dengan user/customer (isAdmin = false)                    |


### APIs
#### URL
```
https://final-project-ticketing-app-production.up.railway.app
```

