# iLib Perpustakaan API

## Daftar Isi

- [Pra-kata](#pra-kata)

- [Bagaimana *clone project* ini?](#Clone-project-ini)

- [Documentation](#Documentation)
    - [HTTP Requests](#HTTP-Requests)
    - [HTTP Responses](#soal-1b)
    - [HTTP Response Code](#soal-1c)
    - [GET](#soal-1d)
    - [POST](#soal-1e)
    - [PUT](#soal-1f)
    - [DELETE](#soal-1g)

- [Authorization & Authentication](#Authorization-&-Authentication)

- [Contoh penggunaan melalui Postman](#Tutorial-Postman)


---

## Pra-kata
---
iLib merupakan tugas dari matkul `Pemrograman Integratif` yang dibuat oleh `Muhammad Hilmi Ramadhan` dengan `NRP 05311940000044` pada *project-2*. Restful API ini mensimulasikan dari sekumpulan data  `author` yang tersedia pada API ini. Yang mana hal ini di*build* menggunakan bahasa pemrograman `Go` dengan *framework* `echo labstack` tepatnya pada versi `v4`. API ini memiliki 4 metode dan 4 endpoint yang berbeda, yang mana untuk melengkapi syarat dari *Authentication* maka disini saya menerapkan `JWT Token` yang juga akan berkaitan dengan *Authorization* nya nanti. Secara sederhana jika bukan pengguna terdaftar maka pengguna tersebut hanya dapat mengakses `metode GET` dari `author` yang tersedia pada *database*, akan tetapi jika pengguna terdaftar dalam artian melakukan *login* terlebih dahulu yang mana nantinya juga memiliki `JWT Token` yang mengakibatkan pengguna tersebut dapat mengakses `iLib Restful API` secara penuh ke - 4 method yang telah di buat `(GET, POST, PUT, DELETE)`.

<br>


## Clone project ini
---
1. Pastikan sudah terinstal `Golang` beserta dengan `Echo labstack frameworknya`. Jika belum silahkan klik tautan berikut : <br>
[Golang Windows](https://www.tutorialpemrograman.com/others/cara-install-golang-di-windows-10/) , [Golang Linux](https://gegeriyadi.com/install-golang/), [Golang MacOS](https://tecadmin.net/install-go-on-macos/), dan [Echo Labstack Framework](https://echo.labstack.com/)
2. Jika telah tersedia point 1 pada device anda, maka letakkan project ini pada `GO-PATH` yang telah didefinisikan saat instalasi `Golang`.
3. Selanjutnya lakukan inisialisasi dengan command pada terminal/shell anda `go mod init` dan `go mod tidy` untuk otomatisasi *import* terhadap `library/file-file` pendukung yang dibutuhkan. Tak lupa, jika pada header `file import go` juga disesuaikan terhadap penyimpanan `foldering device`  masing-masing.
4. Lalu lakukan `run and execution` program dengan command <br> `go run main.go`

<br>


## Documentation
---
1. HTTP Requests <br>
Segala permintaan yang dikirim oleh `API` dibuat seaman mungkin menggunakan HTTPS dengan salah satu metode berikut, yang mana bergantung dari `reqeust` yang diambil : 

    * `GET` yang berguna untuk mendapatkan seluruh list yang ada dari *resources*
    * `POST` yang berguna untuk menambah atau membuat sebuah *inputan* data yang baru terhadap *resources*
    * `PUT` yang berguna untuk melakukan *edit* data dari yang telah tersedia dari *resources*
    * `DELETE` yang berguna untuk menghapus data yang tersedia dari *resources*<br> \
Terkhusus pada `metode POST & PUT` pada `body request` harus dilengkapi beberapa parameter pada `header` khusus seperti diantaranya jika pada `POST` harus menyertakan `Nama_author` yang akan ditargetkan dan pada `PUT` harus menyertakan `Id_author` dan `Nama_author`.

2. HTTP Responses <br>
Setiap response yang dikembalikan dari `requests method` akan mengembalikan `status` dan jika berhasil maka akan ditampilkan dalam bentuk `json single-record queries and an array for list queries`. Tidak hanya itu, akan diberikan juga sebuah pesan berupa `message : "Success"` jika berhasil dan yang terakhir tentunya `data` yang berupa `json array for list queries`. Adapun contohnya sebagai berikut :  <br>
```json
{
    "status": 200,
    "message": "Success",
    "data": [
        {
            "id": 1,
            "nama": "Archibald Gislason"
        },
        {
            "id": 2,
            "nama": "Mr. Saul Muller DDS"
        },
        {
            "id": 3,
            "nama": "Virgie Paucek I"
        },
        {
            "id": 4,
            "nama": "Anastacio Cremin"
        },
        {
            "id": 5,
            "nama": "Cody Eichmann"
        },
        {
            "id": 6,
            "nama": "Stanford Hirthe"
        },
        {
            "id": 7,
            "nama": "Efrain Ullrich"
        },
        {
            "id": 8,
            "nama": "Dr. Carmel Bergstrom IV"
        },
        {
            "id": 9,
            "nama": "Prof. Rosendo Price I"
        },
        {
            "id": 10,
            "nama": "Demarco Carroll"
        },
        {
            "id": 11,
            "nama": "Chadrick Howell"
        },
        {
            "id": 12,
            "nama": "Prof. Jerry Labadie"
        },
        {
            "id": 13,
            "nama": "Mr. Columbus Wisoky"
        },
        {
            "id": 14,
            "nama": "Jody Hodkiewicz"
        },
        {
            "id": 15,
            "nama": "Carmen Bechtelar II"
        },
        {
            "id": 16,
            "nama": "Alfreda Kilback"
        },
        {
            "id": 17,
            "nama": "Prof. Vallie Fay PhD"
        },
        {
            "id": 18,
            "nama": "Caroline Mills"
        },
        {
            "id": 19,
            "nama": "Jules Ullrich"
        },
        {
            "id": 20,
            "nama": "Nadia Greenholt"
        },
        {
            "id": 21,
            "nama": "R.V. Hari Ginardi"
        },
        {
            "id": 22,
            "nama": "Baskoro Adi Pratomo"
        },
        {
            "id": 24,
            "nama": "Bapak Ridho Rahman Hariadi"
        }
    ]
}
```

3. HTTP Response Codes <br>
Setiap *response* akan dikembalikan melalui `HTTP status code` sebagai berikut :
    * `200` `OK` yang bermakna *requests* yang diminta berhasil di dapatkan.
    * `400` `Bad Request` yang bermakna terjadi sebuah kesalahan terhadap *requests* nya (*security*, *malformed*, *data validation*, dll.)
    * `401` `Unauthorized` yang bermakna bahwa `API` gagal mengenali kredensial dari validasinya.
    * `403` `Forbidden` yang bermakna bahwa kredensial tidak dapat memiliki akses terhadap *resources* yang diminta.
    * `404` `Method not allowed` yang bermakna bahwa *resources* tidak tepat dalam pengaksesan `method` nya. (Membutuhkan *requests* `method` yang spesifik seperti `GET`, `POST`, dll).
    * `500` `Server Error` yang bermakna bahwa terjadi kegalatan yang disebabkan oleh server (Bisa server blm menyala dari mysqlnya ataupun hal lain).

<br>

4. GET <br>
Pada `metode` ini siapapun dapat melihat isi dari `endpoint` ini tanpa perlu adanya `login` atau `Authentication` dengan `JWT Token`

    **`GET`**    /localhost:1234/author

    **Contoh Code Go**
    ```go
    func get() {
    fmt.Println("1. Performing Http Get...")
    resp, err := http.Get("https://localhost:1234/author")
    if err != nil {
        log.Fatalln(err)
    }

    defer resp.Body.Close()
    bodyBytes, _ := ioutil.ReadAll(resp.Body)

    // Convert response body to string
    bodyString := string(bodyBytes)
    fmt.Println("API Response as String:\n" + bodyString)

    // Convert response body to Todo struct
    var todoStruct Todo
    json.Unmarshal(bodyBytes, &todoStruct)
    fmt.Printf("API Response as struct %+v\n", todoStruct)
    }
    ```
    **Response**
    ```
    Status : 200
    ```
    ```json
    {
    "status": 200,
    "message": "Success",
    "data": [
        {
            "id": 1,
            "nama": "Archibald Gislason"
        },
        {
            "id": 2,
            "nama": "Mr. Saul Muller DDS"
        }
    ]
    }
    ```
    
<br>

5. POST<br>
Pada `metode` ini hanya pengguna terdaftar yang dapat mengakses `endpoint` tentunya memerlukan `login` atau `Authentication` dengan `JWT Token` yang telah didapat dimasukan ke dalam `Bearer Token` nya. `metode` ini memungkinkan pengguna terdaftar untuk menambahkan penulis pada `iLib Perpustakaan API`
    
    **`POST`**    /localhost:1234/author

    **Authorization**
    | Type | Token|
    | -----| -----|
    | Bearer Token | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBfZGF0ZSI6MTYyMDYzOTE5NiwibGV2ZWwiOiJBcHBsaWNhdGlvbiIsInVzZXJuYW1lIjoibWhpbG1pIn0.Kfpzf7VSO_-STPFxhUs7uHq4v-I0iStg8riOrPIEZww |
    <br>

    **Parameters**

    | KEY  | VALUE |
    | ------------- | ------------- |
    | Nama_author  | Bapak Ridho Rahman Hariadi  |
    

    <br>

    **Contoh Code**
    Go
    ```go
    func post() {
    fmt.Println("2. Performing Http Post...")
    todo := Todo{1, 2, "lorem ipsum dolor sit amet", true}
    jsonReq, err := json.Marshal(todo)
    resp, err := http.Post("https://localhost:1234/author", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
    if err != nil {
        log.Fatalln(err)
    }

    defer resp.Body.Close()
    bodyBytes, _ := ioutil.ReadAll(resp.Body)

    // Convert response body to string
    bodyString := string(bodyBytes)
    fmt.Println(bodyString)

    // Convert response body to Todo struct
    var todoStruct Todo
    json.Unmarshal(bodyBytes, &todoStruct)
    fmt.Printf("%+v\n", todoStruct)
    }
    ```
     **Response**
    ```
    Status : 200
    ```
    ```json
    {
        "status": 200,
        "message": "Success",
        "data": {
            "Id yang terakhir dimasukan adalah": 24
        }
    }
    ```
    <br>

6. PUT<br>
Pada `metode` ini hanya pengguna terdaftar yang dapat mengakses `endpoint` tentunya memerlukan `login` atau  `Authentication` dengan `JWT Token` yang telah didapat dimasukan ke dalam `Bearer Token` nya. `metode` ini memungkinkan pengguna terdaftar `edit` data dari *resources* `author` yang ada pada `API` ini.

    **`PUT`**    /localhost:1234/author

    **Authorization**
    | Type | Token|
    | -----| -----|
    | Bearer Token | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBfZGF0ZSI6MTYyMDYzOTE5NiwibGV2ZWwiOiJBcHBsaWNhdGlvbiIsInVzZXJuYW1lIjoibWhpbG1pIn0.Kfpzf7VSO_-STPFxhUs7uHq4v-I0iStg8riOrPIEZww |
    <br>

    **Parameters**

    | KEY  | VALUE |
    | ------------- | ------------- |
    | Id_author  | 1  |
    | Nama_author | Muhammad Hilmi Ramadhan |
    

    <br>

    **Contoh Code**
    Go
    ```go
    func put() {
    fmt.Println("3. Performing Http Put...")
    todo := Todo{1, 2, "lorem ipsum dolor sit amet", true}
    jsonReq, err := json.Marshal(todo)
    req, err := http.NewRequest(http.MethodPut, "https://localhost:1234/author", bytes.NewBuffer(jsonReq))
    req.Header.Set("Content-Type", "application/json; charset=utf-8")
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatalln(err)
    }

    defer resp.Body.Close()
    bodyBytes, _ := ioutil.ReadAll(resp.Body)

    // Convert response body to string
    bodyString := string(bodyBytes)
    fmt.Println(bodyString)

    // Convert response body to Todo struct
    var todoStruct Todo
    json.Unmarshal(bodyBytes, &todoStruct)
    fmt.Printf("API Response as struct:\n%+v\n", todoStruct)
    }
    ```
     **Response**
    ```
    Status : 200
    ```
    ```json
    {
        "status": 200,
            "message": "Success",
            "data": {
                "Data terupdate sebanyak": 1
            }
    }
    ```
    <br>

7. DELETE<br>
Pada `metode` ini hanya pengguna terdaftar yang dapat mengakses `endpoint` tentunya memerlukan `login` atau  `Authentication` dengan `JWT Token` yang telah didapat dimasukan ke dalam `Bearer Token` nya. `metode` ini memungkinkan pengguna terdaftar `delete` data dari *resources* `author` yang ada pada `API` ini.

    **`DELETE`**    /localhost:1234/author

    **Authorization**
    | Type | Token|
    | -----| -----|
    | Bearer Token | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBfZGF0ZSI6MTYyMDYzOTE5NiwibGV2ZWwiOiJBcHBsaWNhdGlvbiIsInVzZXJuYW1lIjoibWhpbG1pIn0.Kfpzf7VSO_-STPFxhUs7uHq4v-I0iStg8riOrPIEZww |
    <br>

    **Parameters**

    | KEY  | VALUE |
    | ------------- | ------------- |
    | Id_author  | 24  |
    

    <br>

    **Contoh Code**
    Go
    ```go
    func delete() {
    fmt.Println("4. Performing Http Delete...")
    todo := Todo{1, 2, "lorem ipsum dolor sit amet", true}
    jsonReq, err := json.Marshal(todo)
    req, err := http.NewRequest(http.MethodDelete, "https://localhost:1234/author", bytes.NewBuffer(jsonReq))
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatalln(err)
    }

    defer resp.Body.Close()
    bodyBytes, _ := ioutil.ReadAll(resp.Body)

    // Convert response body to string
    bodyString := string(bodyBytes)
    fmt.Println(bodyString)
    }
    ```
     **Response**
    ```
    Status : 200
    ```
    ```json
    {
        "status": 200,
            "message": "Success",
            "data": {
            "Data terhapus sebanyak": 1
            }
    }
    ```

<br>

## Authorization & Authentication
---
Pada *project* ini saya mengenakan bantuan dari `JWT` yang mana saya integrasikan dari `library` tautan berikut `https://wwww.github.com/dgrijalva/jwt-go`. Pada dasarnya `JWT` akan muncul ketika pengguna melakukan validasi `login` sederhana yang dengan kata lain pengguna harus terdaftar terlebih dahulu di *database* guna dapat mengakses ke-4 metode dari `Restful API` yang telah dibuat. Oleh karena itu saya telah menyiapkan sebuah akun *dummy* dengan `username : mhilmi` dan `password : asdf`. Secara konseptual nantinya akan di cocokan antara `password plainteks` dengan `password cipherteks` yang telah dikonversikan ke bentuk `bcrypt`. Berikut merupakan gambar dari `JWT Token` yang diperoleh oleh *user* setelah melakukan validasi *login* sederhana melalui `POSTMAN` :
![Postman2](./ReadmeAsset/postman/setupPostman2.png)<br>
Gambar diatas dapat diperoleh melalui `metode POST` dengan mengarahkan url ke `localhost:1234/login?username=mhilmi&password=asdf`
<br>

Adapun untuk *Authorization* dan *Authentication* untuk *Guest* atau pengguna tidak terdaftar tanpa melalui *login* maka yang dapat diakses hanya `metode GET` dari Penulis Buku yang tersedia di *database* seperti gambar berikut :
![Postman2](./ReadmeAsset/postman/GuesAuthorandAuth.png)<br>


<br>


## Tutorial Postman
---
Postman merupakan sebuah *tools* dan *testing* dari sebuah `API`. Disini akan ditunjukan bagaimana melakukan konfigurasi sebagaimana mestinya agar `iLib Restful API` dapat dicoba : <br>
1. Pertama-tama maka buatlah sebuah *collection* khusus guna menampung foldering dari testing `iLib API` <br>
![Postman1](./ReadmeAsset/postman/setupPostman1.png)
2. Setelah sudah diberi nama bebas, maka berikut salah satu contoh testing dengan metode `POST` dalam menambahkan author.Tentunya dengan melalui otentikasi `JWT Token` terelebih dahulu dengan login melalui `Postman` yang me-return `JWT`, kemudian baru dapat melakukan `POST`. Pada kasus ini akan login dengan `username` mhilmi dan `password` asdf <br>
sehingga url untuk mendapatkan `JWT Token` akan berubah menjadi `localhost:1234/login?username=mhilmi&password=asdf`
![Postman2](./ReadmeAsset/postman/setupPostman2.png)<br>
Setelah mendapatkan `token` seperti diatas, maka mari kita coba untuk `POST` nya, sebelum itu maka set terlebih dahulu pada `Authorization` kita masukan `Bearer Token` yang di dapat dari `JWT Token` setelah login tadi dengan url `localhost:1234/author` sebagai berikut :
![Postman2](./ReadmeAsset/postman/setupPostman3.png)
Mari kita coba untuk `POST`
![Postman2](./ReadmeAsset/postman/setupPostman4.png)
Dan telah berhasil. Kira kira seperti itu dalam tutorial postman dalam pengujian dari `iLib Restful API`.