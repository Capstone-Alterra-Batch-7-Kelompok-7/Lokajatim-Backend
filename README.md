# Lokajatim Backend

Lokajatim Backend adalah backend service yang dikembangkan menggunakan Golang untuk mendukung aplikasi **Lokajatim**, platform yang menyediakan informasi dan layanan terkait Jawa Timur.

## 📑 Fitur Utama
- **Autentikasi dan Autorisasi**: Mendukung login, registrasi, dan manajemen peran pengguna.
- **Manajemen Data**: CRUD untuk data layanan dan informasi yang disediakan.
- **API Terstruktur**: Mendukung integrasi dengan frontend melalui REST API.
- **Keamanan**: JWT untuk autentikasi, validasi input, dan sanitasi data.
- **Testing**: Unit testing untuk memastikan stabilitas aplikasi.

## 🛠️ Teknologi yang Digunakan
- **Bahasa Pemrograman**: Golang
- **Framework**: Echo (untuk HTTP server)
- **Database**: MySQL
- **ORM**: GORM
- **Middleware**: JWT, CORS
- **Deployment**: AWS
- **Tooling**: Swagger untuk dokumentasi API

## 📂 Struktur Proyek
```
Lokajatim-Backend/
├── config/         # Konfigurasi aplikasi (database, JWT, dll.)
├── controllers/    # Logika bisnis dan handler untuk HTTP request
├── models/         # Model database
├── routes/         # Routing untuk endpoint API
├── services/       # Logika layanan yang terpisah dari controller
├── utils/          # Fungsi pendukung
├── tests/          # File unit testing
└── main.go         # Entry point aplikasi
```

## 🚀 Cara Menjalankan Proyek

### Prasyarat
Pastikan Anda sudah menginstal:
- **Go**: Versi 1.21 atau lebih baru
- **Database**: MySQL
- **Git**: Untuk meng-clone repository

### Langkah-Langkah
1. Clone repository ini:
   ```bash
   git clone https://github.com/Capstone-Alterra-Batch-7-Kelompok-7/Lokajatim-Backend.git
   cd Lokajatim-Backend
   ```

2. Konfigurasi file `.env`:
   Buat file `.env` di root project dan tambahkan konfigurasi berikut sebagai contoh:
   ```
   DATABASE_USER="root"
   DATABASE_PASSWORD=""
   DATABASE_HOST="localhost"
   DATABASE_PORT="3306"
   DATABASE_NAME="lokajatim_db"
   JWT_SECRET_KEY="your_jwt_secret_key"
   GEMINI_API_KEY="your_gemini_api_key"
   MIDTRANS_SERVER_KEY="your_midtrans_server_key"
   MIDTRANS_CLIENT_KEY="your_midtrans_client_key"
   SMTP_EMAIL="your_smtp_email"
   SMTP_PASSWORD="your_smtp_password"
   SMTP_HOST="smtp.gmail.com"
   SMTP_PORT="587"
   ```

3. Jalankan perintah berikut untuk menginstal dependency:
   ```bash
   go mod tidy
   ```

4. Migrasikan database:
   ```bash
   go run main.go migrate
   ```

5. Jalankan server:
   ```bash
   go run main.go
   ```

Aplikasi akan berjalan pada `http://localhost:8080`.

## 📜 Dokumentasi API
Gunakan Swagger untuk melihat dokumentasi API. Setelah server berjalan, buka:
```
http://localhost:8080/swagger/index.html
```

## 🧪 Testing
Untuk menjalankan unit test, gunakan perintah berikut:
```bash
go test ./...
```

---

**Dikembangkan oleh Kelompok 7, Capstone Project Alterra Batch 7.**
