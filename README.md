# final-project-management-akademik
Ticketing app merupakan kumpulan API yang disusun dengan bahasa pemrograman Go yang bertujuan untuk memudahkan pembelian tiket event secara online. Customer dapat melihat event apa saja yang akan berlangsung berdasarkan category, melakukan transaksi pembelian tiket, melakukan top up saldo wallet, dan mendapatkan tiket yang disertai dengan qr code. 

Namun, sebelum customer dapat melakukan transaksi pembelian tiket, customer harus melakukan registrasi terlebih dahulu dengan memasukkan data diri seperti: nama lengkap, alamat, email, dll. Setelah berhasil, customer dapat memulai pengalaman bertransaksi dan melakukan login ke aplikasi dengan memasukkan email dan password yang sudah terdaftar. 

### Terdapat 2 jenis API :
| Path               | Deskripsi                                                                               |
|--------------------|-----------------------------------------------------------------------------------------|
| {{URL}} / bo / ... | API untuk back office, hanya boleh diakses oleh user dengan role admin (isAdmin = true) | 
| {{URL}} / ...      | API yang berhubungan langsung dengan user/customer (isAdmin = false)                    |


Di sisi back office, admin dapat melakukan CRUD category, event, ticket, dan melihat seluruh transaksi.

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

### APIs
#### URL
```
https://final-project-ticketing-app-production.up.railway.app
```
### REGISTRATION
```
  POST /registration
```
Saat melakukan registration, customer akan mengisi field-field seperti dibawah ini. Kemudian, saat selesai insert ke tabel customer, secara otomatis juga customer memiliki wallet (relasi ke tabel wallet) dengan balance 0, dan account_number berjumlah 8 digit angka.
##### Parameter
| Parameter    | Tipe Data | Deskripsi                                                                                |
|--------------|-----------|------------------------------------------------------------------------------------------|
| full_name    | string    | Nama harus dalam alfabet, special character tidak diperbolehkan                          | 
| birth_date   | string    | Format birth_date harus yyyy-MM-dd                                                       |
| address      | string    | Alamat customer                                                                          |
| phone_number | string    | Format harus diawali dengan 0 atau +62 atau 62 dengan jumlah maksimal 13 digit           |
| email        | string    | Format harus sesuai dengan email, contoh: author@email.com                               |
| password     | string    | Password customer, saat diinsert ke dalam db akan di encrypt terlebih dahulu dengan hash |
| is_admin     | boolean   | Boolean dengan value true / false                                                        |
##### Contoh Request
```json
{
  "full_name": "Mega Aulia R",
  "birth_date": "1997-01-14",
  "address": "Griya Praja Mukti",
  "phone_number": "087839221567",
  "email": "mega@emails.com",
  "password": "adminOK123",
  "is_admin": true
}
```
##### Contoh Response Sukses
```json
{
  "data": {
    "id": 1,
    "full_name": "Mega Aulia R",
    "birth_date": "1997-01-14T00:00:00Z",
    "address": "Griya Praja Mukti",
    "phone_number": "087839221567",
    "email": "mega@email.com",
    "password": "$2a$14$kX77sehmcIUGQUIfU0o3LeSPIx8EcRapTjW.OLFVaYlg2/l8RkctG",
    "created_at": "2023-02-22T04:37:12.368301Z",
    "updated_at": "2023-02-22T04:37:12.368301Z",
    "is_admin": true,
    "Token": ""
  },
  "message": "Success insert customer"
}
```
##### Contoh Response Gagal - Email sudah terdaftar
```json
{
  "error_message": [
    "email already registered"
  ],
  "error_status": "Invalid parameter"
}
```
##### Contoh Response Gagal - Salah format birth_date
```json
{
  "error_message": [
    "parameter 'birth_date' must be in format yyyy-MM-dd"
  ],
  "error_status": "Invalid parameter"
}
```
##### Contoh Response Gagal - Salah format full_name
```json
{
  "error_message": [
    "parameter 'full_name' must be in alphabet only"
  ],
  "error_status": "Invalid parameter"
}
```
##### Contoh Response Gagal - Salah format phone_number
```json
{
  "error_message": [
    "prefix 'phone_number' must be '08' or '62' or '+62' and max 13 digit"
  ],
  "error_status": "Invalid parameter"
}
```


### LOGIN
```
  POST /login
```
##### Parameter
| Parameter | Tipe Data | Deskripsi                                   |
|-----------|-----------|---------------------------------------------|
| email     | string    | harus dengan email yang terdaftar           | 
| password  | string    | harus sesuai dengan password yang terdaftar |
##### Contoh Request
```json
{
  "email": "mega@email.com",
  "password": "adminOK123"
}
```
##### Contoh Response Sukses
```json
{
  "message": "Success",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzcwNDEyNTQsIkVtYWlsIjoibWVnYUBlbWFpbC5jb20iLCJQYXNzd29yZCI6ImFkbWluT0sxMjMiLCJSb2xlIjoiYWRtaW4ifQ.qfctmzyfBT97it3VMUQfPkRGoo8HKus_xX51vgI3j6Q"
}

```
##### Contoh Response Gagal - Salah Token
```json
{
  "error": "Unauthorized, error parsing JWT"
}
```
##### Contoh Response Gagal - Salah User Role
```json
{
  "error": "forbidden access to API"
}
```
##### Contoh Response Gagal - Salah password
```json
{
  "error_message": [
    "incorrect password"
  ],
  "error_status": "Invalid parameter"
}
```

### WALLET
##### Parameter
| Parameter      | Tipe Data  | Deskripsi                                   |
|----------------|------------|---------------------------------------------|
| balance        | float      | Balance amount yang akan ditop-up ke wallet | 
| account_number | float      | account_number customer                     |
#### 1. Get Wallet By Customer Id
```
  GET /customer/:id/wallet
```
##### Contoh Response Sukses
```json
{
  "data": {
    "id": 1,
    "balance": 0,
    "account_name": "Mega Aulia R",
    "account_number": 48498095,
    "created_at": "2023-02-22T04:37:12.372989Z",
    "updated_at": "2023-02-22T04:37:12.372989Z",
    "customer_id": 1
  }
}
```
#### 2. Top Up Wallet
```
  PUT /wallet/top_up
```
##### Contoh Request
```json
{
  "balance": 2000000,
  "account_number": 48498095
}
```
##### Contoh Response Sukses
```json
{
    "data": {
        "id": 1,
        "balance": 4000000,
        "account_name": "Mega Aulia R",
        "account_number": 48498095,
        "created_at": "2023-02-22T04:37:12.372989Z",
        "updated_at": "2023-02-22T06:51:00.598452Z",
        "customer_id": 1
    },
    "message": "Success top up wallet for account number: 48498095"
}
```

### CATEGORY
##### Parameter
| Parameter     | Tipe Data | Deskripsi                                                            |
|---------------|-----------|----------------------------------------------------------------------|
| name          | string    | Nama categori event, contohnya: Konser, Seminar, Drama Musikal, dsb. |

#### 1. Create
```
  POST /bo/categories
```
##### Contoh Request
```json
{
  "name": "Concert"
}
```
##### Contoh Response Sukses
```json
{
  "data": {
    "id": 1,
    "name": "Concert",
    "created_at": "2023-02-22T04:41:16.791925Z",
    "updated_at": "2023-02-22T04:41:16.791925Z"
  },
  "message": "Success insert category"
}
```
#### 2. Update
```
  PUT /bo/categories/:id
```
##### Contoh Request
```json
{
  "name": "Drama Musikal"
}
```
##### Contoh Response Sukses
```json
{
  "data": {
    "id": 1,
    "name": "Drama Musikal",
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z"
  },
  "message": "Success update category with ID : 1"
}
```
#### 3. Delete
```
  DELETE /bo/categories/:id
```
##### Contoh Response Sukses
```json
{
  "message": "Success delete category with ID : 1"
}
```
##### Contoh Response Gagal 
```json
{
  "error_message": "category with ID : 2 not found",
  "error_status": "Data not found"
}
```
#### 4. Get All Categories
```
  GET /categories
```
##### Contoh Response Sukses
```json
{
  "data": [
    {
      "id": 1,
      "name": "Drama Musikal",
      "created_at": "2023-02-22T04:41:16.791925Z",
      "updated_at": "2023-02-22T04:42:35.547394Z"
    }
  ]
}
```


### EVENT
##### Parameter
| Parameter     | Tipe Data   | Deskripsi                                                |
|---------------|-------------|----------------------------------------------------------|
| name          | string      | Nama event, contohnya: Konser Akbar All Star             |
| description   | string      | Deskripsi event, contoh: Lokasi event                    |
| start_date    | string      | Tanggal dimulainya event, harus dalam format yyyy-MM-dd  |
| end_date      | string      | Tanggal berakhirnya event, harus dalam format yyyy-MM-dd |
| category_id   | int         | Category_id misalnya id = 1 untuk konser                 |

#### 1. Create
```
  POST /bo/event
```
##### Contoh Request
```json
{
  "name": "Blackpink World Tour 2023",
  "description": "Concert Blackpink ICE BSD",
  "start_date": "2023-05-05",
  "end_date": "2023-05-08",
  "category_id": 2
}
```
##### Contoh Response Sukses
```json
{
  "data": {
    "id": 2,
    "name": "Blackpink World Tour 2023",
    "description": "Concert Blackpink ICE BSD",
    "start_date": "2023-05-05T00:00:00Z",
    "end_date": "2023-05-05T00:00:00Z",
    "created_at": "2023-02-22T04:49:57.289862Z",
    "updated_at": "2023-02-22T04:49:57.289862Z",
    "category_id": 2
  },
  "message": "Success insert event"
}
```
##### Contoh Response Gagal
```json
{
  "error_message": [
    "parameter 'end_date' must be greater than 'start_date'"
  ],
  "error_status": "Invalid parameter"
}
```
```json
{
  "error_message": [
    "parameter 'start_date' must be in format yyyy-MM-dd"
  ],
  "error_status": "Invalid parameter"
}
```
```json
{
  "error_message": [
    "parameter 'end_date' must be in format yyyy-MM-dd"
  ],
  "error_status": "Invalid parameter"
}
```
```json
{
  "error_message": [
    "parameter 'start_date' cannot less than today"
  ],
  "error_status": "Invalid parameter"
}
```

#### 2. Update
```
  PUT /bo/event/:id
```
##### Contoh Request
```json
{
  "name": "Blackpink World Tour 2023",
  "description": "Concert Blackpink ICE BSD",
  "start_date": "2023-05-05",
  "end_date": "2023-05-08",
  "category_id": 2
}
```
##### Contoh Response Sukses
```json
{
  "data": {
    "id": 2,
    "name": "Blackpink World Tour 2023",
    "description": "Concert Blackpink ICE BSD",
    "start_date": "2023-05-05T00:00:00Z",
    "end_date": "2023-05-05T00:00:00Z",
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z",
    "category_id": 2
  },
  "message": "Success update event with ID : 2"
}

```
#### 3. Delete
```
  DELETE /bo/event/:id
```
##### Contoh Response Sukses
```json
{
  "message": "Success delete event with ID : 2"
}
```
##### Contoh Response Gagal
```json
{
  "error_message": "event with ID : 1 not found",
  "error_status": "Data not found"
}
```
#### 4. Get All Event by Category
```
GET /categories/:id/event
```
##### Contoh Response Sukses
```json
{
  "data": [
    {
      "id": 2,
      "name": "Blackpink World Tour 2023",
      "description": "Concert Blackpink ICE BSD",
      "start_date": "2023-05-05T00:00:00Z",
      "end_date": "2023-05-05T00:00:00Z",
      "category_id": 2,
      "category_name": "Concert"
    }
  ]
}
```
#### 5. Get All Event
```
GET /categories/
```
##### Contoh Response Sukses
```json
{
  "data": [
    {
      "id": 2,
      "name": "Blackpink World Tour 2023",
      "description": "Concert Blackpink ICE BSD",
      "start_date": "2023-05-05T00:00:00Z",
      "end_date": "2023-05-05T00:00:00Z",
      "category_id": 2,
      "category_name": "Concert"
    }
  ]
}
```


### TICKET
##### Parameter
| Parameter | Tipe Data | Deskripsi                                                                                                                 |
|-----------|-----------|---------------------------------------------------------------------------------------------------------------------------|
| name      | string    | Nama event, contohnya: Presale day 1, Kelas VVIP, Kelas REGULER                                                           |
| date      | string    | Tanggal sesuai dengan event. harus dalam format yyyy-MM-dd dan dalam rentang waktu start_date dan end_date di tabel event |
| quota     | string    | Quota tiket, akan selalu berkurang ketika ada transaksi yang masuk                                                        |
| price     | int       | Harga tiket                                                                                                               |
| event_id  | int       | EventId dari tiket, contohnya event_id = 1, Blackpink World Tour 2023                                                     |

#### 1. Create
```
  POST /bo/tickets
```
##### Contoh Request
```json
{
  "name": "Presale 1 - BLACKPINK DAY",
  "date": "2023-05-05",
  "quota": 50,
  "price": "2000000",
  "event_id": 3
}
```
##### Contoh Response Sukses
```json
{
  "data": {
    "id": 1,
    "name": "Presale Day 1 - BLACKPINK DAY",
    "date": "2023-05-05T00:00:00Z",
    "quota": 50,
    "price": "2000000",
    "event_id": 3,
    "created_at": "2023-02-22T04:56:51.651793Z",
    "updated_at": "2023-02-22T04:56:51.651793Z"
  },
  "message": "Success insert ticket"
}
```
##### Contoh Response Gagal
```json
{
  "error_message": [
    "parameter 'date' must between 2023-05-05 00:00:00 +0000 +0000 or 2023-05-08 00:00:00 +0000 +0000"
  ],
  "error_status": "Invalid parameter"
}
```
```json
{
  "error_message": [
    "parameter 'date' must be in format yyyy-MM-dd"
  ],
  "error_status": "Invalid parameter"
}
```
```json
{
  "error_message": [
    "parameter 'date' cannot less than today"
  ],
  "error_status": "Invalid parameter"
}
```
#### 2. Update
```
  PUT /bo/tickets/:id
```
##### Contoh Request
```json
{
  "name": "Presale Day 1 - BLACKPINK DAY",
  "date": "2023-05-05",
  "quota": 50,
  "price": "2000000",
  "event_id": 3
}
```
##### Contoh Response Sukses
```json
{
  "data": {
    "id": 1,
    "name": "Presale Day 2 - BLACKPINK DAY",
    "date": "2023-05-05T00:00:00Z",
    "quota": 50,
    "price": "2000000",
    "event_id": 3,
    "created_at": "2023-02-22T04:56:51.651793Z",
    "updated_at": "2023-02-22T04:56:51.651793Z"
  },
  "message": "Success insert ticket"
}

```
#### 3. Delete
```
  DELETE /bo/tickets/:id
```
##### Contoh Response Sukses
```json
{
  "message": "Success delete tickets with ID : 2"
}
```
##### Contoh Response Gagal
```json
{
  "error_message": "tickets with ID : 1 not found",
  "error_status": "Data not found"
}
```

#### 4. Get All Tickets by Event Id
```
  GET /event/:id/tickets
```
##### Contoh Response Sukses
```json
{
  "data": [
    {
      "id": 1,
      "name": "Presale Day 1 - BLACKPINK DAY",
      "date": "2023-05-05T00:00:00Z",
      "quota": 50,
      "price": "2000000",
      "event_id": 3,
      "event_name": "Blackpink World Tour 2023",
      "created_at": "2023-02-22T04:56:51.651793Z",
      "updated_at": "2023-02-22T06:18:31.892316Z"
    },
    {
      "id": 2,
      "name": "Presale Day 2 - BLACKPINK DAY",
      "date": "2023-05-06T00:00:00Z",
      "quota": 50,
      "price": "2000000",
      "event_id": 3,
      "event_name": "Blackpink World Tour 2023",
      "created_at": "2023-02-22T06:13:07.995953Z",
      "updated_at": "2023-02-22T06:19:14.399846Z"
    },
    {
      "id": 3,
      "name": "Presale DAY 3 - BLACKPINK DAY",
      "date": "2023-05-07T00:00:00Z",
      "quota": 50,
      "price": "2000000",
      "event_id": 3,
      "event_name": "Blackpink World Tour 2023",
      "created_at": "2023-02-22T06:17:23.2785Z",
      "updated_at": "2023-02-22T06:17:23.2785Z"
    },
    {
      "id": 4,
      "name": "Presale DAY 4 - BLACKPINK DAY",
      "date": "2023-05-08T00:00:00Z",
      "quota": 50,
      "price": "2000000",
      "event_id": 3,
      "event_name": "Blackpink World Tour 2023",
      "created_at": "2023-02-22T06:18:57.115703Z",
      "updated_at": "2023-02-22T06:18:57.115703Z"
    }
  ]
}
```

#### 5. Get Tickets By Id
```
  GET /tickets/:id
```
##### Contoh Response Sukses
```json
{
  "data": {
    "id": 2,
    "name": "Presale Day 2 - BLACKPINK DAY",
    "date": "2023-05-07T00:00:00Z",
    "quota": 49,
    "price": "2000000",
    "event_id": 3,
    "created_at": "2023-02-22T06:13:07.995953Z",
    "updated_at": "2023-02-22T06:52:38.83204Z"
  }
}
```


### CUSTOMER
##### Parameter
| Parameter    | Tipe Data |  Deskripsi                                                                               |
|--------------|-----------|------------------------------------------------------------------------------------------|
| full_name    | string    | Nama harus dalam alfabet, spesial character tidak diperbolehkan                          | 
| birth_date   | string    | Format birth_date harus yyyy-MM-dd                                                       |
| address      | string    | Alamat customer                                                                          |
| phone_number | string    | Format harus diawali dengan 0 atau +62 atau 62 dengan jumlah maksimal 13 digit           |
| email        | string    | Format harus sesuai dengan email, contoh: author@email.com                               |
| password     | string    | Password customer, saat diinsert ke dalam db akan di encrypt terlebih dahulu dengan hash |
| is_admin     | boolean   | Boolean dengan value true / false                                                        |

#### 1. Update
```
  PUT /customer/:id
```
##### Contoh Request
```json
{
  "full_name": "Mega Aulia R",
  "birth_date": "1997-01-14",
  "address": "Griya Praja Mukti",
  "phone_number": "087839221567",
  "email": "mega@doku.com",
  "password": "password",
  "is_admin": true
}
```
##### Contoh Response Sukses
```json
{
  "data": {
    "id": 1,
    "full_name": "Mega Aulia R",
    "birth_date": "1997-01-14T00:00:00Z",
    "address": "Griya Praja Mukti",
    "phone_number": "087839221567",
    "email": "mega@doku.com",
    "password": "$2a$14$xtZBRWNjX.IKQZ/PPKmM6eI9ZovKqzPkal16UO4cGRrUvArWH5voK",
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z",
    "is_admin": true,
    "Token": ""
  },
  "message": "Success update customer with ID : 1"
}
```
#### 2. Get Customer by Id
```
  GET /customer/:id
```
##### Contoh Response Sukses
```json
{
  "data": {
    "id": 1,
    "full_name": "Mega Aulia R",
    "birth_date": "1997-01-14T00:00:00Z",
    "address": "Griya Praja Mukti",
    "phone_number": "087839221567",
    "email": "mega@email.com",
    "password": "$2a$14$kX77sehmcIUGQUIfU0o3LeSPIx8EcRapTjW.OLFVaYlg2/l8RkctG",
    "created_at": "2023-02-22T04:37:12.368301Z",
    "updated_at": "2023-02-22T04:37:12.368301Z",
    "is_admin": true,
    "Token": ""
  }
}
```
#### 3. Get All Customer
```
  GET /bo/customer
```
##### Contoh Response Sukses
```json
{
  "data": {
    "id": 1,
    "full_name": "Mega Aulia R",
    "birth_date": "1997-01-14T00:00:00Z",
    "address": "Griya Praja Mukti",
    "phone_number": "087839221567",
    "email": "mega@email.com",
    "password": "$2a$14$kX77sehmcIUGQUIfU0o3LeSPIx8EcRapTjW.OLFVaYlg2/l8RkctG",
    "created_at": "2023-02-22T04:37:12.368301Z",
    "updated_at": "2023-02-22T04:37:12.368301Z",
    "is_admin": true,
    "Token": ""
  }
}
```


### TRANSACTION
##### Parameter
| Parameter   | Tipe Data | Deskripsi                                                                                      |
|-------------|-----------|------------------------------------------------------------------------------------------------|
| date        | string    | Tanggal transaksi harus dalam format yyyy-MM-dd dan harus sesuai dengan date pada kolom ticket |
| ticket_id   | int       | TicketId, contoh ticket_id = 1, Presale day 2 konser Blackpink 2023                            |
| customer_id | int       | CustomerId, contohnya customer_id = 1, Mega Aulia                                              |

#### 1. Create
Saat melakukan create transactions, sistem back-end akan memvalidasi semua paramater. Ketika semua validasi sudah lolos, sistem akan melanjutkan mengenerate QR Code dalam bentuk file '.png' dan selanjutnya dilakukan proses upload ke CDN (di sini saya menggunakan cloudinary) agar file QR tersebut dapat diakses oleh customer.
```
  POST /transactions
```
##### Contoh Request
```json
{
  "date": "2023-05-05",
  "ticket_id": 2,
  "customer_id": 1
}
```
##### Contoh Response Sukses
```json
{
  "data": {
    "id": 1,
    "date": "2023-05-05T00:00:00Z",
    "qr_code": "https://res.cloudinary.com/dojng2aaa/image/upload/v1677048758/ticket/rjj28nsgivvxwdkc9q19.png",
    "created_at": "2023-02-22T06:52:38.827112Z",
    "updated_at": "2023-02-22T06:52:38.827112Z",
    "ticket_id": 2,
    "customer_id": 1
  },
  "message": "Success insert transaction"
}
```
##### Contoh Response Gagal - Balance tidak cukup
```json
{
  "error_message": [
    "insufficient balance"
  ],
  "error_status": "Invalid parameter"
}
```
##### Contoh Response Gagal - Date tidak sesuai dengan data ticket / format salah
```json
{
  "error_message": [
    "parameter 'date' must be in format yyyy-MM-dd",
    "parameter 'start_date' cannot less than today",
    "insufficient balance"
  ],
  "error_status": "Invalid parameter"
}
```
##### Contoh Response Gagal - Quota ticket sudah habis
```json
{
  "error_message": [
    "ticket is sold out"
  ],
  "error_status": "Invalid parameter"
}
```
#### 2. Get All Transaction
```
  GET /bo/transactions
```
##### Contoh Response Sukses
```json
{
  "data": [
    {
      "id": 1,
      "qr_code": "https://res.cloudinary.com/dojng2aaa/image/upload/v1677048758/ticket/rjj28nsgivvxwdkc9q19.png",
      "created_at": "2023-02-22T06:52:38.827112Z",
      "customer_id": "1",
      "customer_name": "Mega Aulia R",
      "email": "mega@email.com",
      "phone_number": "087839221567",
      "ticket_id": "2",
      "ticket_name": "Presale Day 2 - BLACKPINK DAY",
      "ticket_date": "2023-05-07T00:00:00Z",
      "price": "2000000",
      "event_id": "3",
      "event_name": "Blackpink World Tour 2023"
    }
  ]
}
```
#### 3. Get All Transaction By Customer Id
```
  GET /customer/:id/transaction
```
##### Contoh Response Sukses
```json
{
  "data": [
    {
      "id": 1,
      "qr_code": "https://res.cloudinary.com/dojng2aaa/image/upload/v1677048758/ticket/rjj28nsgivvxwdkc9q19.png",
      "created_at": "2023-02-22T06:52:38.827112Z",
      "customer_id": "1",
      "customer_name": "Mega Aulia R",
      "email": "mega@email.com",
      "phone_number": "087839221567",
      "ticket_id": "2",
      "ticket_name": "Presale Day 2 - BLACKPINK DAY",
      "ticket_date": "2023-05-07T00:00:00Z",
      "price": "2000000",
      "event_id": "3",
      "event_name": "Blackpink World Tour 2023"
    },
    {
      "id": 2,
      "qr_code": "https://res.cloudinary.com/dojng2aaa/image/upload/v1677048758/ticket/rjj28nsgivvxwde098374.png",
      "created_at": "2023-02-22T06:52:38.827112Z",
      "customer_id": "1",
      "customer_name": "Mega Aulia R",
      "email": "mega@email.com",
      "phone_number": "087839221567",
      "ticket_id": "5",
      "ticket_name": "REGULER - JUSTIN BIEBER",
      "ticket_date": "2023-05-07T00:00:00Z",
      "price": "2000000",
      "event_id": "7",
      "event_name": "Justin World Tour 2023"
    }
  ]
}
```
#### 4. Get Transaction By Id
```
  GET /transaction/:id
```
##### Contoh Response Sukses
```json
{
  "data": 
    {
      "id": 1,
      "qr_code": "https://res.cloudinary.com/dojng2aaa/image/upload/v1677048758/ticket/rjj28nsgivvxwdkc9q19.png",
      "created_at": "2023-02-22T06:52:38.827112Z",
      "customer_id": "1",
      "customer_name": "Mega Aulia R",
      "email": "mega@email.com",
      "phone_number": "087839221567",
      "ticket_id": "2",
      "ticket_name": "Presale Day 2 - BLACKPINK DAY",
      "ticket_date": "2023-05-07T00:00:00Z",
      "price": "2000000",
      "event_id": "3",
      "event_name": "Blackpink World Tour 2023"
    }
}
```
