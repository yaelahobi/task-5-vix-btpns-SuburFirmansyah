
## Deskripsi
Berdasarkan data yang telah diolah oleh tim Data Analysts, bahwa untuk  
meningkatkan engagement user pada aplikasi m-banking adalah meningkatkan  
aspek memiliki user pada aplikasi tersebut. Saran yang diberikan oleh tim data  
analysts adalah membentuk fitur **personalize user**, salah satunya adalah  
memungkinkan **user dapat mengupload gambar untuk dijadikan foto profilnya**. Tim  
developer bertanggung jawab untuk mengembangkan fitur ini, dan kalian diberikan  
tugas untuk merancang API pada fitur upload, dan menghapus gambar. Beberapa  
ketentuannya antara lain :

- User dapat menambahkan foto profile
- Sistem dapat mengidentifikasi User ( log in / sign up)
- Hanya user yang telah login / sign up yang dapat melakukan delete / tambah  
  foto profil
- User dapat menghapus gambar yang telah di post
- User yang berbeda tidak dapat menghapus / mengubah foto yang telah di  
  buat oleh user lain

Buatlah API menggunakan bahasa GoLang yang sesuai dengan ketentuan dan  
kebutuhan diatas!

## Endpoints
Pada bagian User Endpoint :
1. POST : /users/register, dan gunakan atribut berikut ini :
    - ID (primary key, required)
    - Username (required)
    - Email (unique & required)
    - Password (required & minlength 6)
    - Relasi dengan model Photo (Gunakan constraint cascade)
    - Created At (timestamp)
    - Updated At (timestamp)
2. GET: /users/login
    - Using email & password (required)
3. PUT : /users/:userId (Update User)
4. DELETE : /users/:userId (Delete User)

Photos Endpoint
1. POST : /photos
    - ID
    - Title
    - Caption
    - PhotoUrl
    - UserID
    - Relasi dengan model User
2. GET : /photos
3. PUT : /photoId
4. DELETE : /:photoId

## Requirement
1. Authorization dapat menggunakan tool [Go JWT](https://github.com/dgrijalva/jwt-go)
2. Pastikan hanya user yang membuat foto yang dapat menghapus / mengubah foto

## Repo Structure
Struktur dokumen / environment dari GoLang yang akan dibentuk kurang lebih sebagai berikut :
- **app** : Menampung pembuatan struct dalam kasus ini menggunakan struct User untuk keperluan data dan authentication
- **controllers** : Berisi antara logic database yaitu models dan query
- **database** : Berisi konfigurasi database serta digunakan untuk menjalankan koneksi database dan migration
- **helpers** : Berisi fungsi-fungsi yang dapat digunakan di setiap tempat dalam hal ini jwt, bcrypt, headerValue
- **middlewares** : Berisi fungsi yang digunakan untuk proses otentikasi jwt yang digunakan untuk proteksi api
- **models** : Berisi models yang digunakan untuk relasi database
- **router** : Berisi konfigurasi routing / endpoint yang akan digunakan untuk mengakses api
- **go mod** : Yang digunakan untuk manajemen package / dependency berupa library

## Tools & Libraries
Tools yang dapat kalian gunakan :
- [Gin Gonic Framework](https://github.com/gin-gonic/gin)
- [Gorm](https://gorm.io/index.html)
- [JWT Go](https://github.com/dgrijalva/jwt-go)
- [Go Validator](http://github.com/asaskevich/govalidator)