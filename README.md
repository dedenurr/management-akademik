# final-project-management-akademik
API untuk Manajemen Akademik yang berfungsi untuk memonitoring nilai akademik mahasiswa  dalam perkuliahan pada sebuah universitas. Untuk membangun API ini teridiri dari beberapa entitas diantaranya:
### Entitas
    1. Go versi 1.19
    2. Framework Go Gin
    3. Postgres

## Resource
* Railway : https://management-akademik-production.up.railway.app/
* Postman : https://documenter.getpostman.com/view/16828940/2s93CNNYuT 
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
| Path                   | Deskripsi                                                                  |
|------------------------|----------------------------------------------------------------------------|
| {{URL}} / dosens       | API untuk Data Dosen                                                       | 
| {{URL}} / mahasiswas   | API untuk Data Mahasiswa                                                   |
| {{URL}} / matakuliahs  | API untuk Data MataKuliah                                                  |
| {{URL}} / perkuliahans | API untuk Nilai Perkuliahan                                                |

### APIs
#### URL
```
https://management-akademik-production.up.railway.app/
```

